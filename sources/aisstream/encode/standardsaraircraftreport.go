package encode

import (
	"ais-stream/models"

	ais "github.com/BertoldVdb/go-ais"
	"github.com/aisstream/ais-message-models/golang/aisStream"
)

func (e *Encoder) StandardSarAircraftReport(p aisStream.StandardSearchAndRescueAircraftReport) models.Message {

	// create a packet for re-encoding aisstream as go-ais
	packet := ais.StandardSearchAndRescueAircraftReport{
		Header: ais.Header{
			MessageID:       uint8(p.MessageID),
			RepeatIndicator: uint8(p.RepeatIndicator),
			UserID:          uint32(p.UserID),
		},
		Valid:            p.Valid,
		Altitude:         uint16(p.Altitude),
		Sog:              uint16(p.Sog),
		PositionAccuracy: p.PositionAccuracy,
		Longitude:        ais.FieldLatLonFine(p.Longitude),
		Latitude:         ais.FieldLatLonFine(p.Latitude),
		Cog:              ais.Field10(p.Cog),
		Timestamp:        uint8(p.Timestamp),
		AltFromBaro:      p.AltFromBaro,
		Spare1:           uint8(p.Spare1),
		Dte:              p.Dte,
		Spare2:           uint8(p.Spare2),
		AssignedMode:     p.AssignedMode,
		Raim:             p.Raim,
		CommunicationStateItdma: ais.CommunicationStateItdma{
			CommunicationStateIsItdma: p.CommunicationStateIsItdma,
			CommunicationState:        uint32(p.CommunicationState),
		},
	}

	// re-encode as a go-ais message
	return e.AsMessage(packet)

}
