package packet

import "gukkit/net/data/types"

type TeleportConfirmPacket struct {
	ID         types.VarInt
	TeleportID types.VarInt
}
