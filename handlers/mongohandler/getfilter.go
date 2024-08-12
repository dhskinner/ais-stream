package mongohandler

import (
	"ais-stream/models"
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (p *MongoHandler) GetFilter(name string) (*FilterConfig, error) {

	filter := bson.D{{Key: "name", Value: name}}
	result := FilterConfig{}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	err := p.collections[filtersCollection].FindOne(timeoutCtx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &result, nil
		}
		slog.Error("error retrieving filter", "error", err, "collection", filtersCollection)
	}
	return &result, err

}

type FilterConfig struct {
	Name              string           `bson:"name"`
	MessageWhitelist  *WhitelistConfig `bson:"messagewhitelist,omitempty"`
	VesselWhitelist   *WhitelistConfig `bson:"vesselwhitelist,omitempty"`
	PositionWhitelist *WhitelistConfig `bson:"positionwhitelist,omitempty"`
}

type WhitelistConfig struct {
	Name        string              `bson:"name"`
	Boundary    *models.Boundary    `bson:"boundary,omitempty"`
	Mmsis       []models.MMSI       `bson:"mmsis,omitempty"`
	MessageIds  []models.MessageId  `bson:"messageids,omitempty"`
	ShipTypeIds []models.ShipTypeId `bson:"shiptypeids,omitempty"`
}
