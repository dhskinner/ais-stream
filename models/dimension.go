package models

import ais "github.com/BertoldVdb/go-ais"

type Dimension ais.FieldDimension

// Dimension represents the encoding of the dimension
// type Dimension struct {
// 	A uint16
// 	B uint16
// 	C uint8
// 	D uint8
// }

func (d *Dimension) Length() Metres {
	return d.A + d.B
}

func (d *Dimension) Beam() Metres {
	return uint16(d.C) + uint16(d.D)
}
