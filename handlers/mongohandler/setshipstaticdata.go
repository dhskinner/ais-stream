package mongohandler

import (
	"ais-stream/models"
	"log/slog"
	"strings"
	"time"

	ais "github.com/BertoldVdb/go-ais"
	"go.mongodb.org/mongo-driver/bson"
)

func (p *MongoHandler) setShipStaticData(message models.Message) {

	packet, ok := message.Packet.(ais.ShipStaticData)
	if !ok {
		slog.Error("error processing ShipStaticData - could not unqueue packet")
		return
	}

	// check validity
	// shiptype := models.ShipTypeId(packet.Type)
	// if !p.filter.IsWhitelisted(packet.UserID, nil, &shiptype, nil, p.filter.VesselWhitelist) {
	// 	return
	// }

	updates := bson.M{}
	updates["mmsi"] = packet.UserID
	updates["time"] = time.Unix(message.TagBlock.Time, 0)
	if len(packet.Name) > 0 {
		updates["name"] = strings.TrimSpace(packet.Name)
	}
	if len(packet.CallSign) > 0 {
		updates["callsign"] = strings.TrimSpace(packet.CallSign)
	}
	if packet.ImoNumber > 0 {
		updates["imo"] = packet.ImoNumber
	}
	if packet.Type > 0 {
		updates["shiptype"] = models.ShipTypeId(packet.Type)
	}
	if packet.Dimension.A > 0 ||
		packet.Dimension.B > 0 ||
		packet.Dimension.C > 0 ||
		packet.Dimension.D > 0 {
		updates["dimension"] = models.Dimension(packet.Dimension)
	}
	if packet.MaximumStaticDraught > 0 {
		updates["draught"] = models.Metres(packet.MaximumStaticDraught)
	}
	// TODO
	// var dest *models.Destination
	// if len(packet.Destination) > 0 {
	// 	dest = &models.Destination{Destination: strings.TrimSpace(packet.Destination)}
	// }
	// if packet.Eta.Day > 0 &&
	// 	packet.Eta.Month > 0 {
	// 	if dest == nil {
	// 		dest = &models.Destination{}
	// 	}
	// 	eta := models.ETA{
	// 		Month:  packet.Eta.Month,
	// 		Day:    packet.Eta.Day,
	// 		Hour:   packet.Eta.Hour,
	// 		Minute: packet.Eta.Minute,
	// 	}
	// 	dest.ETA = *eta.AsTime()
	// }
	// if dest != nil {
	// 	dest.Time = time.Unix(message.TagBlock.Time, 0)
	// 	updates["dest"] = dest
	// }
	if len(packet.Destination) > 0 {
		updates["dest"] = strings.TrimSpace(packet.Destination)
		eta := models.ETA{
			Month:  packet.Eta.Month,
			Day:    packet.Eta.Day,
			Hour:   packet.Eta.Hour,
			Minute: packet.Eta.Minute,
		}
		updates["eta"] = *eta.AsTime()
	}
	updates["source"] = message.TagBlock.Source
	updates["class"] = models.AisClassA

	p.Upsert(packet.UserID, vesselsCollection, updates)

}
