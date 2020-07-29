package packet

import "gukkit/net/data/types"

const (
	SpawnExperienceOrbPacketID = types.VarInt(0x01)
)

type SpawnExperienceOrbPacket struct {
	Clientbound

	EntityID types.VarInt
	X        types.Double
	Y        types.Double
	Z        types.Double
	Count    types.Short //The amount of experience this orb will reward once collected
}
