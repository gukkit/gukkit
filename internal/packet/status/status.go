package status

import (
	"gukkit/internal/packet"
	"gukkit/net/data/types"
)

const (
	ResponseID = 0
	PongID     = 1
	RequestID  = 0
	PingID     = 1
)

type ResponsePacket struct {
	packet.Clientbound
}

type PongPacket struct {
	packet.Clientbound

	Payload types.Long
}

type PingPacket struct {
	packet.Serverbound

	Payload types.Long
}

type RequestPacket struct {
	packet.Serverbound
}
