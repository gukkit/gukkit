package packet

import "gukkit/net/data/types"

type QueryBlockNBTPacket struct {
	ID            types.VarInt
	TransactionID types.VarInt
	Location      types.Position
}
