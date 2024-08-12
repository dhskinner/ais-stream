package models

import (
	"time"
)

type AtlasVessel struct {
	Beam             Metres      `bson:"beam,omitempty"`
	CallSign         string      `bson:"callsign,omitempty"`
	AisClass         string      `bson:"class,omitempty"`
	Altitude         Metres      `bson:"alt,omitempty"`
	CourseOverGround Degrees     `bson:"cog,omitempty,truncate"`
	Dimension        *Dimension  `bson:"dimension,omitempty,truncate"`
	Destination      string      `bson:"dest,omitempty"`
	Draught          Metres      `bson:"draught,omitempty"`
	ETA              time.Time   `bson:"eta,omitempty"`
	ImoNumber        uint32      `bson:"imo,omitempty"`
	Length           Metres      `bson:"length,omitempty"`
	Mmsi             MMSI        `bson:"mmsi"`
	Name             string      `bson:"name,omitempty"`
	Navigation       string      `bson:"nav,omitempty"`
	Position         Coordinates `bson:"pos,omitempty"`
	SpeedOverGround  Knots       `bson:"sog,omitempty"`
	Source           string      `bson:"source,omitempty"`
	State            string      `bson:"state,omitempty"`
	Time             time.Time   `bson:"time"`
	ShipType         ShipType    `bson:"shiptype,omitempty"`
	Vendor           *Vendor     `bson:"vendor,omitempty"`
	Style            *Style      `bson:"style,omitempty"`
	Metadata         *Metadata   `bson:"metadata,omitempty"`
}

func NewAtlasVessel(in *VesselInfo) *AtlasVessel {

	return &AtlasVessel{
		Time:             in.Time,
		Mmsi:             in.Mmsi,
		Name:             in.Name,
		CallSign:         in.CallSign,
		ImoNumber:        in.ImoNumber,
		ShipType:         ShipType(in.ShipTypeId),
		Dimension:        in.Dimension,
		Length:           in.Dimension.Length(),
		Beam:             in.Dimension.Beam(),
		Draught:          in.MaximumStaticDraught,
		AisClass:         in.AisClass.AsString(),
		Position:         in.Position,
		SpeedOverGround:  in.SpeedOverGround,
		CourseOverGround: in.CourseOverGround,
		Source:           in.Source,
		State:            in.State,
		Navigation:       in.NavigationId.AsShortString(),
		Destination:      in.Destination,
		Altitude:         in.Altitude,
		Style:            in.Style,
		Metadata: &Metadata{
			Fleet:        in.Fleet,
			HomePort:     in.HomePort,
			HomePosition: in.HomePosition,
			Organisation: in.Organisation,
		},
	}
}
