package encode

import (
	"ais-stream/models"

	ais "github.com/BertoldVdb/go-ais"
	"github.com/aisstream/ais-message-models/golang/aisStream"
)

func (e *Encoder) SafetyBroadcastMessage(p aisStream.SafetyBroadcastMessage) models.Message {

	// create a packet for re-encoding aisstream as go-ais
	packet := ais.SafetyBroadcastMessage{
		Header: ais.Header{
			MessageID:       uint8(p.MessageID),
			RepeatIndicator: uint8(p.RepeatIndicator),
			UserID:          uint32(p.UserID),
		},
		Valid: p.Valid,
		Spare: uint8(p.Spare),
		Text:  p.Text,
	}

	// re-encode as a go-ais message
	return e.AsMessage(packet)

}
