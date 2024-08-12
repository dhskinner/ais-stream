package encode

import (
	"ais-stream/models"

	ais "github.com/BertoldVdb/go-ais"
	"github.com/aisstream/ais-message-models/golang/aisStream"
)

func (e *Encoder) StaticDataReport(p aisStream.StaticDataReport) models.Message {

	// create a packet for re-encoding aisstream as go-ais
	packet := ais.StaticDataReport{
		Header: ais.Header{
			MessageID:       uint8(p.MessageID),
			RepeatIndicator: uint8(p.RepeatIndicator),
			UserID:          uint32(p.UserID),
		},
		Valid:      p.Valid,
		Reserved:   uint8(p.Reserved),
		PartNumber: p.PartNumber,
		ReportA: ais.StaticDataReportA{
			Valid: p.ReportA.Valid,
			Name:  p.ReportA.Name,
		},
		ReportB: ais.StaticDataReportB{
			Valid:          p.ReportB.Valid,
			ShipType:       uint8(p.ReportB.ShipType),
			VendorIDName:   p.ReportB.VendorIDName,
			VenderIDModel:  uint8(p.ReportB.VenderIDModel),
			VenderIDSerial: uint32(p.ReportB.VenderIDSerial),
			CallSign:       p.ReportB.CallSign,
			Dimension: ais.FieldDimension{
				A: uint16(p.ReportB.Dimension.A),
				B: uint16(p.ReportB.Dimension.B),
				C: uint8(p.ReportB.Dimension.C),
				D: uint8(p.ReportB.Dimension.D),
			},
			FixType: uint8(p.ReportB.FixType),
			Spare:   uint8(p.ReportB.Spare),
		},
	}

	// re-encode as a go-ais message
	return e.AsMessage(packet)

}
