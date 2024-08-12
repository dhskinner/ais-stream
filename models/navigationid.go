package models

type NavigationId uint8

const (
	NavigationUnderway             NavigationId = 0
	Navigationanchor               NavigationId = 1
	NavigationNUC                  NavigationId = 2
	NavigationRAM                  NavigationId = 3
	NavigationConstrainedByDraught NavigationId = 4
	NavigationMoored               NavigationId = 5
	NavigationAground              NavigationId = 6
	NavigationFishing              NavigationId = 7
	NavigationSailing              NavigationId = 8
	NavigationReserved1            NavigationId = 9
	Navigationreserved2            NavigationId = 10
	NavigationTowingAstern         NavigationId = 11
	NavigationPushingAhead         NavigationId = 12
	NavigationReserved3            NavigationId = 13
	NavigationAIS                  NavigationId = 14
	NavigationUndefined            NavigationId = 15
)

func (n NavigationId) AsShortString() string {
	return navigationLabelShort[n]
}

func (n NavigationId) AsLongString() string {
	return navigationLabelLong[n]
}

var navigationLabelShort []string = []string{
	"Underway",
	"Anchored",
	"Not under command",
	"RAM",
	"Draught constrained",
	"Moored",
	"Aground",
	"Fishing",
	"Sailing",
	"Reserved",
	"Reserved",
	"Towing",
	"Pushing",
	"Reserved",
	"SART, MOB, EPIRB",
	"Undefined",
}

var navigationLabelLong []string = []string{
	"Under way using engine",
	"At anchor",
	"Not under command",
	"Restricted maneuverability",
	"Constrained by her draught",
	"Moored",
	"Aground",
	"Engaged in fishing",
	"Under way sailing",
	"Reserved for future amendment of navigational status for ships carrying DG, HS, or MP, or IMO hazard or pollutant category C, high speed craft (HSC)",
	"Reserved for future amendment of navigational status for ships carrying dangerous goods (DG), harmful substances (HS) or marine pollutants (MP), or IMO hazard or pollutant category A, wing in ground (WIG)",
	"Power-driven vessel towing astern (regional use)",
	"Power-driven vessel pushing ahead or towing alongside (regional use)",
	"Reserved for future use",
	"AIS-SART (active), MOB-AIS, EPIRB-AIS",
	"Undefined (also used by AIS-SART, MOB-AIS and EPIRB-AIS under test)",
}
