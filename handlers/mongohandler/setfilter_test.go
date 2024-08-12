package mongohandler

import (
	"ais-stream/models"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// flag whether to delete created data on exit (e.g. for testing)
const SET_FILTER_CLEANUP bool = false

// short delay on start for resources to fully load
const SET_FILTER_STARTUP_DELAY = time.Duration(2 * time.Second)

func TestSetFilter(t *testing.T) {

	// create a filter to upload
	config := &FilterConfig{
		Name: "test",
		MessageWhitelist: &WhitelistConfig{
			Name: "messages",
			//Boundary:    nil,
			//Mmsis:       []models.MMSI{},
			MessageIds: models.StandardMessages,
			//ShipTypeIds: []models.ShipTypeId{},
		},
		VesselWhitelist: &WhitelistConfig{
			Name: "vessels",
			//Boundary:    nil,
			//Mmsis:       []models.MMSI{},
			MessageIds:  models.StandardMessages,
			ShipTypeIds: models.SpecialCraft,
		},
		PositionWhitelist: &WhitelistConfig{
			Name:     "positions",
			Boundary: models.BoundaryAustralia,
			//Mmsis:       []models.MMSI{},
			MessageIds:  models.StandardMessages,
			ShipTypeIds: models.SpecialCraft,
		},
	}

	// load required resources to insert configs
	handler := New(
		30*time.Second,
		"LOCAL_MONGODB_CONNECTION",
	)
	handler.Connect("default")

	// short delay for everything to fire up (mongo takes a while)
	time.Sleep(SET_FILTER_STARTUP_DELAY)

	// upload the new config
	filter := bson.D{{Key: "name", Value: config.Name}}
	opts := options.FindOneAndReplace().SetUpsert(true)
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	result := handler.collections[filtersCollection].FindOneAndReplace(timeoutCtx, filter, config, opts)
	assert.NoError(t, result.Err(), "error setting filter config")

	// cleanup if required (e.g. while testing)
	if SET_FILTER_CLEANUP {
		timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
		defer cancel()
		result := handler.collections[filtersCollection].FindOneAndDelete(timeoutCtx, filter)
		assert.NoError(t, result.Err(), "error cleaning up filter config")
	}
}
