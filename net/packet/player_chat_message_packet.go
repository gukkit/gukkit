package packet

import "gukkit/net/data/types"

type PlayerChatMessagePacket struct {
	ID      types.VarInt
	Message types.String
}
