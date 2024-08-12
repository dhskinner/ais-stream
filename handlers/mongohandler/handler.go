package mongohandler

import (
	"ais-stream/handlers"
	"ais-stream/handlers/deduplicator"
	"ais-stream/handlers/mongohandler/filter"
	"ais-stream/models"
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/BertoldVdb/go-ais"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName        string = "ais2"
	stationsCollection  string = "stations"
	positionsCollection string = "positions"
	vesselsCollection   string = "vessels"
	safetyCollection    string = "safety"
	atonCollection      string = "aton"
	filtersCollection   string = "filters"

	MONGO_CONNECTION_RETRY_SECS = 10
)

// Default handler for incoming sentences
type MongoHandler struct {
	queue            chan models.Message
	addressKey       string
	isConnected      bool
	duplicateTimeout time.Duration
	client           *mongo.Client
	collections      map[string]*mongo.Collection
	filter           *filter.Filter
}

func New(
	duplicateTimeout time.Duration,
	addressKey string,
) *MongoHandler {

	// create a new mongohandler
	p := &MongoHandler{
		queue:            make(chan models.Message, 1024),
		isConnected:      false,
		duplicateTimeout: duplicateTimeout,
		addressKey:       addressKey,
		collections:      make(map[string]*mongo.Collection),
	}
	return p
}

func (p *MongoHandler) Connect(filterName string) error {

	uri := os.Getenv(p.addressKey)
	if uri == "" {
		err := fmt.Errorf("you must set your '%s' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable", p.addressKey)
		slog.Error("error getting database address", "error", err)
		return err
	}

	var err error = nil
	p.client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		slog.Error("error connecting to mongodb", "uri", uri, "error", err)
		return err
	}
	p.isConnected = true
	opts := &options.CollectionOptions{
		BSONOptions: &options.BSONOptions{
			IntMinSize:     true,
			OmitZeroStruct: true,
		},
	}
	p.collections[atonCollection] = p.client.Database(databaseName).Collection(atonCollection, opts)
	p.collections[safetyCollection] = p.client.Database(databaseName).Collection(safetyCollection, opts)
	p.collections[vesselsCollection] = p.client.Database(databaseName).Collection(vesselsCollection, opts)
	p.collections[positionsCollection] = p.client.Database(databaseName).Collection(positionsCollection, opts)
	p.collections[stationsCollection] = p.client.Database(databaseName).Collection(stationsCollection, opts)
	p.collections[filtersCollection] = p.client.Database(databaseName).Collection(filtersCollection, opts)

	// add a filter (n.b. shiptypes only filters positions)
	config, err := p.GetFilter(filterName)
	if err != nil {
		slog.Error("error getting filter configuration", "error", err)
		return err
	}

	// TODO this is clunky af - sort this all out
	messageWhitelist := filter.NewWhitelist(
		config.MessageWhitelist.Boundary,
		config.MessageWhitelist.Mmsis,
		config.MessageWhitelist.MessageIds,
		config.MessageWhitelist.ShipTypeIds,
	)
	vesselWhitelist := filter.NewWhitelist(
		config.VesselWhitelist.Boundary,
		config.VesselWhitelist.Mmsis,
		config.VesselWhitelist.MessageIds,
		config.VesselWhitelist.ShipTypeIds,
	)
	positionWhitelist := filter.NewWhitelist(
		config.PositionWhitelist.Boundary,
		config.PositionWhitelist.Mmsis,
		config.PositionWhitelist.MessageIds,
		config.PositionWhitelist.ShipTypeIds,
	)
	f := filter.New(config.Name, messageWhitelist, vesselWhitelist, positionWhitelist, p)
	p.filter = f
	return nil

}

func (p *MongoHandler) Disconnect() error {

	p.isConnected = false
	if p.isConnected {
		if err := p.client.Disconnect(context.TODO()); err != nil {
			return err
		}
	}
	return nil

}

func (p *MongoHandler) Message(message models.Message) error {

	if message == nil {
		return fmt.Errorf("error cannot add nil message")
	}
	p.queue <- message
	return nil

}

