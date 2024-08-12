package handlers

import (
	"ais-stream/models"
	"fmt"
	"log/slog"

	ais "github.com/BertoldVdb/go-ais"
)

type Flag uint8

const (
	OK Flag = iota
	FilteredMessage
	FilteredPosition
	FilteredVessel
	Duplicate
)

const PRINT_MESSAGES = false

// this is just a function to print a stream of crap to the console
func Print(counter uint64, flag Flag, message models.Message) {

	// ignore any 'nil' packets with no header
	if message.Packet == nil && len(message.Payload) > 0 {
		slog.Error("error no packet is present in message", "message", message)
		return
	}

	if !PRINT_MESSAGES {
		return
	}

	str := ""
	switch flag {
	case FilteredMessage:
		str = "(filtered message)"
	case FilteredPosition:
		str = "(filtered message type)"
	case FilteredVessel:
		str = "(filtered vessel)"
	case Duplicate:
		str = "(duplicate)"
	case OK:
		if message.Packet != nil {
			switch message.Packet.(type) {
			case ais.PositionReport:
				str = positionReport(message.Packet.(ais.PositionReport))
			case ais.ShipStaticData:
				str = shipStaticData(message.Packet.(ais.ShipStaticData))
			case ais.StandardSearchAndRescueAircraftReport:
				str = standardSearchAndRescueAircraftReport(message.Packet.(ais.StandardSearchAndRescueAircraftReport))
			case ais.SafetyBroadcastMessage:
				str = safetyBroadcastMessage(message.Packet.(ais.SafetyBroadcastMessage))
			case ais.StandardClassBPositionReport:
				str = standardClassBPositionReport(message.Packet.(ais.StandardClassBPositionReport))
			case ais.ExtendedClassBPositionReport:
				str = extendedClassBPositionReport(message.Packet.(ais.ExtendedClassBPositionReport))
			case ais.StaticDataReport:
				str = staticDataReport(message.Packet.(ais.StaticDataReport))
			case ais.LongRangeAisBroadcastMessage:
				str = longRangeAisBroadcastMessage(message.Packet.(ais.LongRangeAisBroadcastMessage))
			case ais.BaseStationReport:
				str = baseStationReport(message.Packet.(ais.BaseStationReport))
			case ais.AddressedBinaryMessage:
				//str = addressedBinaryMessage(message.Packet.(ais.AddressedBinaryMessage))
			case ais.BinaryAcknowledge:
				//str = binaryAcknowledge(message.Packet.(ais.BinaryAcknowledge))
			case ais.BinaryBroadcastMessage:
				//str = binaryBroadcastMessage(message.Packet.(ais.BinaryBroadcastMessage))
			case ais.CoordinatedUTCInquiry:
				//str = coordinatedUTCInquiry(message.Packet.(ais.CoordinatedUTCInquiry))
			case ais.AddessedSafetyMessage:
				str = addessedSafetyMessage(message.Packet.(ais.AddessedSafetyMessage))
			case ais.Interrogation:
				//str = interrogation(message.Packet.(ais.Interrogation))
			case ais.AssignedModeCommand:
				//str = assignedModeCommand(message.Packet.(ais.AssignedModeCommand))
			case ais.GnssBroadcastBinaryMessage:
				//str = gnssBroadcastBinaryMessage(message.Packet.(ais.GnssBroadcastBinaryMessage))
			case ais.DataLinkManagementMessage:
				//str = dataLinkManagementMessage(message.Packet.(ais.DataLinkManagementMessage))
			case ais.AidsToNavigationReport:
				str = aidsToNavigationReport(message.Packet.(ais.AidsToNavigationReport))
			case ais.ChannelManagement:
				//str = channelManagement(message.Packet.(ais.ChannelManagement))
			case ais.GroupAssignmentCommand:
				//str = groupAssignmentCommand(message.Packet.(ais.GroupAssignmentCommand))
			case ais.SingleSlotBinaryMessage:
				//str = singleSlotBinaryMessage(message.Packet.(ais.SingleSlotBinaryMessage))
			case ais.MultiSlotBinaryMessage:
				//str = multiSlotBinaryMessage(message.Packet.(ais.MultiSlotBinaryMessage))
			default:
			}
		}
	}

	header := message.Packet.GetHeader()
	fmt.Printf("%d %d src: %-16s mmsi: %-9d id: %2d msg: %s %s\r\n",
		counter,
		message.TagBlock.Time,
		message.TagBlock.Source,
		header.UserID,
		header.MessageID,
		getMessageName(header.MessageID),
		str,
	)

}

