package models

import "go.mongodb.org/mongo-driver/bson"

type NavigationInfo uint8

func (in NavigationInfo) MarshalBSON() ([]byte, error) {

	b, err := bson.Marshal(&struct {
		Id    NavigationId `bson:"id"`
		Label string       `bson:"label"`
	}{
		Id:    NavigationId(in),
		Label: NavigationId(in).AsShortString(),
	})
	return b, err

}

func (in *NavigationInfo) UnmarshalBSON(data []byte) error {

	aux := &struct {
		Id NavigationId `bson:"id"`
	}{}
	if err := bson.Unmarshal(data, aux); err != nil {
		return err
	}
	*in = NavigationInfo(aux.Id)
	return nil

}
