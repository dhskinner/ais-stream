package encode

import (
	"ais-stream/models"

	ais "github.com/BertoldVdb/go-ais"
	"github.com/aisstream/ais-message-models/golang/aisStream"
)

func (e *Encoder) PositionReportClassBExtended(p aisStream.ExtendedClassBPositionReport) models.Message {

	// create a packet for re-encoding aisstream as go-ais
	packet := ais.ExtendedClassBPositionReport{
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
		Name:             p.Name,
		Type:             uint8(p.Type),
		Dimension: ais.FieldDimension{
			A: uint16(p.Dimension.A),
			B: uint16(p.Dimension.B),
			C: uint8(p.Dimension.C),
		},
		FixType:      uint8(p.FixType),
		Raim:         p.Raim,
		Dte:          p.Dte,
		AssignedMode: p.AssignedMode,
		Spare3:       uint8(p.Spare3),
	}

	// re-encode as a go-ais message
	return e.AsMessage(packet)

}
