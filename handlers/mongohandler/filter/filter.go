package filter

import (
	"ais-stream/interfaces"
	"ais-stream/models"
)

// Filter screens out processing by message id, for vessel updates, or position updates
type Filter struct {
	Name              string
	MessageWhitelist  *Whitelist
	VesselWhitelist   *Whitelist
	PositionWhitelist *Whitelist
	database          *Database
}

// create a new filter
func New(
	name string,
	messageWhitelist *Whitelist,
	vesselWhitelist *Whitelist,
	positionWhitelist *Whitelist,
	handler interfaces.Handler) *Filter {

	f := &Filter{
		Name:              name,
		MessageWhitelist:  messageWhitelist,
		VesselWhitelist:   vesselWhitelist,
		PositionWhitelist: positionWhitelist,
		database:          NewDatabase(handler),
	}
	return f

}

// check a whitelist against key parameters (messagetype is ignored)
func (f *Filter) IsWhitelisted(
	mmsi models.MMSI,
	messageId *models.MessageId,
	shiptype *models.ShipTypeId,
	position *models.Coordinates,
	whitelist *Whitelist) bool {

	if whitelist == nil || mmsi == 0 {
		return false
	}

	// update the in-memory database
	record := f.database.GetAndUpdate(mmsi, shiptype, position)

	// check the message id
	if messageId != nil && !whitelist.IsMessageIdIncluded(*messageId) {
		return false
	}

	// check the last known vessel position
	if !whitelist.IsPositionIncluded(record.Position) {
		return false
	}

	// check the shiptype and/or mmsi as a pair
	if !whitelist.IsShipTypeIncluded(record.ShipType) &&
		!whitelist.IsMmsiIncluded(record.Mmsi) {
		return false
	}

	return true
}

// test whether the message is acceptable to our whitelist
func (f *Filter) IsMessageIncluded(message models.Message) bool {

	if message.Packet == nil || f.MessageWhitelist == nil {
		return false
	}

	// check the message parameters that are known at this stage
	// if a parameter is not known, the last known value is used
	header := message.Packet.GetHeader()
	return f.IsWhitelisted(header.UserID, (*models.MessageId)(&header.MessageID), nil, nil, f.MessageWhitelist)

}

// test whether the position is acceptable to our whitelist
func (f *Filter) IsPositionIncluded(pos *models.VesselPosition) bool {

	if pos == nil || f.PositionWhitelist == nil {
		return false
	}

	// check the vessel parameters that are known at this stage
	// if a parameter is not known, the last known value is used
	return f.IsWhitelisted(pos.Mmsi, nil, nil, &pos.Position, f.PositionWhitelist)

}
