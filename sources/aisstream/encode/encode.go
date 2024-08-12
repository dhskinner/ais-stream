package encode

import (
	"ais-stream/models"

	ais "github.com/BertoldVdb/go-ais"
	aisnmea "github.com/BertoldVdb/go-ais/aisnmea"
	nmea "github.com/adrianmo/go-nmea"
)

type Encoder struct {
	AisCodec  *ais.Codec
	NmeaCodec *aisnmea.NMEACodec
}

func NewEncoder() *Encoder {

	e := &Encoder{}
	e.AisCodec = ais.CodecNew(false, false)
	e.AisCodec.DropSpace = true
	e.NmeaCodec = aisnmea.NMEACodecNew(ais.CodecNew(false, false))
	return e

}

func (e *Encoder) AsMessage(packet ais.Packet) models.Message {

	// re-encode as a VDM packet
	bytes := e.AisCodec.EncodePacket(packet)
	tb := nmea.TagBlock{
		Text:   AisstreamTag,
		Source: AisstreamSource,
	}

	vdm := &aisnmea.VdmPacket{
		Channel:     ChannelA,
		TalkerID:    AisstreamTalkerId,
		MessageType: AisstreamMessageId,
		Payload:     bytes,
		Packet:      packet,
		TagBlock:    tb,
	}

	return models.Message(vdm)

}
