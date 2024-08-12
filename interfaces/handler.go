package interfaces

import (
	"ais-stream/models"
)

type Handler interface {
	Message(message models.Message) error
	GetRecord(mmsi models.MMSI) (*Record, error)
}

type Record struct {
	Mmsi     models.MMSI        `bson:"mmsi"`
	ShipType models.ShipTypeId  `bson:"type"`
	Position models.Coordinates `bson:"pos"`
}
