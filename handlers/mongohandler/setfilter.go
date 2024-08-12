package mongohandler

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (p *MongoHandler) SetFilter(config *FilterConfig) error {

	filter := bson.D{{Key: "name", Value: config.Name}}
	opts := options.FindOneAndReplace().SetUpsert(true)
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	result := p.collections[filtersCollection].FindOneAndReplace(timeoutCtx, filter, config, opts)
	if result.Err() != nil {
		slog.Error("error retrieving filter", "error", result.Err(), "collection", filtersCollection)
	}
	return result.Err()

}
