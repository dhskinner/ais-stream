package models

//	"style": {
//	  "iconColor": "red",
//	  "trackColor": "red",
//	  "fleetColor": "gold",
//	  "orgColor": "gold",
//	  "zIndex": 150
//	}
type Style struct {
	VesselColor string `bson:"vesselcolor,omitempty"`
	FleetColor  string `bson:"fleetcolor,omitempty"`
	OrgColor    string `bson:"orgcolor,omitempty"`
	ZIndex      uint8  `bson:"zindex,omitempty"`
}
