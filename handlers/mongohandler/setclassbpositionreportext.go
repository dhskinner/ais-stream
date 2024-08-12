package mongohandler

import (
	"ais-stream/models"
	"log/slog"
	"strings"
	"time"

	ais "github.com/BertoldVdb/go-ais"
	"go.mongodb.org/mongo-driver/bson"
)

func (p *MongoHandler) setExtendedClassBPositionReport(message models.Message) {

	packet, ok := message.Packet.(ais.ExtendedClassBPositionReport)
	if !ok {
		slog.Error("error processing ExtendedClassBPositionReport - could not unqueue packet")
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
	if len(packet.Name) > 0 {
		updates["name"] = strings.TrimSpace(packet.Name)
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
	updates["class"] = models.AisClassB

	p.Upsert(packet.UserID, vesselsCollection, updates)

}
