package models

type AisClass uint8

const (
	AisClassUnknown AisClass = 0
	AisClassA       AisClass = 1
	AisClassB       AisClass = 2
	AisAtoN         AisClass = 3
	AisStation      AisClass = 4
	AisAircraft     AisClass = 5
)

func (s AisClass) AsString() string {

	switch AisClass(s) {
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

func AisClassIdFromString(in string) AisClass {

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