// Processes incoming sentences - default handler just outputs these to the console
func (p *MongoHandler) Process(ctx context.Context, wg *sync.WaitGroup, filterName string) {

	// tell the caller we've stopped
	defer wg.Done()

	for {

		// run a new worker
		err := p.run(ctx, filterName)
		if err != nil {
			slog.Error("mongo handler: error", "error", err)
		}

		// on error (and if not cancelled), automatically restart
		select {
		case <-ctx.Done():
			slog.Info("mongo handler: stopped worker")
			return
		default:
			time.Sleep(time.Duration(MONGO_CONNECTION_RETRY_SECS) * time.Second)
		}
	}
}

// Processes incoming sentences - default handler just outputs these to the console
func (p *MongoHandler) run(ctx context.Context, filterName string) error {

	// create a deduplicator to tag repeated messages
	dedup := deduplicator.New(p.duplicateTimeout)
	go dedup.Start(ctx)

	// connect to mongodb
	err := p.Connect(filterName)
	if err != nil {
		return err
	}
	defer p.Disconnect()

	var counter uint64 = 0

worker:
	for {
		select {

		// check for cancel signal
		case <-ctx.Done():
			break worker

		// receive incoming sentences
		case message := <-p.queue:

			// filter message types that are not of interest
			counter++
			if !p.filter.IsMessageIncluded(message) {
				handlers.Print(counter, handlers.FilteredMessage, message)
				continue
			}

			// filter duplicate messages
			isDuplicate, _, _ := dedup.IsDuplicate(message)
			if isDuplicate {
				handlers.Print(counter, handlers.Duplicate, message)
				continue
			}

			// extract and save relevant data
			switch message.Packet.(type) {
			case ais.PositionReport:
				p.setPositionReport(message)
			case ais.ShipStaticData:
				p.setShipStaticData(message)
			case ais.StandardSearchAndRescueAircraftReport:
				p.setSearchAndRescueAircraftReport(message)
			case ais.SafetyBroadcastMessage:
				p.setSafetyBroadcastMessage(message)
			case ais.StandardClassBPositionReport:
				p.setStandardClassBPositionReport(message)
			case ais.ExtendedClassBPositionReport:
				p.setExtendedClassBPositionReport(message)
			case ais.StaticDataReport:
				p.setStaticDataReport(message)
			case ais.LongRangeAisBroadcastMessage:
				p.setLongRangeAisBroadcastMessage(message)
			case ais.BaseStationReport:
				p.setBaseStationReport(message)
			case ais.AidsToNavigationReport:
				p.setAidsToNavigationReport(message)
			default:
				// everything else is todo
			}
			handlers.Print(counter, handlers.OK, message)
		}
	}
	return nil

}

// insert each position report to Mongodb (this assumes that only
// inserts or deletes are allowable for a time series collection)
func (p *MongoHandler) InsertPosition(position *models.VesselPosition) error {

	// check the filter
	if !p.filter.IsPositionIncluded(position) {
		return nil
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	_, err := p.collections[positionsCollection].InsertOne(timeoutCtx, position)
	if err != nil {
		slog.Error("error inserting position", "error", err, "position", position)
	}
	return err

}

// upsert each report to Mongodb - note that this merges updates into
// the existing document. If there are two vessels using the same MMSI
// (as they tend to do) then data may flip flop - nothing we can do about that....
func (p *MongoHandler) Upsert(mmsi uint32, collection string, doc primitive.M) error {

	// check the filter
	var shiptype *models.ShipTypeId = nil
	var position *models.Coordinates = nil
	value, ok := doc["pos"]
	if ok {
		pos, ok := value.(models.Coordinates)
		if ok {
			position = &pos
		}
	}
	value, ok = doc["shiptype"]
	if ok {
		typ, ok := value.(models.ShipTypeId)
		if ok {
			shiptype = &typ
		}
	}

	if !p.filter.IsWhitelisted(mmsi, nil, shiptype, position, p.filter.VesselWhitelist) {
		return nil
	}

	filter := bson.D{{Key: "mmsi", Value: mmsi}}
	opts := options.Update().SetUpsert(true)
	update := bson.D{{Key: "$set", Value: doc}}
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	_, err := p.collections[collection].UpdateOne(timeoutCtx, filter, update, opts)
	if err != nil {
		slog.Error("error upserting document", "error", err, "collection", collection, "doc", doc)
	}
	return err

}
