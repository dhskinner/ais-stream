package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type ShipRecord struct {
	Mmsi     MMSI        `bson:"mmsi"`
	ShipType ShipTypeId  `bson:"shiptype"`
	Position Coordinates `bson:"pos"`
}

func (in *ShipRecord) UnmarshalBSON(data []byte) error {

	aux := &struct {
		Mmsi MMSI       `bson:"mmsi"`
		Type ShipTypeId `bson:"shiptype,omitempty"`
	}{}
	if err := bson.Unmarshal(data, aux); err != nil {
		//return err
		in.Mmsi = aux.Mmsi
		in.ShipType = 0
		return nil
	}
	in.Mmsi = aux.Mmsi
	in.ShipType = aux.Type
	return nil

}
