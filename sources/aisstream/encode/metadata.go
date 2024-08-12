package encode

type Metadata struct {
	Mmsi       float64 `json:"MMSI"`
	MmsiString float64 `json:"MMSI_String"`
	ShipName   string  `json:"ShipName"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	TimeUtc    string  `json:"time_utc"`
}
