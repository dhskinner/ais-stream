package models

type AtonId uint8

const (
	AtonUnknown                         AtonId = 0
	AtonReferencePoint                  AtonId = 1
	AtonRacon                           AtonId = 2
	AtonFixedStructure                  AtonId = 3
	AtonEmergencyWreckMarkingBuoy       AtonId = 4
	AtonLightWithoutSectors             AtonId = 5
	AtonLightWithSectors                AtonId = 6
	AtonLeadingLightFront               AtonId = 7
	AtonLeadingLightRear                AtonId = 8
	AtonBeaconCardinalN                 AtonId = 9
	AtonBeaconCardinalE                 AtonId = 10
	AtonBeaconCardinalS                 AtonId = 11
	AtonBeaconCardinalW                 AtonId = 12
	AtonBeaconPort                      AtonId = 13
	AtonBeaconStarboard                 AtonId = 14
	AtonBeaconPreferredChannelPort      AtonId = 15
	AtonBeaconPreferredChannelStarboard AtonId = 16
	AtonBeaconIsolatedDanger            AtonId = 17
	AtonBeaconSafeWater                 AtonId = 18
	AtonBeaconSpecialMark               AtonId = 19
	AtonCardinalMarkN                   AtonId = 20
	AtonCardinalMarkE                   AtonId = 21
	AtonCardinalMarkS                   AtonId = 22
	AtonCardinalMarkW                   AtonId = 23
	AtonPortMark                        AtonId = 24
	AtonStarboardMark                   AtonId = 25
	AtonPreferredChannelPort            AtonId = 26
	AtonPreferredChannelStarboard       AtonId = 27
	AtonIsolatedDanger                  AtonId = 28
	AtonSafeWater                       AtonId = 29
	AtonSpecialMark                     AtonId = 30
	AtonLightVesselLANBYRigs            AtonId = 31
)

func (n AtonId) AsString() string {
	if n > 31 {
		return atonLabel[0]
	}
	return atonLabel[n]
}

var atonLabel []string = []string{

	"Not specified",
	"Reference point",
	"RACON",
	"Fixed structures off-shore, such as oil platforms, wind farms",
	"Emergency Wreck Marking Buoy",
	// Fixed
	"Light, without sectors",
	"Light, with sectors",
	"Leading Light Front",
	"Leading Light Rear",
	"Beacon, Cardinal N",
	"Beacon, Cardinal E",
	"Beacon, Cardinal S",
	"Beacon, Cardinal W",
	"Beacon, Port hand",
	"Beacon, Starboard hand",
	"Beacon, Preferred Channel port hand",
	"Beacon, Preferred Channel starboard hand",
	"Beacon, Isolated danger",
	"Beacon, Safe water",
	"Beacon, Special mark",
	// Floating
	"Cardinal Mark N",
	"Cardinal Mark E",
	"Cardinal Mark S",
	"Cardinal Mark W",
	"Port hand Mark",
	"Starboard hand Mark",
	"Preferred Channel Port hand",
	"Preferred Channel Starboard hand",
	"Isolated danger",
	"Safe Water",
	"Special Mark",
	"Light Vessel/LANBY/Rigs",
}
