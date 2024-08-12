package encode

import (
	"ais-stream/models"

	ais "github.com/BertoldVdb/go-ais"
	"github.com/aisstream/ais-message-models/golang/aisStream"
)

func (e *Encoder) PositionReportClassA(p aisStream.PositionReport) models.Message {

	// create a packet for re-encoding aisstream as go-ais
	packet := ais.PositionReport{
		Header: ais.Header{
			MessageID:       uint8(p.MessageID),
			RepeatIndicator: uint8(p.RepeatIndicator),
			UserID:          uint32(p.UserID),
		},
		Valid:                     p.Valid,
		NavigationalStatus:        uint8(p.NavigationalStatus),
		RateOfTurn:                int16(p.RateOfTurn),
		Sog:                       ais.Field10(p.Sog),
		PositionAccuracy:          p.PositionAccuracy,
		Longitude:                 ais.FieldLatLonFine(p.Longitude),
		Latitude:                  ais.FieldLatLonFine(p.Latitude),
		Cog:                       ais.Field10(p.Cog),
		TrueHeading:               uint16(p.TrueHeading),
		Timestamp:                 uint8(p.Timestamp),
		SpecialManoeuvreIndicator: uint8(p.SpecialManoeuvreIndicator),
		Spare:                     uint8(p.Spare),
		Raim:                      p.Raim,
		CommunicationStateNoItdma: ais.CommunicationStateNoItdma{
			CommunicationState: uint32(p.CommunicationState),
		},
	}

	// re-encode as a go-ais message
	return e.AsMessage(packet)

}
