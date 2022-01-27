package packet

import "gukkit/net/data/types"

type ClickWindowPacket struct {
	ID          types.VarInt
	WindowID    types.UnsignedByte
	StateID     types.VarInt
	Slot        types.Short
	Button      types.Byte
	Mode        types.VarInt
	ArrayLength types.VarInt
}
