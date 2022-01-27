package types

type Slot struct {
	Present   bool
	ItemID    Byte
	ItemCount Byte
	// NBT       nbt.EndTag
}
