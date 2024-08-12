package models

type VendorInfo struct {
	Name   string `bson:"name,omitempty"`
	Model  uint8  `bson:"model,omitempty"`
	Serial uint32 `bson:"serial,omitempty"`
}
