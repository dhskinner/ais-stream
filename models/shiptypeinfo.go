package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type ShipTypeInfo uint8

func (in ShipTypeInfo) MarshalBSON() ([]byte, error) {

	b, err := bson.Marshal(&struct {
		Id       ShipTypeId `bson:"id"`
		Label    string     `bson:"label"`
		Category string     `bson:"category"`
	}{
		Id:       ShipTypeId(in),
		Label:    ShipTypeId(in).AsString(),
		Category: ShipTypeId(in).AsCategory().AsString(),
	})
	return b, err

}

func (in *ShipTypeInfo) UnmarshalBSON(data []byte) error {

	aux := &struct {
		Id ShipTypeId `bson:"id,omitempty"`
	}{}
	if err := bson.Unmarshal(data, aux); err != nil {
		return err
	}
	*in = ShipTypeInfo(aux.Id)
	return nil

}