func positionReport(in ais.PositionReport) string {
	return fmt.Sprintf("lat: %.4f lon: %.4f", float64(in.Latitude), float64(in.Longitude))
}
func shipStaticData(in ais.ShipStaticData) string {
	return in.Name
}
func standardSearchAndRescueAircraftReport(in ais.StandardSearchAndRescueAircraftReport) string {
	return fmt.Sprintf("lat: %.4f lon: %.4f", float64(in.Latitude), float64(in.Longitude))
}
func safetyBroadcastMessage(in ais.SafetyBroadcastMessage) string {
	return in.Text
}
func standardClassBPositionReport(in ais.StandardClassBPositionReport) string {
	return fmt.Sprintf("lat: %.4f lon: %.4f", float64(in.Latitude), float64(in.Longitude))
}
func extendedClassBPositionReport(in ais.ExtendedClassBPositionReport) string {
	return fmt.Sprintf("lat: %.4f lon: %.4f", float64(in.Latitude), float64(in.Longitude))
}
func staticDataReport(in ais.StaticDataReport) string {
	return in.ReportA.Name
}
func longRangeAisBroadcastMessage(in ais.LongRangeAisBroadcastMessage) string {
	return fmt.Sprintf("lat: %.4f lon: %.4f", float64(in.Latitude), float64(in.Longitude))
}
func baseStationReport(in ais.BaseStationReport) string {
	return fmt.Sprintf("lat: %.4f lon: %.4f", float64(in.Latitude), float64(in.Longitude))
}

//	func addressedBinaryMessage(in ais.AddressedBinaryMessage) string {
//		return ""
//	}
//
//	func binaryAcknowledge(in ais.BinaryAcknowledge) string {
//		return ""
//	}
//
//	func binaryBroadcastMessage(in ais.BinaryBroadcastMessage) string {
//		return ""
//	}
//
//	func coordinatedUTCInquiry(in ais.CoordinatedUTCInquiry) string {
//		return ""
//	}
func addessedSafetyMessage(in ais.AddessedSafetyMessage) string {
	return in.Text
}

//	func interrogation(in ais.Interrogation) string {
//		return ""
//	}
//
//	func assignedModeCommand(in ais.AssignedModeCommand) string {
//		return ""
//	}
//
//	func gnssBroadcastBinaryMessage(in ais.GnssBroadcastBinaryMessage) string {
//		return ""
//	}
//
//	func dataLinkManagementMessage(in ais.DataLinkManagementMessage) string {
//		return ""
//	}
func aidsToNavigationReport(in ais.AidsToNavigationReport) string {
	return fmt.Sprintf("lat: %.4f lon: %.4f", float64(in.Latitude), float64(in.Longitude))
}

// func channelManagement(in ais.ChannelManagement) string {
// 	return ""
// }
// func groupAssignmentCommand(in ais.GroupAssignmentCommand) string {
// 	return ""
// }
// func singleSlotBinaryMessage(in ais.SingleSlotBinaryMessage) string {
// 	return ""
// }
// func multiSlotBinaryMessage(in ais.MultiSlotBinaryMessage) string {
// 	return ""
// }

func getMessageName(id uint8) string {
	var messageName []string = []string{
		"Reserved",
		"Position report",                       //1	Scheduled position report (Class A shipborne mobile equipment)
		"Position report",                       //2	Assigned scheduled position report; (Class A shipborne mobile equipment)
		"Position report",                       //3	Special position report, response to interrogation; (Class A shipborne mobile equipment)
		"Base station report",                   //4	Position, UTC, date and current slot number of base station
		"Static & voyage data",                  //5	Scheduled static and voyage related vessel data report; (Class A shipborne mobile equipment)
		"Binary addressed message",              //6	Binary data for addressed communicationView Binary Messages
		"Binary acknowledgement",                //7	Acknowledgement of received addressed binary data
		"Binary broadcast message",              //8	Binary data for broadcast communicationView Binary Messages
		"Standard SAR aircraft position report", //9	Position report for airborne stations involved in SAR operations, only
		"UTC/date inquiry",                      //10	Request UTC and date
		"UTC/date response",                     //11	Current UTC and date if available
		"Addressed safety related message",      //12	Safety related data for addressed communication
		"Safety related acknowledgement",        //13	Acknowledgement of received addressed safety related message
		"Safety related broadcast message",      //14	Safety related data for broadcast communication
		"Interrogation",                         //15	Request for a specific message type (can result in multiple responses from one or several stations
		"Assignment mode command",               //16	Assignment of a specific report behaviour by competent authority using a Base station
		"DGNSS broadcast binary message",        //17	DGNSS corrections provided by a base station
		"Standard Class B position report",      //18	Standard position report for Class B shipborne mobile equipment to be used instead of Messages 1, 2, 3
		"Extended Class B position report",      //19	Extended position report for class B shipborne mobile equipment; contains additional static information
		"Data link management message",          //20	Reserve slots for Base station(s)
		"Aids-to-Navigation report",             //21	Position and status report for aids-to-navigation
		"Channel management",                    //22	Management of channels and transceiver modes by a Base station
		"Group assignment command",              //23	Assignment of a specific report behaviour by competent authority using a Base station to a specific group of mobiles
		"Static data report",                    //24	Additional data assigned to an MMSI Part A: Name, Part B: Static Data
		"Single slot binary message",            //25	Short unscheduled binary data transmission (broadcast or addressed)
		"Multiple slot binary message",          //26	Scheduled binary data transmission (broadcast or addressed)
		"Long range position report",            //27	Scheduled position report; Class A shipborne mobile equipment outside base station coverage
	}

	if int(id) >= len(messageName) {
		return "error"
	}
	return messageName[id]
}
