package models

import "time"

type VesselInfo struct {
	Time                 time.Time    `bson:"time"`                         // ok
	Mmsi                 MMSI         `bson:"mmsi"`                         // ok
	Name                 string       `bson:"name,omitempty"`               // ok
	CallSign             string       `bson:"callsign,omitempty"`           // ok
	ImoNumber            uint32       `bson:"imo,omitempty"`                // ok
	Fleet                string       `bson:"fleet,omitempty"`              // not populated
	HomePort             Coordinates  `bson:"homeport,omitempty"`           // not populated
	Organisation         string       `bson:"org,omitempty"`                // not populated
	Style                Style        `bson:"style,omitempty"`              // not populated
	ShipType             ShipTypeId   `bson:"shiptype,omitempty,truncate"`  // ok
	Dimension            Dimension    `bson:"dimension,omitempty,truncate"` // ok
	MaximumStaticDraught Metres       `bson:"draught,omitempty,truncate"`   // ok
	AisClass             AisClassId   `bson:"class,omitempty"`              // ok
	Vendor               VendorInfo   `bson:"vendor,omitempty"`             // ok
	Position             Coordinates  `bson:"pos,omitempty,truncate"`       // ok
	Altitude             Metres       `bson:"alt,omitempty,truncate"`       // ok
	SpeedOverGround      Knots        `bson:"sog,omitempty,truncate"`       // ok
	CourseOverGround     Degrees      `bson:"cog,omitempty,truncate"`       // ok
	Source               string       `bson:"source,omitempty"`             // ok
	State                string       `bson:"state,omitempty"`              // ok
	NavigationStatus     NavigationId `bson:"nav,omitempty"`                // ok
	Destination          string       `bson:"dest,omitempty"`               // ok
	ETA                  time.Time    `bson:"eta,omitempty"`                // ok
}

func (v *VesselInfo) IsValid() bool {

	if v.Mmsi == 0 ||
		v.Mmsi > 999999999 ||
		v.Time.IsZero() {
		return false
	}

	return true
}

// {
//   "mmsi": 503372200,
//   "name": "CG24",
//   "fleet": "QF2 Brisbane",
//   "home_port": {
//     "type": "Point",
//     "coordinates": [
//       153.189728,
//       -27.452103
//     ]
//   },
//   "org": "QF2",
//   "style": {
//     "iconColor": "red",
//     "trackColor": "red",
//     "fleetColor": "gold",
//     "orgColor": "gold",
//     "zIndex": 150
//   }
// }

// {
//   "_id": {
//     "$oid": "661df6672884b4ba2d55611c"
//   },
//   "mmsi": 503466000,
//   "time": {
//     "$date": "2024-04-16T03:54:15.000Z"
//   },
//   "unix": {
//     "$numberLong": "1713239655"
//   },
//   "src": "Cape Jervis B",
//   "state": "SA",
//   "tags": [
//     "amsa"
//   ],
//   "nmea": "\\s:Cape Jervis B,c:1713239655*72\\!AIVDM,1,1,2,B,17P9840tj89p6`ccWkcprW0L0D03,0*7C",
//   "msg": 1,
//   "cog": 228.1999969482422,
//   "sog": 13.600000381469727,
//   "pos": {
//     "type": "Point",
//     "coordinates": [
//       138.085155,
//       -35.613148333333335
//     ]
//   }
// }
