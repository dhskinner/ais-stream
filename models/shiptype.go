package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type ShipType uint8

func (in ShipType) MarshalBSON() ([]byte, error) {

	b, err := bson.Marshal(&struct {
		Id       ShipTypeId `bson:"id,omitempty"`
		Label    string     `bson:"label,omitempty"`
		Category string     `bson:"category,omitempty"`
	}{
		Id:       ShipTypeId(in),
		Label:    ShipTypeId(in).AsString(),
		Category: ShipTypeId(in).AsCategory().AsString(),
	})
	return b, err

}

func (in *ShipType) UnmarshalBSON(data []byte) error {

	aux := &struct {
		Id ShipTypeId `bson:"id,omitempty"`
	}{}
	if err := bson.Unmarshal(data, aux); err != nil {
		return err
	}
	*in = ShipType(aux.Id)
	return nil

}
