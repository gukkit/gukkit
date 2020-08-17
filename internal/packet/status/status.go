package status

import (
	"gukkit/net/data/types"
	"gukkit/net/packet"
)

const (
	ResponseID = types.VarInt(0)
	RequestID  = types.VarInt(0)
	PingPongID = types.VarInt(1)
)

type ResponsePacket struct {
	packet.Clientbound

	Message types.String
}

func (pk ResponsePacket) Encode(w packet.Writer) (err error) {
	if err = ResponseID.Encode(w); err != nil {
		return
	}

	if err = pk.Message.Encode(w); err != nil {
		return
	}

	return
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

func (pk PingPongPacket) Encode(w packet.Writer) (err error) {
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
