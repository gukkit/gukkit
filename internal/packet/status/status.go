package status

import (
	"gukkit/net/data/types"
	"gukkit/net/packet"
)

const (
	ResponseID = 0
	PingPongID = types.VarInt(1)
	RequestID  = 0
)

type ResponsePacket struct {
	packet.Clientbound
}

type RequestPacket struct {
	packet.Serverbound

	ID types.VarInt
}

type PingPongPacket struct {
	packet.Serverbound
	packet.Clientbound

	ID      types.VarInt
	Payload types.Long
}

func (pk *PingPongPacket) Encode(w packet.Writer) (err error) {
	if err = PingPongID.Encode(w); err != nil {
		return
	}

	err = pk.Payload.Encode(w)
	return
}

func (pk *PingPongPacket) Decode(r packet.Reader) (err error) {
	if pk.ID != PingPongID {
		err = packet.DecodeIDNotEqualErr
		return
	}

	_, err = pk.Payload.Decode(r)
	return
}
