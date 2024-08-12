// Run a test routine to inspect each vessel record and clean it up as needed

package mongohandler

import (
	"ais-stream/models"
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestVesselClean(t *testing.T) {

	var count int

	// load required resources
	handler := New(
		30*time.Second,
		"LOCAL_MONGODB_CONNECTION",
	)
	handler.Connect("default")

	// short delay for everything to fire up (mongo takes a while)
	time.Sleep(SET_FILTER_STARTUP_DELAY)

	// loop through every vessel and load it into a data model
	filter := bson.D{{}}
	cursor, err := handler.collections[vesselsCollection].Find(context.TODO(), filter)
	assert.NoError(t, err, "error getting vessels cursor")

	for cursor.Next(context.TODO()) {
		var doc models.VesselInfo
		if err := cursor.Decode(&doc); err != nil {

			var bsonDoc bson.D
			if err := cursor.Decode(&bsonDoc); err != nil {
				log.Printf("Error decoding to bson: %+v - %+v", err, doc)
			}

			// result, err := handler.collections[vesselsCollection].ReplaceOne(context.TODO(), filter, doc)
			// log.Printf("Replaced mmsi: %d - %+v (error %+v)", doc.Mmsi, result, err)
			log.Printf("Replaced mmsi: %d - %+v (error %+v)", doc.Mmsi, bsonDoc, err)

		}
		//fmt.Printf("%+v\n", doc)
		count++
	}
	if err := cursor.Err(); err != nil {
		assert.NoError(t, err, "error traversing vessels cursor")
	}
	fmt.Printf("done count: %d\n", count)
}
