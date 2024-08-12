package models

type MessageId uint8

const (
	Message01_PositionReport                        MessageId = 1
	Message02_PositionReport                        MessageId = 2
	Message03_PositionReport                        MessageId = 3
	Message04_BaseStationReport                     MessageId = 4
	Message05_ShipStaticData                        MessageId = 5
	Message06_AddressedBinaryMessage                MessageId = 6
	Message07_BinaryAcknowledge                     MessageId = 7
	Message08_BinaryBroadcastMessage                MessageId = 8
	Message09_StandardSearchAndRescueAircraftReport MessageId = 9
	Message10_CoordinatedUTCInquiry                 MessageId = 10
	Message11_BaseStationReport                     MessageId = 11
	Message12_AddessedSafetyMessage                 MessageId = 12
	Message13_BinaryAcknowledge                     MessageId = 13
	Message14_SafetyBroadcastMessage                MessageId = 14
	Message15_Interrogation                         MessageId = 15
	Message16_AssignedModeCommand                   MessageId = 16
	Message17_GnssBroadcastBinaryMessage            MessageId = 17
	Message18_StandardClassBPositionReport          MessageId = 18
	Message19_ExtendedClassBPositionReport          MessageId = 19
	Message20_DataLinkManagementMessage             MessageId = 20
	Message21_AidsToNavigationReport                MessageId = 21
	Message22_ChannelManagement                     MessageId = 22
	Message23_GroupAssignmentCommand                MessageId = 23
	Message24_StaticDataReport                      MessageId = 24
	Message25_SingleSlotBinaryMessage               MessageId = 25
	Message26_MultiSlotBinaryMessage                MessageId = 26
	Message27_LongRangeAisBroadcastMessage          MessageId = 27
)

var StandardMessages = []MessageId{
	Message01_PositionReport,
	Message02_PositionReport,
	Message03_PositionReport,
	Message04_BaseStationReport,
	Message05_ShipStaticData,
	Message09_StandardSearchAndRescueAircraftReport,
	Message11_BaseStationReport,
	Message14_SafetyBroadcastMessage,
	Message18_StandardClassBPositionReport,
	Message19_ExtendedClassBPositionReport,
	Message21_AidsToNavigationReport,
	Message24_StaticDataReport,
	Message27_LongRangeAisBroadcastMessage,
}

func (n MessageId) AsString() string {
	if n > 27 {
		return messageIdLabel[0]
	}
	return messageIdLabel[n]
}

var messageIdLabel []string = []string{
	"Undefined",
	"Scheduled position report",
	"Assigned scheduled position report",
	"Special position report",
	"Base station report",
	"Static and voyage related data",
	"Binary addressed message",
	"Binary acknowledgement",
	"Binary broadcast message",
	"Standard SAR aircraft position report",
	"UTC/date inquiry",
	"UTC/date response",
	"Addressed safety related message",
	"Safety related acknowledgement",
	"Safety related broadcast message",
	"Interrogation",
	"Assignment mode command",
	"DGNSS broadcast binary message",
	"Standard Class B equipment position report",
	"Extended Class B equipment position report",
	"Data link management message",
	"Aids-to-Navigation report",
	"Channel management",
	"Group assignment command",
	"Static data report",
	"Single slot binary message",
	"Multiple slot binary message with Communications State",
	"Position report for long range applications",
}
