package filter

import (
	"ais-stream/models"
	"sync"
)

// set a whitelist of parameters to process:
// - message id's = messages 1 to 27 (message 0 is not included)
// - bounds = the lat/lon of area to include for vessel updates (based on last known vessel position)
// - mmsis = whitelist of vessels of interest to explicitly include (regardless of ship type), and
// - shipTypes = whitelist of vessel types to include
type Whitelist struct {
	filterby    FilterCategory
	boundary    *models.Boundary
	mmsis       sync.Map
	messageIds  [28]bool
	shipTypeIds [100]bool
}

// create a new filter
func NewWhitelist(
	boundary *models.Boundary,
	mmsis []models.MMSI,
	messageIds []models.MessageId,
	shipTypeIds []models.ShipTypeId) *Whitelist {

	f := &Whitelist{
		boundary: boundary,
	}

	// initialise the filter boundary
	if boundary != nil {
		f.filterby = f.filterby | FilterByBoundary
	}

	// initialise a map of mmsi's
	if len(mmsis) > 0 {
		f.filterby = f.filterby | FilterByMmsi
	}
	for _, mmsi := range mmsis {
		f.mmsis.Store(mmsi, true)
	}

	// initialise an array of 27 messages (0 is unused)
	if len(messageIds) > 0 {
		f.filterby = f.filterby | FilterByMessageId
	}
	for _, id := range messageIds {
		f.messageIds[id] = true
	}

	// initialise an array of 99 vessel types (0 is unused)
	if len(shipTypeIds) > 0 {
		f.filterby = f.filterby | FilterByShipType
	}
	for _, id := range shipTypeIds {
		f.shipTypeIds[id] = true
	}
	return f

}

// test whether the position is within our whitelisted boundary
func (f *Whitelist) IsPositionIncluded(position models.Coordinates) bool {

	// if the category is not set, then let it through
	if !f.filterby.Includes(FilterByBoundary) || f.boundary == nil {
		return true
	}

	// test whether the position is inside the defined bounds
	return f.boundary.Contains(position)

}

// test whether the mmsi is within our whitelisted set of vessels to include
func (f *Whitelist) IsMmsiIncluded(mmsi models.MMSI) bool {

	// if the category is not set, then let it through
	if !f.filterby.Includes(FilterByMmsi) {
		return true
	}

	// is the vessel mmsi present?
	_, ok := f.mmsis.Load(mmsi)
	return ok

}

// test whether the messageid is present in our whitelist
func (f *Whitelist) IsMessageIdIncluded(id models.MessageId) bool {

	// if the category is not set, then let it through
	if !f.filterby.Includes(FilterByMessageId) {
		return true
	}

	if uint8(id) == 0 || uint8(id) > 27 || !f.messageIds[id] {
		return false
	}
	return true

}

// test whether the messageid is present in our whitelist
func (f *Whitelist) IsShipTypeIncluded(id models.ShipTypeId) bool {

	// if the category is not set, then let it through
	if !f.filterby.Includes(FilterByShipType) {
		return true
	}

	if uint8(id) == 0 || uint8(id) > 99 || !f.shipTypeIds[id] {
		return false
	}
	return true

}
