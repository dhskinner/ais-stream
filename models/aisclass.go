package models

type AisClass string

const (
	AisClassUnknown AisClass = "Unknown"
	AisClassA       AisClass = "A"
	AisClassB       AisClass = "B"
	AisAtoN         AisClass = "AtoN"
	AisBaseStation  AisClass = "Base"
	AisAircraft     AisClass = "Aircraft"
)
