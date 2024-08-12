package mongohandler

import (
	"ais-stream/interfaces"
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (p *MongoHandler) GetRecord(mmsi uint32) (*interfaces.Record, error) {

	filter := bson.D{{Key: "mmsi", Value: mmsi}}
	result := interfaces.Record{Mmsi: mmsi}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	err := p.collections[vesselsCollection].FindOne(timeoutCtx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &result, nil
		}
		slog.Error("error retrieving record", "error", err, "collection", vesselsCollection)
	}
	return &result, err

}
