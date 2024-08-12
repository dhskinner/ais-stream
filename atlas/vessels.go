package atlas

import (
	"ais-stream/models"
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Agregate vessels from local mongodb to Atlas at set intervals
func (a *Atlas) runVesselsAggregation(minutes int) error {

	// allow a small buffer on updates so they overlap a little
	// TODO this needs to be transferred to a config file
	intervalBufferMinutes := 1

	// create the aggregation document
	pipeline := bson.A{
		bson.D{
			{"$match",
				bson.D{
					{"state", "QLD"},
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
				},
			},
		},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$mmsi"},
					{"vessels", bson.D{{"$push", "$$ROOT"}}},
				},
			},
		},
		bson.D{
			{"$set",
				bson.D{
					{"vessels",
						bson.D{
							{"$sortArray",
								bson.D{
									{"input", "$vessels"},
									{"sortBy", bson.D{{"time", -1}}},
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{"$replaceRoot",
				bson.D{
					{"newRoot",
						bson.D{
							{"$arrayElemAt",
								bson.A{
									"$vessels",
									0,
								},
							},
						},
					},
				},
			},
		},
	}

	// pass the pipeline to the Aggregate() method
	readCtx, readCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer readCancel()
	cursor, err := a.localCollections[localVesselsCollectionName].Aggregate(readCtx, pipeline)
	if err != nil {
		slog.Error("error running vessels aggregation", "error", err)
		return err
	}

	// collate the results
	var vessels []*models.VesselInfo
	if err = cursor.All(readCtx, &vessels); err != nil {
		slog.Error("error collating vessel aggregation results", "error", err)
		return err
	}

	// run a bulk update to insert into atlas
	m := []mongo.WriteModel{}
	for _, vessel := range vessels {
		if !vessel.IsValid() {
			continue
		}
		// TODO sort out the aggregation to marshall straight to desired form
		v := models.NewAtlasVessel(vessel)
		filter := bson.D{{Key: "mmsi", Value: vessel.Mmsi}}
		m = append(m, mongo.NewReplaceOneModel().SetUpsert(true).SetFilter(filter).SetReplacement(v))
	}
	writeCtx, writeCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer writeCancel()
	results, err := a.atlasCollections[atlasVesselsCollectionName].BulkWrite(writeCtx, m)
	if err != nil {
		slog.Error("error bulk writing vessels to atlas", "error", err, "results", results)
		return err
	}
	return nil
}
