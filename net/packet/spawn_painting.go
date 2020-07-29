package packet

import "gukkit/net/data/types"

const (
	SpawnPaintingPacketID = types.VarInt(0x04)

	//Direction
	East  = 3 //东
	South = 0 //南
	West  = 1 //西
	North = 2 //北
)

type SpawnPaintingPacket struct {
	Clientbound

	EntityID   types.VarInt
	EntityUUID types.UUID
	Motive     types.VarInt
	Location   types.Position
	Direction  types.Byte //only value[0,1,2,3]
}
