package transport

import (
	"gukkit/net/packet"
)

type RecvPacketNode struct {
	pk packet.Packet

	Next *RecvPacketNode
}
