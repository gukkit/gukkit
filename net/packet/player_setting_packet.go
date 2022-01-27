package packet

import "gukkit/net/data/types"

// Displayed Skin Parts flags:

// Bit 0 (0x01): Cape enabled
// Bit 1 (0x02): Jacket enabled
// Bit 2 (0x04): Left Sleeve enabled
// Bit 3 (0x08): Right Sleeve enabled
// Bit 4 (0x10): Left Pants Leg enabled
// Bit 5 (0x20): Right Pants Leg enabled
// Bit 6 (0x40): Hat enabled

type PlayerSettingPacket struct {
	ID                  types.VarInt
	ViewDistance        types.Byte
	ChatMode            types.VarInt //0: enabled, 1: commands only, 2: hidden
	ChatColor           types.Boolean
	DisplayedSkinParts  types.UnsignedByte
	MainHand            types.VarInt //0: Left, 1: Right.
	EnableTextFiltering types.Boolean
	AllowServerListings types.Boolean
}
