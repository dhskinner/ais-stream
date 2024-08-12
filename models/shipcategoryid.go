package models

type ShipCategoryId uint8

const (
	ShipCategoryUnknown        ShipCategoryId = 0
	ShipCategoryReserved       ShipCategoryId = 1
	ShipCategoryWingInGround   ShipCategoryId = 2
	ShipCategorySpecial1       ShipCategoryId = 3
	ShipCategoryHighSpeedCraft ShipCategoryId = 4
	ShipCategorySpecial2       ShipCategoryId = 5
	ShipCategoryPassenger      ShipCategoryId = 6
	ShipCategoryCargo          ShipCategoryId = 7
	ShipCategoryTanker         ShipCategoryId = 8
	ShipCategoryOther          ShipCategoryId = 9
)

func (s ShipCategoryId) AsString() string {
	index := uint8(s)
	return shipCategoryLabel[index]
}

var shipCategoryLabel []string = []string{
	"Unknown",
	"Reserved",
	"Wing In Ground",
	"Special Category",
	"High-Speed Craft",
	"Special Category",
	"Passenger",
	"Cargo",
	"Tanker",
	"Other",
}
