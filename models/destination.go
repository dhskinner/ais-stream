package models

import "time"

type Destination struct {
	Time        time.Time `bson:"time"`           // ok
	Destination string    `bson:"dest,omitempty"` // ok
	ETA         time.Time `bson:"eta,omitempty"`  // ok
}
