package encode

import (
	"ais-stream/models"

	ais "github.com/BertoldVdb/go-ais"
	"github.com/aisstream/ais-message-models/golang/aisStream"
)

func (e *Encoder) LongRangeAisBroadcastMessage(p aisStream.LongRangeAisBroadcastMessage) models.Message {

	// create a packet for re-encoding aisstream as go-ais
	packet := ais.LongRangeAisBroadcastMessage{
		Header: ais.Header{
			MessageID:       uint8(p.MessageID),
			RepeatIndicator: uint8(p.RepeatIndicator),
			UserID:          uint32(p.UserID),
		},
		Valid:              p.Valid,
		PositionAccuracy:   p.PositionAccuracy,
		Raim:               p.Raim,
		NavigationalStatus: uint8(p.NavigationalStatus),
		Longitude:          ais.FieldLatLonCoarse(p.Longitude),
		Latitude:           ais.FieldLatLonCoarse(p.Latitude),
		Sog:                uint8(p.Sog),
		Cog:                uint16(p.Cog),
		PositionLatency:    p.PositionLatency,
		Spare:              p.Spare,
	}

	// re-encode as a go-ais message
	return e.AsMessage(packet)

}
