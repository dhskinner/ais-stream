package models

// [lat, lon] coordinate pair
type Coordinates []Degrees

func (p Coordinates) Latitude() Degrees {
	if len(p) < 2 {
		return 0
	}
	return p[0]
}

func (p Coordinates) Longitude() Degrees {
	if len(p) < 2 {
		return 0
	}
	return p[1]
}

func NewCoordinates(latitude float32, longitude float32) Coordinates {
	return Coordinates{latitude, longitude}
}

func (p Coordinates) IsValid() bool {
	return p.Latitude() <= 90 &&
		p.Latitude() >= -90 &&
		p.Longitude() <= 180 &&
		p.Longitude() >= -180 &&
		!(p.Longitude() == 0 && p.Latitude() == 0)
}

func (p *Coordinates) AsState() State {

	if p.Longitude() < 108 ||
		p.Latitude() > 9.5 ||
		p.Longitude() > 156.5 ||
		p.Latitude() < -45 {
		return StateInt
	}
	if p.Longitude() < 129 {
		return StateWA
	}
	if p.Latitude() > -26 && p.Longitude() < 138 {
		return StateNT
	}
	if p.Latitude() < -26 && p.Longitude() < 141 {
		return StateSA
	}
	if p.Latitude() > -28.16432571631462 { // Pt Danger
		return StateQLD
	}
	if p.Latitude() > -37.50526603006263 {
		return StateNSW
	}
	if p.Latitude() > -39.55 {
		return StateVIC
	}

	return ""
}
