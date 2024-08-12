package models

import "go.mongodb.org/mongo-driver/bson"

type Aton struct {
	Id          AtonId `bson:"id,omitempty"`
	Label       string `bson:"label,omitempty"`
	OffPosition bool   `bson:"offposition,omitempty"`
	Virtual     bool   `bson:"virtual,omitempty"`
}

func (in *Aton) MarshalBSON() ([]byte, error) {

	if in == nil {
		return bson.Marshal(&struct{}{})
	}

	return bson.Marshal(&struct {
		Id          AtonId `bson:"id,omitempty"`
		Label       string `bson:"label,omitempty"`
		OffPosition bool   `bson:"offpos,omitempty"`
		Virtual     bool   `bson:"virtual,omitempty"`
	}{
		Id:          AtonId(in.Id),
		Label:       AtonId(in.Id).AsString(),
		OffPosition: in.OffPosition,
		Virtual:     in.Virtual,
	})
}

func (in *Aton) UnmarshalBSON(data []byte) error {

	aux := &struct {
		Id          uint8 `bson:"id,omitempty"`
		OffPosition bool  `bson:"offpos,omitempty"`
		Virtual     bool  `bson:"virtual,omitempty"`
	}{}

	if err := bson.Unmarshal(data, aux); err != nil {
		return err
	}
	in.Id = (AtonId)(aux.Id)
	in.OffPosition = aux.OffPosition
	in.Virtual = aux.Virtual

	return nil
}
