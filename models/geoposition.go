package models

type GeoPosition struct {
	Type        string    `bson:"type"`
	Coordinates []Degrees `bson:"coordinates,truncate"`
}

func NewGeoPosition(pos Coordinates) GeoPosition {
	return GeoPosition{
		Type:        "Point",
		Coordinates: []Degrees{pos.Longitude(), pos.Latitude()},
	}
}

func (p *GeoPosition) Latitude() Degrees {
	return p.Coordinates[1]
}

func (p *GeoPosition) Longitude() Degrees {
	return p.Coordinates[0]
}

func (p *GeoPosition) IsValid() bool {
	return p.Latitude() <= 90 &&
		p.Latitude() >= -90 &&
		p.Longitude() <= 180 &&
		p.Longitude() >= -180 &&
		!(p.Longitude() == 0 && p.Latitude() == 0)
}
