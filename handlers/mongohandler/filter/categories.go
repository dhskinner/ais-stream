package filter

type FilterCategory uint8

const (
	FilterByMessageId FilterCategory = 1
	FilterByBoundary  FilterCategory = 2
	FilterByMmsi      FilterCategory = 4
	FilterByShipType  FilterCategory = 8
)

func (f FilterCategory) Includes(flag FilterCategory) bool {

	return f&flag == flag

}
