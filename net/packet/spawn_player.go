package packet

import "gukkit/net/data/types"

const (
	SpawnPlayerPacketID = types.VarInt(0x05)
)

type SpawnPlayerPacket struct {
	Clientbound

	EntityID   types.VarInt
	PlayerUUID types.VarInt
	X          types.Double
	Y          types.Double
	Z          types.Double
	Yaw        types.Double
	Pitch      types.Double
}
