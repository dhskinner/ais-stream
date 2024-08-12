package models

type Metadata struct {
	Comment      string      `bson:"comment,omitempty"`
	Fleet        string      `bson:"fleet,omitempty"`
	HomePort     string      `bson:"homeport,omitempty"`
	HomePosition Coordinates `bson:"homepos,omitempty"`
	Organisation string      `bson:"org,omitempty"`
	Style        Style       `bson:"style,omitempty"`
}
