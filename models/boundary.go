package models

type Boundary struct {
	LatMin float32
	LatMax float32
	LonMin float32
	LonMax float32
}

func (b *Boundary) Contains(position Coordinates) bool {

	// test whether the boundary and position are valid
	if b == nil || !position.IsValid() {
		return false
	}

	// test whether the position is inside the defined bounds
	latitude := position.Latitude()
	longitude := position.Longitude()
	if latitude < b.LatMin ||
		latitude > b.LatMax ||
		longitude < b.LonMin ||
		longitude > b.LonMax {
		return false
	}

	return true

}

var BoundaryAustralia *Boundary = &Boundary{LatMin: -60, LatMax: -9, LonMin: 100, LonMax: 160}
