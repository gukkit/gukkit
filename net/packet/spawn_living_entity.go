package packet

import "gukkit/net/data/types"

const (
	SpawnLivingEntityPacketID = types.VarInt(0x03)
)

type SpawnLivingEntityPacket struct {
	Clientbound

	EntityID   types.VarInt
	EntityUUID types.UUID
	Type       types.VarInt
	X          types.Double
	Y          types.Double
	Z          types.Double
	Yaw        types.Double
	Pitch      types.Double
	HeadPitch  types.Double
	VelocityX  types.Short
	VelocityY  types.Short
	VelocityZ  types.Short
}
