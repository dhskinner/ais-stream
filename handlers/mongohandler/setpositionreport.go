package mongohandler

import (
	"ais-stream/models"
	"log/slog"
	"time"

	ais "github.com/BertoldVdb/go-ais"
	"go.mongodb.org/mongo-driver/bson"
)

func (p *MongoHandler) setPositionReport(message models.Message) {

	packet, ok := message.Packet.(ais.PositionReport)
	if !ok {
		slog.Error("error processing PositionReport - could not unqueue packet")
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
	p.InsertPosition(doc)

	updates := bson.M{}
	updates["mmsi"] = packet.UserID
	updates["time"] = time.Unix(message.TagBlock.Time, 0)
	pos := models.NewCoordinates(float32(packet.Latitude), float32(packet.Longitude))
	if pos.IsValid() {
		updates["pos"] = pos
		updates["state"] = pos.AsState()
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
	updates["source"] = message.TagBlock.Source
	updates["nav"] = models.NavigationId(packet.NavigationalStatus)
	p.Upsert(packet.UserID, vesselsCollection, updates)

}
