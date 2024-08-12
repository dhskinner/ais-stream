package models

import ais "github.com/BertoldVdb/go-ais"

type Dimension ais.FieldDimension

// Dimension represents the encoding of the dimension
func (d *Dimension) Length() Metres {
	if d == nil {
		return 0
	}
	return d.A + d.B
}

func (d *Dimension) Beam() Metres {
	if d == nil {
		return 0
	}
	return uint16(d.C) + uint16(d.D)
}
