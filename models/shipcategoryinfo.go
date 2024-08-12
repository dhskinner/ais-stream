package models

import "go.mongodb.org/mongo-driver/bson"

type ShipCategoryInfo uint8

func (in ShipCategoryInfo) MarshalBSON() ([]byte, error) {

	b, err := bson.Marshal(&struct {
		Id    ShipCategoryId `bson:"id,omitempty"`
		Label string         `bson:"label,omitempty"`
	}{
		Id:    ShipCategoryId(in),
		Label: ShipCategoryId(in).AsString(),
	})
	return b, err

}

func (in *ShipCategoryInfo) UnmarshalBSON(data []byte) error {

	aux := &struct {
		Id ShipCategoryId `bson:"id,omitempty"`
	}{}
	if err := bson.Unmarshal(data, aux); err != nil {
		return err
	}
	*in = ShipCategoryInfo(aux.Id)
	return nil

}
