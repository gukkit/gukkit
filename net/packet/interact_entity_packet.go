package packet

import "gukkit/net/data/types"

type InteractEntityPacket struct {
	ID       types.VarInt
	EntityID types.VarInt
	Type     types.VarInt
	TargetX  types.Float
	TargetY  types.Float
	TargetZ  types.Float
	Hand     types.VarInt
	Sneaking types.Boolean
}
