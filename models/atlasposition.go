package models

import (
	"time"
)

type AtlasPosition struct {
	Time             time.Time   `bson:"time"`
	Min              time.Time   `bson:"min"`
	Mmsi             MMSI        `bson:"mmsi"`
	Position         Coordinates `bson:"pos"`
	SpeedOverGround  Knots       `bson:"sog,omitempty,truncate"`
	CourseOverGround Degrees     `bson:"cog,omitempty,truncate"`
	Altitude         Metres      `bson:"alt,omitempty,truncate"`
}
