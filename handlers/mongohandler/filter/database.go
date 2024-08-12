package filter

import (
	"ais-stream/interfaces"
	"ais-stream/models"
	"sync"
)

// 'Database' is a grand name to describe a simple in-memory sync-map to store a couple of
// vessel paramaters for filtering - remember theres around 300,000 vessels to store in RAM
type Database struct {
	handler interfaces.Handler
	records sync.Map
}

// create a new filter
func NewDatabase(handler interfaces.Handler) *Database {

	return &Database{
		handler: handler,
	}

}

func (d *Database) Set(record *interfaces.Record) {

	d.records.Store(record.Mmsi, record)

}

func (d *Database) Update(
	record *interfaces.Record,
	shiptype *models.ShipTypeId,
	position *models.Coordinates,
) {

	if shiptype != nil {
		record.ShipType = *shiptype
	}
	if position != nil {
		record.Position = *position
	}
	d.Set(record)

}

func (d *Database) GetAndUpdate(
	mmsi models.MMSI,
	shiptype *models.ShipTypeId,
	position *models.Coordinates,
) *interfaces.Record {

	record := d.Get(mmsi)
	go d.Update(record, shiptype, position)
	return record

}

// This will be slow initially until an in-memory cache is built up of
// mmsi-type keypairs - remember there's about 300K of these so need to
// consider RAM requirements
func (d *Database) Get(mmsi models.MMSI) *interfaces.Record {

	// is the vessel mmsi present in our in-memory cache?
	value, ok := d.records.Load(mmsi)
	if ok {
		record, ok := value.(*interfaces.Record)
		if ok {
			return record
		}
	}

	// nope - try to get it from mongo
	record, err := d.handler.GetRecord(mmsi)

	// cache the value in memory (if its unknown to mongo, then it will be nil)
	if err == nil {
		d.records.Store(mmsi, record)
	}
	return record

}
