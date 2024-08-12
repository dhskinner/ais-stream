package models

import (
	"time"
)

type AtlasVessel struct {
	Time             time.Time      `bson:"time"`
	Mmsi             MMSI           `bson:"mmsi"`
	Name             string         `bson:"name"`
	CallSign         string         `bson:"callsign"`
	ImoNumber        uint32         `bson:"imo"`
	Metadata         *AtlasMetadata `bson:"metadata"`
	ShipType         ShipTypeInfo   `bson:"shiptype"`
	Dimension        Dimension      `bson:"dimension,truncate"`
	Length           Metres         `bson:"length,truncate"`
	Beam             Metres         `bson:"beam,truncate"`
	Draught          Metres         `bson:"draught,truncate"`
	AisClass         string         `bson:"aisclass"`
	Position         Coordinates    `bson:"pos"`
	SpeedOverGround  Knots          `bson:"sog,truncate"`
	CourseOverGround Degrees        `bson:"cog,truncate"`
	Source           string         `bson:"source"`
	State            string         `bson:"state"`
	NavigationStatus string         `bson:"nav"`
	Destination      string         `bson:"dest"`
	ETA              time.Time      `bson:"eta"`
	//Vendor           VendorInfo    `bson:"vendor"`
}

func NewAtlasVessel(in *VesselInfo) *AtlasVessel {

	return &AtlasVessel{
		Time:             in.Time,
		Mmsi:             in.Mmsi,
		Name:             in.Name,
		CallSign:         in.CallSign,
		ImoNumber:        in.ImoNumber,
		ShipType:         ShipTypeInfo(in.ShipType),
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
		NavigationStatus: in.NavigationStatus.AsShortString(),
		Destination:      in.Destination,
		ETA:              in.ETA,
				//Vendor:           &in.Vendor,
	}

}
