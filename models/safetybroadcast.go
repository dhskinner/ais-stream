package models

type SafetyBroadcastMessage struct {
	Mmsi  MMSI   `aisWidth:"38"`
	Valid bool   `aisEncodeMaxLen:"1008"`
	Text  string `aisWidth:"-1"`
}
