package packet

import "gukkit/net/data/types"

const (
	SpawnWeatherEntityPacketID = types.VarInt(0x02)
)

type SpawnWeatherEntityPacket struct {
	Clientbound

	EntityID types.VarInt
	Type     types.Byte //the global entity type, currently always 1 for thunderbolt
	X        types.Double
	Y        types.Double
	Z        types.Double
}
