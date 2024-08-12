package atlas

import (
	"ais-stream/models"
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Aggregate vessels from local mongodb to Atlas at set intervals
func (a *Atlas) runPositionsAggregation(minutes int, minimumSog float32) error {

	// allow a small buffer on updates so they overlap a little
	// TODO this needs to be transferred to a config file
	intervalBufferMinutes := 1

	// create the aggregation document
	pipeline := bson.A{
		bson.D{
			{"$match",
				bson.D{
					{"$expr",
						bson.D{
							{"$gte",
								bson.A{
									"$time",
									bson.D{
										{"$dateSubtract",
											bson.D{
												{"startDate", "$$NOW"},
												{"unit", "minute"},
												{"amount", minutes + intervalBufferMinutes},
											},
										},
									},
								},
							},
						},
					},
					{"sog", bson.D{{"$gte", minimumSog}}},
				},
			},
		},
		bson.D{
			{"$set",
				bson.D{
					{"min",
						bson.D{
							{"$dateTrunc",
								bson.D{
									{"date", "$time"},
									{"unit", "minute"},
									{"binSize", 1},
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{"$group",
				bson.D{
					{"_id",
						bson.D{
							{"mmsi", "$mmsi"},
							{"min", "$min"},
						},
					},
					{"pos", bson.D{{"$push", "$$ROOT"}}},
				},
			},
		},
		bson.D{
			{"$set",
				bson.D{
					{"pos",
						bson.D{
							{"$sortArray",
								bson.D{
									{"input", "$pos"},
									{"sortBy", bson.D{{"time", -1}}},
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "vessels"},
					{"localField", "_id.mmsi"},
					{"foreignField", "mmsi"},
					{"as", "vessel"},
				},
			},
		},
		bson.D{
			{"$set",
				bson.D{
					{"vessel",
						bson.D{
							{"$arrayElemAt",
								bson.A{
									"$vessel",
									0,
								},
							},
						},
					},
					{"pos",
						bson.D{
							{"$arrayElemAt",
								bson.A{
									"$pos",
									0,
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{"$replaceWith",
				bson.D{
					{"mmsi", "$_id.mmsi"},
					{"min", "$_id.min"},
					{"time", "$pos.time"},
					{"cog", "$pos.cog"},
					{"sog", "$pos.sog"},
					{"alt", "$pos.alt"},
					{"pos", "$pos.pos"},
					{"state", "$vessel.state"},
				},
			},
		},
		bson.D{{"$match", bson.D{{"state", "QLD"}}}},
	}

	// pass the pipeline to the Aggregate() method
	readCtx, readCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer readCancel()
	cursor, err := a.localCollections[localVesselsCollectionName].Aggregate(readCtx, pipeline)
	if err != nil {
		slog.Error("error running positions aggregation", "error", err)
		return err
	}

	// collate the results
	var positions []models.AtlasPosition
	if err = cursor.All(readCtx, &positions); err != nil {
		slog.Error("error collating positions aggregation results", "error", err)
		return err
	}

	// run a bulk update to insert into atlas
	m := []mongo.WriteModel{}
	for _, position := range positions {
		if !position.Position.IsValid() {
			continue
		}
		filter := bson.D{{Key: "mmsi", Value: position.Mmsi}, {Key: "min", Value: position.Min}}
		m = append(m, mongo.NewReplaceOneModel().SetUpsert(true).SetFilter(filter).SetReplacement(position))
	}
	writeCtx, writeCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer writeCancel()
	results, err := a.atlasCollections[atlasPositionsCollectionName].BulkWrite(writeCtx, m)
	if err != nil {
		slog.Error("error bulk writing positions to atlas", "error", err, "results", results)
		return err
	}
	return nil
}
