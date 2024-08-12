package aisstream

import (
	"ais-stream/interfaces"
	"ais-stream/sources"
	"ais-stream/sources/aisstream/encode"
	"context"
	"encoding/json"
	"log/slog"
	"sync"
	"time"

	"github.com/aisstream/ais-message-models/golang/aisStream"
	"github.com/gorilla/websocket"
)

var messageTypes []aisStream.AisMessageTypes = []aisStream.AisMessageTypes{
	"PositionReport",
	"ExtendedClassBPositionReport",
	"LongRangeAisBroadcastMessage",
	"SafetyBroadcastMessage",
	"ShipStaticData",
	"StandardClassBPositionReport",
	"StandardSearchAndRescueAircraftReport",
	"StaticDataReport",
}

// "2024-04-20 04:46:07.219885807 +0000 UTC"
const timeLayout = "2006-01-02 15:04:05.000000000 -0700 UTC"

type aisClient struct {
	ctx     context.Context
	config  *sources.Config
	handler interfaces.Handler
}

func Client(ctx context.Context, wg *sync.WaitGroup, config *sources.Config, hd interfaces.Handler) {

	// tell the caller we've stopped
	defer wg.Done()

	// the following must be present
	retry, err := config.GetRetry()
	if err != nil {
		slog.Error("aissstream: error", "error", err)
		return
	}

	// service the message stream (blocking)
	c := &aisClient{ctx: ctx, config: config, handler: hd}
	for {

		// run a new worker
		err := c.run()
		if err != nil {
			slog.Error("aissstream: error", "error", err)
		}

		// on error (and if not cancelled), automatically restart
		select {
		case <-ctx.Done():
			slog.Info("aisstream: stopped worker")
			return
		default:
			time.Sleep(time.Duration(retry) * time.Second)
		}
	}
}

// setup a new connection to aisstream and service messages
func (c *aisClient) run() error {

	// get some environment variables - the following keys must be present:
	boundary, err := c.config.GetBoundary()
	if err != nil {
		return err
	}
	apiKey, err := c.config.GetApiKey()
	if err != nil {
		return err
	}
	address, err := c.config.GetAddress()
	if err != nil {
		return err
	}

	// connect to the webstream
	ws, _, err := websocket.DefaultDialer.Dial(address, nil)
	if err != nil {
		slog.Info("aisstream: dial failed", "address", address)
		return err
	}
	slog.Info("aisstream: connected", "address", address)

	// defer a function for orderly shutdown
	defer func() {
		ws.Close()
		slog.Info("aisstream: closed webstream")
	}()

	subMsg := aisStream.SubscriptionMessage{
		APIKey:             apiKey,
		BoundingBoxes:      boundary,
		FilterMessageTypes: messageTypes,
	}

	subMsgBytes, _ := json.Marshal(subMsg)
	if err := ws.WriteMessage(websocket.TextMessage, subMsgBytes); err != nil {
		slog.Info("aisstream: subscription failed", "error", err)
		return err
	}

	// create a new encoder to transform aisstream back to nmea sentences
	encoder := encode.NewEncoder()

	slog.Info("aisstream: starting worker")

worker:
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return err
		}
		var packet aisStream.AisStreamMessage

		err = json.Unmarshal(p, &packet)
		if err != nil {
			return err
		}

		// check if there is a valid timestamp
		timestamp := time.Now().Unix()
		metadata, ok := packet.GetMetaDataOk()
		if ok {
			ts := metadata["time_utc"]
			timeStr, ok := ts.(string)
			if ok {
				tm, err := time.Parse(timeLayout, timeStr)
				if err == nil {
					timestamp = tm.Unix()
				}
			}
		}

		switch packet.MessageType {

		// Type 1,2,3
		case aisStream.POSITION_REPORT:
			message := encoder.PositionReportClassA(*packet.Message.PositionReport)
			message.TagBlock.Time = timestamp
			c.handler.Message(message)

		// Type 5
		case aisStream.SHIP_STATIC_DATA:
			message := encoder.ShipStaticData(*packet.Message.ShipStaticData)
			message.TagBlock.Time = timestamp
			c.handler.Message(message)

		// Type 9
		case aisStream.STANDARD_SEARCH_AND_RESCUE_AIRCRAFT_REPORT:
			message := encoder.StandardSarAircraftReport(*packet.Message.StandardSearchAndRescueAircraftReport)
			message.TagBlock.Time = timestamp
			c.handler.Message(message)

		// Type 14
		case aisStream.SAFETY_BROADCAST_MESSAGE:
			message := encoder.SafetyBroadcastMessage(*packet.Message.SafetyBroadcastMessage)
			message.TagBlock.Time = timestamp
			c.handler.Message(message)

		// Type 18
		case aisStream.STANDARD_CLASS_B_POSITION_REPORT:
			message := encoder.PositionReportClassB(*packet.Message.StandardClassBPositionReport)
			message.TagBlock.Time = timestamp
			c.handler.Message(message)

		// Type 19
		case aisStream.EXTENDED_CLASS_B_POSITION_REPORT:
			message := encoder.PositionReportClassBExtended(*packet.Message.ExtendedClassBPositionReport)
			message.TagBlock.Time = timestamp
			c.handler.Message(message)

		// Type 24
		case aisStream.STATIC_DATA_REPORT:
			message := encoder.StaticDataReport(*packet.Message.StaticDataReport)
			message.TagBlock.Time = timestamp
			c.handler.Message(message)

		// Type 27
		case aisStream.LONG_RANGE_AIS_BROADCAST_MESSAGE:
			message := encoder.LongRangeAisBroadcastMessage(*packet.Message.LongRangeAisBroadcastMessage)
			message.TagBlock.Time = timestamp
			c.handler.Message(message)

		// dont bother processing these
		case aisStream.UNKNOWN_MESSAGE:
		case aisStream.ADDRESSED_SAFETY_MESSAGE:
		case aisStream.ADDRESSED_BINARY_MESSAGE:
		case aisStream.AIDS_TO_NAVIGATION_REPORT:
		case aisStream.ASSIGNED_MODE_COMMAND:
		case aisStream.BASE_STATION_REPORT:
		case aisStream.BINARY_ACKNOWLEDGE:
		case aisStream.BINARY_BROADCAST_MESSAGE:
		case aisStream.CHANNEL_MANAGEMENT:
		case aisStream.COORDINATED_UTC_INQUIRY:
		case aisStream.DATA_LINK_MANAGEMENT_MESSAGE:
		case aisStream.DATA_LINK_MANAGEMENT_MESSAGE_DATA:
		case aisStream.GROUP_ASSIGNMENT_COMMAND:
		case aisStream.GNSS_BROADCAST_BINARY_MESSAGE:
		case aisStream.INTERROGATION:
		case aisStream.MULTI_SLOT_BINARY_MESSAGE:
		case aisStream.SINGLE_SLOT_BINARY_MESSAGE:
		}

		// check for cancel signal
		select {
		case <-c.ctx.Done():
			break worker
		default:
		}
	}

	slog.Info("aisstream: stopping worker")
	return nil
}
