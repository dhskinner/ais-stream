package models

type AtlasMetadata struct {
	Mmsi         MMSI        `bson:"mmsi"`
	Comment      string      `bson:"comment,omitempty"`
	Fleet        string      `bson:"fleet,omitempty"`    // not populated
	HomePort     GeoPosition `bson:"homeport,omitempty"` // not populated
	Organisation string      `bson:"org,omitempty"`      // not populated
	Style        Style       `bson:"style,omitempty"`    // not populated
}
