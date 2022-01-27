package packet

import "gukkit/net/data/types"

// Window type	ID	Meaning

// Enchantment Table
// 	0	Topmost enchantment.
// 	1	Middle enchantment.
// 	2	Bottom enchantment.
// Lectern
// 	1	Previous page (which does give a redstone output).
// 	2	Next page.
// 	3	Take Book.
// 	100+page	Opened page number - 100 + number.
// Stonecutter
// 	Recipe button number - 4*row + col. Depends on the item.
// Loom
// 	Recipe button number - 4*row + col. Depends on the item.

type ClickWindowButtonPacket struct {
	ID       types.VarInt
	WindowID types.Byte
	ButtonID types.Byte
}
