package mongohandler

import (
	"ais-stream/models"
	"log/slog"
	"time"

	ais "github.com/BertoldVdb/go-ais"
	"go.mongodb.org/mongo-driver/bson"
)

func (p *MongoHandler) setSafetyBroadcastMessage(message models.Message) {

	packet, ok := message.Packet.(ais.SafetyBroadcastMessage)
	if !ok {
		slog.Error("error processing SafetyBroadcastMessage - could not unqueue packet")
		return
	}

	updates := bson.M{}
	updates["mmsi"] = packet.UserID
	updates["time"] = time.Unix(message.TagBlock.Time, 0)
	updates["text"] = packet.Text
	updates["source"] = message.TagBlock.Source

	p.Upsert(packet.UserID, safetyCollection, updates)

}
