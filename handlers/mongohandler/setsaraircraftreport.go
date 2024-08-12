package mongohandler

import (
	"ais-stream/models"
	"log/slog"
	"time"

	ais "github.com/BertoldVdb/go-ais"
	"go.mongodb.org/mongo-driver/bson"
)

func (p *MongoHandler) setSearchAndRescueAircraftReport(message models.Message) {

	packet, ok := message.Packet.(ais.StandardSearchAndRescueAircraftReport)
	if !ok {
		slog.Error("error processing StandardSearchAndRescueAircraftReport - could not unqueue packet")
		return
	}

	doc := &models.VesselPosition{
		Time:             time.Unix(message.TagBlock.Time, 0),
		Mmsi:             packet.UserID,
		Position:         models.NewCoordinates(float32(packet.Latitude), float32(packet.Longitude)),
		SpeedOverGround:  float32(packet.Sog),
		CourseOverGround: float32(packet.Cog),
		Source:           message.TagBlock.Source,
	}
	if packet.Altitude <= 4094 {
		doc.Altitude = packet.Altitude
	}
	p.InsertPosition(doc)

	// // check validity
	// pos := models.NewCoordinates(float32(packet.Latitude), float32(packet.Longitude))
	// if !p.filter.IsWhitelisted(packet.UserID, nil, nil, &pos, p.filter.VesselWhitelist) {
	// 	return
	// }

	updates := bson.M{}
	updates["mmsi"] = packet.UserID
	updates["time"] = time.Unix(message.TagBlock.Time, 0)
	updates["source"] = message.TagBlock.Source
	updates["class"] = models.AisAircraft
	pos := models.NewCoordinates(float32(packet.Latitude), float32(packet.Longitude))
	if pos.IsValid() {
		updates["pos"] = pos
		updates["state"] = pos.AsState()
	}
	if packet.Altitude <= 4094 {
		updates["alt"] = packet.Altitude
	}
	if packet.Cog >= 0 && packet.Cog < 360 {
		updates["cog"] = packet.Cog
	} else {
		updates["cog"] = 0
	}
	if packet.Sog > 0 {
		updates["sog"] = packet.Sog
	} else {
		updates["sog"] = 0
	}
	p.Upsert(packet.UserID, vesselsCollection, updates)

}
