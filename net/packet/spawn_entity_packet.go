package packet

import (
	"gukkit/net/data/types"
)

const (
	SpawnEntityPacketID = types.VarInt(0x00)
)

type SpawnEntityPacket struct {
	Clientbound

	EntityID   types.VarInt
	ObjectUUID types.UUID
	Type       types.VarInt
	X          types.Double
	Y          types.Double
	Z          types.Double
	Pitch      types.Angle
	Yaw        types.Angle
	Data       types.Int
	VelocityX  types.Short
	VelocityY  types.Short
	VelocityZ  types.Short
}

func (pk *SpawnEntityPacket) Encode(w Writer) (err error) {
	err = SpawnEntityPacketID.Encode(w)
	err = pk.EntityID.Encode(w)
	err = pk.ObjectUUID.Encode(w)
	err = pk.Type.Encode(w)
	err = pk.X.Encode(w)
	err = pk.Y.Encode(w)
	err = pk.Z.Encode(w)
	err = pk.Pitch.Encode(w)
	err = pk.Yaw.Encode(w)
	err = pk.Data.Encode(w)
	err = pk.VelocityX.Encode(w)
	err = pk.VelocityY.Encode(w)
	err = pk.VelocityZ.Encode(w)

	return
}
