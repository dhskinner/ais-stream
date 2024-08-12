package encode

import (
	"ais-stream/models"

	ais "github.com/BertoldVdb/go-ais"
	"github.com/aisstream/ais-message-models/golang/aisStream"
)

func (e *Encoder) PositionReportClassB(p aisStream.StandardClassBPositionReport) models.Message {

	// create a packet for re-encoding aisstream as go-ais
	packet := ais.StandardClassBPositionReport{
		Header: ais.Header{
			MessageID:       uint8(p.MessageID),
			RepeatIndicator: uint8(p.RepeatIndicator),
			UserID:          uint32(p.UserID),
		},
		Valid:            p.Valid,
		Spare1:           uint8(p.Spare1),
		Sog:              ais.Field10(p.Sog),
		PositionAccuracy: p.PositionAccuracy,
		Longitude:        ais.FieldLatLonFine(p.Longitude),
		Latitude:         ais.FieldLatLonFine(p.Latitude),
		Cog:              ais.Field10(p.Cog),
		TrueHeading:      uint16(p.TrueHeading),
		Timestamp:        uint8(p.Timestamp),
		Spare2:           uint8(p.Spare2),
		ClassBUnit:       p.ClassBUnit,
		ClassBDisplay:    p.ClassBDisplay,
		ClassBDsc:        p.ClassBDsc,
		ClassBBand:       p.ClassBBand,
		ClassBMsg22:      p.ClassBMsg22,
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
