package models

import "go.mongodb.org/mongo-driver/bson"

type AisClassInfo uint8

func (in AisClassInfo) MarshalBSON() ([]byte, error) {

	b, err := bson.Marshal(&struct {
		Id    AisClassId `bson:"id"`
		Label string     `bson:"label"`
	}{
		Id:    AisClassId(in),
		Label: AisClassId(in).AsString(),
	})
	return b, err

}

func (in *AisClassInfo) UnmarshalBSON(data []byte) error {

	aux := &struct {
		Id AisClassId `bson:"id"`
	}{}
	if err := bson.Unmarshal(data, aux); err != nil {
		return err
	}
	*in = AisClassInfo(aux.Id)
	return nil

}
