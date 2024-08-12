package encode

import (
	"ais-stream/models"

	ais "github.com/BertoldVdb/go-ais"
	"github.com/aisstream/ais-message-models/golang/aisStream"
)

func (e *Encoder) ShipStaticData(p aisStream.ShipStaticData) models.Message {

	// create a packet for re-encoding aisstream as go-ais
	packet := ais.ShipStaticData{
		Header: ais.Header{
			MessageID:       uint8(p.MessageID),
			RepeatIndicator: uint8(p.RepeatIndicator),
			UserID:          uint32(p.UserID),
		},
		Valid:      p.Valid,
		AisVersion: uint8(p.AisVersion),
		ImoNumber:  uint32(p.ImoNumber),
		CallSign:   p.CallSign,
		Name:       p.Name,
		Type:       uint8(p.Type),
		Dimension: ais.FieldDimension{
			A: uint16(p.Dimension.A),
			B: uint16(p.Dimension.B),
			C: uint8(p.Dimension.C),
			D: uint8(p.Dimension.D),
		},
		FixType: uint8(p.FixType),
		Eta: ais.FieldETA{
			Month:  uint8(p.Eta.Month),
			Day:    uint8(p.Eta.Day),
			Hour:   uint8(p.Eta.Hour),
			Minute: uint8(p.Eta.Minute),
		},
		MaximumStaticDraught: ais.Field10(p.MaximumStaticDraught),
		Destination:          p.Destination,
		Dte:                  p.Dte,
		Spare:                p.Spare,
	}

	// re-encode as a go-ais message
	return e.AsMessage(packet)

}
