package models

type ShipTypeId int32

const (
	ShipTypeUnknown              ShipTypeId = 0
	ShipTypeReserved1            ShipTypeId = 1
	ShipTypeReserved2            ShipTypeId = 2
	ShipTypeReserved3            ShipTypeId = 3
	ShipTypeReserved4            ShipTypeId = 4
	ShipTypeReserved5            ShipTypeId = 5
	ShipTypeReserved6            ShipTypeId = 6
	ShipTypeReserved7            ShipTypeId = 7
	ShipTypeReserved8            ShipTypeId = 8
	ShipTypeReserved9            ShipTypeId = 9
	ShipTypeReserved10           ShipTypeId = 10
	ShipTypeReserved11           ShipTypeId = 11
	ShipTypeReserved12           ShipTypeId = 12
	ShipTypeReserved13           ShipTypeId = 13
	ShipTypeReserved14           ShipTypeId = 14
	ShipTypeReserved15           ShipTypeId = 15
	ShipTypeReserved16           ShipTypeId = 16
	ShipTypeReserved17           ShipTypeId = 17
	ShipTypeReserved18           ShipTypeId = 18
	ShipTypeReserved19           ShipTypeId = 19
	ShipTypeWingInGround         ShipTypeId = 20
	ShipTypeWingInGroundA        ShipTypeId = 21
	ShipTypeWingInGroundB        ShipTypeId = 22
	ShipTypeWingInGroundC        ShipTypeId = 23
	ShipTypeWingInGroundD        ShipTypeId = 24
	ShipTypeReserved25           ShipTypeId = 25
	ShipTypeReserved26           ShipTypeId = 26
	ShipTypeReserved27           ShipTypeId = 27
	ShipTypeReserved28           ShipTypeId = 28
	ShipTypeReserved29           ShipTypeId = 29
	ShipTypeFishing              ShipTypeId = 30
	ShipTypeTowing               ShipTypeId = 31
	ShipTypeTowingLong           ShipTypeId = 32
	ShipTypeDredging             ShipTypeId = 33
	ShipTypeDiving               ShipTypeId = 34
	ShipTypeMilitary             ShipTypeId = 35
	ShipTypeSailing              ShipTypeId = 36
	ShipTypePleasure             ShipTypeId = 37
	ShipTypeReserved38           ShipTypeId = 38
	ShipTypeReserved39           ShipTypeId = 39
	ShipTypeHighSpeedCraft       ShipTypeId = 40
	ShipTypeHighSpeedCraftA      ShipTypeId = 41
	ShipTypeHighSpeedCraftB      ShipTypeId = 42
	ShipTypeHighSpeedCraftC      ShipTypeId = 43
	ShipTypeHighSpeedCraftD      ShipTypeId = 44
	ShipTypeReserved45           ShipTypeId = 45
	ShipTypeReserved46           ShipTypeId = 46
	ShipTypeReserved47           ShipTypeId = 47
	ShipTypeReserved48           ShipTypeId = 48
	ShipTypeHighSpeedCraftNoInfo ShipTypeId = 49
	ShipTypePilot                ShipTypeId = 50
	ShipTypeSAR                  ShipTypeId = 51
	ShipTypeTug                  ShipTypeId = 52
	ShipTypePortTender           ShipTypeId = 53
	ShipTypeAntiPollution        ShipTypeId = 54
	ShipTypeLawEnforcement       ShipTypeId = 55
	ShipTypeReserved56           ShipTypeId = 56
	ShipTypeReserved57           ShipTypeId = 57
	ShipTypeMedicalTransport     ShipTypeId = 58
	ShipTypeNoncombatant         ShipTypeId = 59
	ShipTypePassenger            ShipTypeId = 60
	ShipTypePassengerA           ShipTypeId = 61
	ShipTypePassengerB           ShipTypeId = 62
	ShipTypePassengerC           ShipTypeId = 63
	ShipTypePassengerD           ShipTypeId = 64
	ShipTypeReserved65           ShipTypeId = 65
	ShipTypeReserved66           ShipTypeId = 66
	ShipTypeReserved67           ShipTypeId = 67
	ShipTypeReserved68           ShipTypeId = 68
	ShipTypePassengerNoInfo      ShipTypeId = 69
	ShipTypeCargo                ShipTypeId = 70
	ShipTypeCargoA               ShipTypeId = 71
	ShipTypeCargoB               ShipTypeId = 72
	ShipTypeCargoC               ShipTypeId = 73
	ShipTypeCargoD               ShipTypeId = 74
	ShipTypeReserved75           ShipTypeId = 75
	ShipTypeReserved76           ShipTypeId = 76
	ShipTypeReserved77           ShipTypeId = 77
	ShipTypeReserved78           ShipTypeId = 78
	ShipTypeCargoNoInfo          ShipTypeId = 79
	ShipTypeTanker               ShipTypeId = 80
	ShipTypeTankerA              ShipTypeId = 81
	ShipTypeTankerB              ShipTypeId = 82
	ShipTypeTankerC              ShipTypeId = 83
	ShipTypeTankerD              ShipTypeId = 84
	ShipTypeReserved85           ShipTypeId = 85
	ShipTypeReserved86           ShipTypeId = 86
	ShipTypeReserved87           ShipTypeId = 87
	ShipTypeReserved88           ShipTypeId = 88
	ShipTypeTankerNoInfo         ShipTypeId = 89
	ShipTypeOther                ShipTypeId = 90
	ShipTypeOtherA               ShipTypeId = 91
	ShipTypeOtherB               ShipTypeId = 92
	ShipTypeOtherC               ShipTypeId = 93
	ShipTypeOtherD               ShipTypeId = 94
	ShipTypeReserved95           ShipTypeId = 95
	ShipTypeReserved96           ShipTypeId = 96
	ShipTypeReserved97           ShipTypeId = 97
	ShipTypeReserved98           ShipTypeId = 98
	ShipTypeOtherNoInfo          ShipTypeId = 99
)

