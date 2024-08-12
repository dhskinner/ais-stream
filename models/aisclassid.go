package models

type AisClassId uint8

const (
	AisClassUnknown AisClassId = 0
	AisClassA       AisClassId = 1
	AisClassB       AisClassId = 2
	AisAtoN         AisClassId = 3
	AisStation      AisClassId = 4
	AisAircraft     AisClassId = 5
)

func (s AisClassId) AsString() string {

	switch AisClassId(s) {
	case AisClassA:
		return "A"
	case AisClassB:
		return "B"
	case AisAtoN:
		return "Aid to Navigation"
	case AisStation:
		return "Base Station"
	case AisAircraft:
		return "Aircraft"
	default:
		return "Unknown"
	}

}

func AisClassIdFromString(in string) AisClassId {

	switch in {
	case "A":
		return AisClassA
	case "B":
		return AisClassB
	case "AtoN":
		fallthrough
	case "Aid to Navigation":
		return AisAtoN
	case "Base Station":
		fallthrough
	case "Base":
		return AisStation
	case "Aircraft":
		return AisAircraft
	default:
		return AisClassUnknown
	}

}
