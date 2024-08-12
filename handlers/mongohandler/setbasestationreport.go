package mongohandler

import (
	"ais-stream/models"
	"log/slog"
	"time"

	ais "github.com/BertoldVdb/go-ais"
	"go.mongodb.org/mongo-driver/bson"
)

func (p *MongoHandler) setBaseStationReport(message models.Message) {

	packet, ok := message.Packet.(ais.BaseStationReport)
	if !ok {
		slog.Error("error processing BaseStationReport - could not unqueue packet")
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
	pos := models.NewCoordinates(float32(packet.Latitude), float32(packet.Longitude))
	if pos.IsValid() {
		updates["pos"] = pos
		updates["state"] = pos.AsState()
	}
	updates["class"] = models.AisStation
	updates["source"] = message.TagBlock.Source

	p.Upsert(packet.UserID, stationsCollection, updates)

}