func (s ShipTypeId) AsString() string {
	if s > 99 {
		return shipTypeLabel[0]
	}
	return shipTypeLabel[s]
}

func (s ShipTypeId) AsCategory() ShipCategoryId {

	switch i := uint8(s); i {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
		return ShipCategoryReserved
	case 20, 21, 22, 23, 24, 25, 26, 27, 28, 29:
		return ShipCategoryWingInGround
	case 30, 31, 32, 33, 34, 35, 36, 37, 38, 39:
		return ShipCategorySpecial1
	case 40, 41, 42, 43, 44, 45, 46, 47, 48, 49:
		return ShipCategoryHighSpeedCraft
	case 50, 51, 52, 53, 54, 55, 56, 57, 58, 59:
		return ShipCategorySpecial2
	case 60, 61, 62, 63, 64, 65, 66, 67, 68, 69:
		return ShipCategoryPassenger
	case 70, 71, 72, 73, 74, 75, 76, 77, 78, 79:
		return ShipCategoryCargo
	case 80, 81, 82, 83, 84, 85, 86, 87, 88, 89:
		return ShipCategoryTanker
	case 90, 91, 92, 93, 94, 95, 96, 97, 98, 99:
		return ShipCategoryOther
	default:
		return ShipCategoryUnknown
	}

}

var shipTypeLabel []string = []string{
	"Not available",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Reserved",
	"Wing in ground (WIG), all ships of this type",
	"Wing in ground (WIG), Hazardous category A",
	"Wing in ground (WIG), Hazardous category B",
	"Wing in ground (WIG), Hazardous category C",
	"Wing in ground (WIG), Hazardous category D",
	"Wing in ground (WIG), Reserved for future use",
	"Wing in ground (WIG), Reserved for future use",
	"Wing in ground (WIG), Reserved for future use",
	"Wing in ground (WIG), Reserved for future use",
	"Wing in ground (WIG), Reserved for future use",
	"Fishing",
	"Towing",
	"Towing: length exceeds 200m or breadth exceeds 25m",
	"Dredging or underwater ops",
	"Diving ops",
	"Military ops",
	"Sailing",
	"Pleasure Craft",
	"Reserved",
	"Reserved",
	"High speed craft (HSC), all ships of this type",
	"High speed craft (HSC), Hazardous category A",
	"High speed craft (HSC), Hazardous category B",
	"High speed craft (HSC), Hazardous category C",
	"High speed craft (HSC), Hazardous category D",
	"High speed craft (HSC), Reserved for future use",
	"High speed craft (HSC), Reserved for future use",
	"High speed craft (HSC), Reserved for future use",
	"High speed craft (HSC), Reserved for future use",
	"High speed craft (HSC), No additional information",
	"Pilot Vessel",
	"Search and Rescue vessel",
	"Tug",
	"Port Tender",
	"Anti-pollution equipment",
	"Law Enforcement",
	"Spare - Local Vessel",
	"Spare - Local Vessel",
	"Medical Transport",
	"Noncombatant ship according to RR Resolution No. 18",
	"Passenger, all ships of this type",
	"Passenger, Hazardous category A",
	"Passenger, Hazardous category B",
	"Passenger, Hazardous category C",
	"Passenger, Hazardous category D",
	"Passenger, Reserved for future use",
	"Passenger, Reserved for future use",
	"Passenger, Reserved for future use",
	"Passenger, Reserved for future use",
	"Passenger, No additional information",
	"Cargo, all ships of this type",
	"Cargo, Hazardous category A",
	"Cargo, Hazardous category B",
	"Cargo, Hazardous category C",
	"Cargo, Hazardous category D",
	"Cargo, Reserved for future use",
	"Cargo, Reserved for future use",
	"Cargo, Reserved for future use",
	"Cargo, Reserved for future use",
	"Cargo, No additional information",
	"Tanker, all ships of this type",
	"Tanker, Hazardous category A",
	"Tanker, Hazardous category B",
	"Tanker, Hazardous category C",
	"Tanker, Hazardous category D",
	"Tanker, Reserved for future use",
	"Tanker, Reserved for future use",
	"Tanker, Reserved for future use",
	"Tanker, Reserved for future use",
	"Tanker, No additional information",
	"Other Type, all ships of this type",
	"Other Type, Hazardous category A",
	"Other Type, Hazardous category B",
	"Other Type, Hazardous category C",
	"Other Type, Hazardous category D",
	"Other Type, Reserved for future use",
	"Other Type, Reserved for future use",
	"Other Type, Reserved for future use",
	"Other Type, Reserved for future use",
	"Other Type, no additional information",
}

var SpecialCraft = []ShipTypeId{
	ShipTypeTowing,
	ShipTypeTowingLong,
	ShipTypeDredging,
	ShipTypeDiving,
	ShipTypeMilitary,
	ShipTypeHighSpeedCraft,
	ShipTypeHighSpeedCraftA,
	ShipTypeHighSpeedCraftB,
	ShipTypeHighSpeedCraftC,
	ShipTypeHighSpeedCraftD,
	ShipTypeHighSpeedCraftNoInfo,
	ShipTypePilot,
	ShipTypeSAR,
	// ShipTypeTug,
	// ShipTypePortTender,
	ShipTypeAntiPollution,
	ShipTypeLawEnforcement,
	ShipTypeMedicalTransport,
}
