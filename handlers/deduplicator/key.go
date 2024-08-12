package deduplicator

import (
	"ais-stream/models"
	"fmt"

	ais "github.com/BertoldVdb/go-ais"
)

// this is just a function to get a couple of parameters that should be the same
// if a duplicate message is received, and different for non-duplicates
func getKey(message models.Message) string {

	if message.Packet == nil {
		return ""
	}

	switch message.Packet.(type) {
	case ais.PositionReport:
		return positionReport(message.Packet.(ais.PositionReport))
	case ais.StandardSearchAndRescueAircraftReport:
		return standardSearchAndRescueAircraftReport(message.Packet.(ais.StandardSearchAndRescueAircraftReport))
	case ais.StandardClassBPositionReport:
		return standardClassBPositionReport(message.Packet.(ais.StandardClassBPositionReport))
	case ais.ExtendedClassBPositionReport:
		return extendedClassBPositionReport(message.Packet.(ais.ExtendedClassBPositionReport))
	case ais.LongRangeAisBroadcastMessage:
		return longRangeAisBroadcastMessage(message.Packet.(ais.LongRangeAisBroadcastMessage))

	// these are all the same key
	// case ais.SafetyBroadcastMessage:
	// case ais.ShipStaticData:
	// case ais.StaticDataReport:
	// case ais.BaseStationReport:
	// case ais.AidsToNavigationReport:
	// case ais.AddressedBinaryMessage:
	// case ais.BinaryAcknowledge:
	// case ais.BinaryBroadcastMessage:
	// case ais.CoordinatedUTCInquiry:
	// case ais.AddessedSafetyMessage:
	// case ais.Interrogation:
	// case ais.AssignedModeCommand:
	// case ais.GnssBroadcastBinaryMessage:
	// case ais.DataLinkManagementMessage:
	// case ais.ChannelManagement:
	// case ais.GroupAssignmentCommand:
	// case ais.SingleSlotBinaryMessage:
	// case ais.MultiSlotBinaryMessage:
	default:
		in := message.Packet.GetHeader()
		return fmt.Sprintf("%9d%2d", in.UserID, in.MessageID)
	}
}

func positionReport(in ais.PositionReport) string {
	return fmt.Sprintf("%9d%2d%.6f%.6f", in.UserID, in.MessageID, float64(in.Latitude), float64(in.Longitude))
}

func standardSearchAndRescueAircraftReport(in ais.StandardSearchAndRescueAircraftReport) string {
	return fmt.Sprintf("%9d%2d%.6f%.6f", in.UserID, in.MessageID, float64(in.Latitude), float64(in.Longitude))
}

func standardClassBPositionReport(in ais.StandardClassBPositionReport) string {
	return fmt.Sprintf("%9d%2d%.6f%.6f", in.UserID, in.MessageID, float64(in.Latitude), float64(in.Longitude))
}

func extendedClassBPositionReport(in ais.ExtendedClassBPositionReport) string {
	return fmt.Sprintf("%9d%2d%.6f%.6f", in.UserID, in.MessageID, float64(in.Latitude), float64(in.Longitude))
}

func longRangeAisBroadcastMessage(in ais.LongRangeAisBroadcastMessage) string {
	return fmt.Sprintf("%9d%2d%.6f%.6f", in.UserID, in.MessageID, float64(in.Latitude), float64(in.Longitude))
}
