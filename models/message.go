package models

import (
	"time"

	"github.com/BertoldVdb/go-ais/aisnmea"
)

type Message = *aisnmea.VdmPacket

func Timestamp(m Message) time.Time {

	t1 := time.Now()
	t2 := time.Unix(m.TagBlock.Time, 0)

	// sanity check
	if m.TagBlock.Time == 0 || t2.After(t1) {
		return t1
	}

	return t2

}
