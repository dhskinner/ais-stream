package atlas

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO this needs to be transferred to a config file
const (
	localConnectionKey           string = "LOCAL_MONGODB_CONNECTION"
	localDatabaseName            string = "ais2"
	localPositionsCollectionName string = "positions"
	localVesselsCollectionName   string = "vessels"

	atlasConnectionKey           string = "ATLAS_MONGODB_CONNECTION"
	atlasDatabaseName            string = "ais"
	atlasPositionsCollectionName string = "positions"
	atlasVesselsCollectionName   string = "vessels"

	vesselsInterval time.Duration = 1 * time.Minute
	vesselsOffset   time.Duration = 0

	positionsInterval   time.Duration = 1 * time.Minute
	positionsOffset     time.Duration = 10 * time.Second
	positionsMinimimSog float32       = 0.1
)

// Client to aggregate local database updates into atlas
// This runs on a ticker performing:
// - position updates for Queensland each one minute
// - vessel updates for Queensland each 5 minutes

type Atlas struct {
	localClient      *mongo.Client
	atlasClient      *mongo.Client
	localCollections map[string]*mongo.Collection
	atlasCollections map[string]*mongo.Collection
}

func New() *Atlas {

	a := &Atlas{
		localCollections: make(map[string]*mongo.Collection),
		atlasCollections: make(map[string]*mongo.Collection),
	}
	return a

}

func (a *Atlas) Connect() error {

	// get the environment variables
	localConnectionString := os.Getenv(localConnectionKey)
	if localConnectionString == "" {
		err := fmt.Errorf("you must set your '%s' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable", localConnectionKey)
		slog.Error("error getting database address", "error", err)
		return err
	}

	atlasConnectionString := os.Getenv(atlasConnectionKey)
	if atlasConnectionString == "" {
		err := fmt.Errorf("you must set your '%s' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable", atlasConnectionKey)
		slog.Error("error getting database address", "error", err)
		return err
	}

	// connect to Atlas hosted mongodb
	var err error = nil
	a.localClient, err = newClient(localConnectionString)
	if err != nil {
		return err
	}

	// connect to the local MongoDB database
	a.atlasClient, err = newClient(atlasConnectionString)
	if err != nil {
		return err
	}

	// fire up some collections
	opts := &options.CollectionOptions{
		BSONOptions: &options.BSONOptions{
			IntMinSize:     true,
			OmitZeroStruct: true,
		},
	}

	a.localCollections[localPositionsCollectionName] = a.localClient.Database(localDatabaseName).Collection(localPositionsCollectionName, opts)
	a.localCollections[localVesselsCollectionName] = a.localClient.Database(localDatabaseName).Collection(localVesselsCollectionName, opts)
	a.atlasCollections[atlasPositionsCollectionName] = a.atlasClient.Database(atlasDatabaseName).Collection(atlasPositionsCollectionName, opts)
	a.atlasCollections[atlasVesselsCollectionName] = a.atlasClient.Database(atlasDatabaseName).Collection(atlasVesselsCollectionName, opts)
	return nil

}

func (a *Atlas) Disconnect() error {

	if err := a.atlasClient.Disconnect(context.TODO()); err != nil {
		return err
	}

	if err := a.localClient.Disconnect(context.TODO()); err != nil {
		return err
	}
	return nil

}

func (a *Atlas) Process(ctx context.Context, wg *sync.WaitGroup) {

	// tell the caller we've stopped
	defer wg.Done()

	// fire up mongo
	err := a.Connect()
	if err != nil {
		return
	}
	defer a.Disconnect()

worker:
	for {
		select {

		// check for cancel signal
		case <-ctx.Done():
			break worker

		// wait for a ticker, then update vessels
		case <-time.After(timeToGo(vesselsInterval, vesselsOffset)):
			a.runVesselsAggregation(int(vesselsInterval.Minutes()))

		// wait for a ticker, then update positions
		case <-time.After(timeToGo(positionsInterval, positionsOffset)):
			a.runPositionsAggregation(int(positionsInterval.Minutes()), positionsMinimimSog)
		}
	}
}

func newClient(connectionString string) (*mongo.Client, error) {

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()
	client, err := mongo.Connect(timeoutCtx, opts)
	if err != nil {
		slog.Error("error connecting to mongodb", "error", err)
		return nil, err
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		slog.Error("error could not ping mongodb atlas", "error", err)
		return nil, err
	}

	slog.Info("successfully connected to mongodb atlas")
	return client, nil

}
