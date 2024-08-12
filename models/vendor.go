package models

type Vendor struct {
	Name   string `bson:"name"`
	Model  uint8  `bson:"model"`
	Serial uint32 `bson:"serial"`
}
