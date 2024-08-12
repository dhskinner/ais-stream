package models

import "time"

// ETA represents the encoding of the estimated time of arrival
type ETA struct {
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
}

func (e *ETA) AsTime() *time.Time {
	now := time.Now()
	year := now.UTC().Year()
	if e.Month < uint8(now.Month()) {
		year += 1
	}
	eta := time.Date(year, time.Month(e.Month), int(e.Day), int(e.Hour), int(e.Minute), 0, 0, time.UTC)
	return &eta
}
