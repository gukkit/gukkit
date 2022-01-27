package nbt

import "bytes"

type NBT struct {
	bytes.Buffer
	pointID int8
	point   Tag
	root    Tag
}

func NewEmptyNBT() *NBT {
	return &NBT{}
}

func NewCompoundNBT(name string) *NBT {
	nbt := NewEmptyNBT()
	nbt.root = NewCompoundTag(name, nil)
	nbt.point = nbt.root
	return nbt
}

func (nbt *NBT) AddByteTag(name string, b byte) *NBT {
	if !nbt.IsCompound() && !nbt.IsList() {
		panic("current nbt tag is not a list or compound tag, we can't write other tag")
	}

	return nbt
}

func (nbt *NBT) IsArray() (yes bool) {
	return nbt.pointID == NBT_TAG_Byte_Array ||
		nbt.pointID == NBT_TAG_Int_Array ||
		nbt.pointID == NBT_TAG_Long_Array
}

func (nbt *NBT) IsCompound() bool {
	return nbt.pointID == NBT_TAG_Compound
}

func (nbt *NBT) IsList() bool {
	return nbt.pointID == NBT_TAG_List
}

func (nbt *NBT) Encode() {
	next := nbt.root
	for next != nil {

	}
}
