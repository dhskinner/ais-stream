package models

import (
	"time"
)

type VesselPosition struct {
	Time             time.Time   `bson:"time"`
	Mmsi             MMSI        `bson:"mmsi"`
	Position         Coordinates `bson:"pos,omitempty,truncate"`
	SpeedOverGround  Knots       `bson:"sog,omitempty,truncate"`
	CourseOverGround Degrees     `bson:"cog,omitempty,truncate"`
	Altitude         Metres      `bson:"alt,omitempty,truncate"`
	Source           string      `bson:"source,omitempty"`
}

func (v *VesselPosition) IsValid() bool {

	if v.Mmsi == 0 ||
		v.Mmsi > 999999999 ||
		v.Time.IsZero() ||
		!v.Position.IsValid() {
		return false
	}

	return true
}
