package models

import (
	"time"
)

type VesselInfo struct {
	Time                 time.Time    `bson:"time"`
	Mmsi                 MMSI         `bson:"mmsi"`
	Name                 string       `bson:"name,omitempty"`
	CallSign             string       `bson:"callsign,omitempty"`
	ShipTypeId           ShipTypeId   `bson:"shiptype,omitempty"`
	ImoNumber            uint32       `bson:"imo,omitempty"`
	Dimension            *Dimension   `bson:"dimension,omitempty,truncate"`
	MaximumStaticDraught Metres       `bson:"draught,omitempty,truncate"`
	AisClass             AisClass     `bson:"class,omitempty"`
	Vendor               *Vendor      `bson:"vendor,omitempty"`
	Position             Coordinates  `bson:"pos,omitempty,truncate"`
	Altitude             Metres       `bson:"alt,omitempty,truncate"`
	SpeedOverGround      Knots        `bson:"sog,omitempty,truncate"`
	CourseOverGround     Degrees      `bson:"cog,omitempty,truncate"`
	Source               string       `bson:"source,omitempty"`
	State                string       `bson:"state,omitempty"`
	NavigationId         NavigationId `bson:"nav,omitempty"`
	Destination          string       `bson:"dest,omitempty"`
	ETA                  time.Time    `bson:"eta,omitempty"`
	Comment              string       `bson:"comment,omitempty"`
	Fleet                string       `bson:"fleet,omitempty"`
	HomePort             string       `bson:"homeport,omitempty"`
	HomePosition         Coordinates  `bson:"homepos,omitempty"`
	Organisation         string       `bson:"org,omitempty"`
	Style                *Style       `bson:"style,omitempty"`
}

func (v *VesselInfo) IsValid() bool {

	if v.Mmsi == 0 ||
		v.Mmsi > 999999999 ||
		v.Time.IsZero() {
		return false
	}

	return true
}
