package models

type Style struct {
	VesselColor string `bson:"vesselcolor,omitempty"`
	FleetColor  string `bson:"fleetcolor,omitempty"`
	OrgColor    string `bson:"orgcolor,omitempty"`
	ZIndex      uint8  `bson:"zindex,omitempty"`
}
