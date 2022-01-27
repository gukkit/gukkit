package packet

import "gukkit/net/data/types"

type PlayerStatusPacket struct {
	ID       types.VarInt
	ActionID types.VarInt
}
