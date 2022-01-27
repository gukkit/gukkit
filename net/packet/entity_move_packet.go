package packet

import "gukkit/net/data/types"

const (
	EntityMovePacketID = types.VarInt(0x29)
)

type EntityMovePacket struct {
	ID       types.VarInt
	Entity   types.VarInt
	DeltaX   types.Short
	DeltaY   types.Short
	DeltaZ   types.Short
	OnGround types.Boolean
}

func (pk *EntityMovePacket) Bytes() {

}
