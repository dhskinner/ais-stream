package mongohandler

import (
	"ais-stream/models"
	"log/slog"
	"strings"
	"time"

	ais "github.com/BertoldVdb/go-ais"
	"go.mongodb.org/mongo-driver/bson"
)

func (p *MongoHandler) setAidsToNavigationReport(message models.Message) {

	packet, ok := message.Packet.(ais.AidsToNavigationReport)
	if !ok {
		slog.Error("error processing AidsToNavigationReport - could not unqueue packet")
		return
	}

	// check validity
	// pos := models.NewCoordinates(float32(packet.Latitude), float32(packet.Longitude))
	// if !p.filter.IsWhitelisted(packet.UserID, nil, nil, &pos, p.filter.VesselWhitelist) {
	// 	return
	// }

	updates := bson.M{}
	updates["mmsi"] = packet.UserID
	updates["time"] = time.Unix(message.TagBlock.Time, 0)

	if len(packet.Name) > 0 {
		updates["name"] = strings.TrimSpace(packet.Name) + strings.TrimSpace(packet.NameExtension)
	}
	pos := models.NewCoordinates(float32(packet.Latitude), float32(packet.Longitude))
	if pos.IsValid() {
		updates["pos"] = pos
		updates["state"] = pos.AsState()
	}
	if packet.Type > 0 {
		updates["aton.id"] = models.AtonId(packet.Type)
	}
	if packet.OffPosition {
		updates["aton.offpos"] = packet.OffPosition
	}
	if packet.VirtualAtoN {
		updates["aton.virtual"] = packet.VirtualAtoN
	}
	if packet.Dimension.A > 0 ||
		packet.Dimension.B > 0 ||
		packet.Dimension.C > 0 ||
		packet.Dimension.D > 0 {
		updates["dimension"] = models.Dimension(packet.Dimension)
	}
	updates["class"] = models.AisAtoN
	updates["source"] = message.TagBlock.Source

	p.Upsert(packet.UserID, atonCollection, updates)

}
