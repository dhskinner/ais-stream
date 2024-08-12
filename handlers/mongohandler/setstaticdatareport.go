package mongohandler

import (
	"ais-stream/models"
	"log/slog"
	"strings"
	"time"

	ais "github.com/BertoldVdb/go-ais"
	"go.mongodb.org/mongo-driver/bson"
)

func (p *MongoHandler) setStaticDataReport(message models.Message) {

	packet, ok := message.Packet.(ais.StaticDataReport)
	if !ok {
		slog.Error("error processing StaticDataReport - could not unqueue packet")
		return
	}

	// check validity
	// shiptype := models.ShipTypeId(packet.ReportB.ShipType)
	// if !p.filter.IsWhitelisted(packet.UserID, nil, &shiptype, nil, p.filter.VesselWhitelist) {
	// 	return
	// }

	updates := bson.M{}
	updates["mmsi"] = packet.UserID
	updates["time"] = time.Unix(message.TagBlock.Time, 0)
	if len(packet.ReportA.Name) > 0 {
		updates["name"] = strings.TrimSpace(packet.ReportA.Name)
	}
	if len(packet.ReportB.CallSign) > 0 {
		updates["callsign"] = strings.TrimSpace(packet.ReportB.CallSign)
	}
	if packet.ReportB.ShipType > 0 {
		updates["shiptype"] = models.ShipTypeId(packet.ReportB.ShipType)
	}
	if packet.ReportB.Dimension.A > 0 ||
		packet.ReportB.Dimension.B > 0 ||
		packet.ReportB.Dimension.C > 0 ||
		packet.ReportB.Dimension.D > 0 {
		updates["dimension"] = models.Dimension(packet.ReportB.Dimension)
	}
	if len(packet.ReportB.VendorIDName) > 0 {
		updates["vendor.name"] = strings.TrimSpace(packet.ReportB.VendorIDName)
	}
	if packet.ReportB.VenderIDSerial > 0 {
		updates["vendor.serial"] = packet.ReportB.VenderIDSerial
	}
	if packet.ReportB.VenderIDModel > 0 {
		updates["vendor.model"] = packet.ReportB.VenderIDModel
	}
	updates["source"] = message.TagBlock.Source
	updates["class"] = models.AisClassA

	p.Upsert(packet.UserID, vesselsCollection, updates)

}
