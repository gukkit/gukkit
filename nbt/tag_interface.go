package nbt

import "io"

var (
	NBT_TAG_End        int8 = 0
	NBT_TAG_Byte       int8 = 1
	NBT_TAG_Short      int8 = 2
	NBT_TAG_Int        int8 = 3
	NBT_TAG_Long       int8 = 4
	NBT_TAG_Float      int8 = 5
	NBT_TAG_Double     int8 = 6
	NBT_TAG_Byte_Array int8 = 7
	NBT_TAG_String     int8 = 8
	NBT_TAG_List       int8 = 9
	NBT_TAG_Compound   int8 = 10
	NBT_TAG_Int_Array  int8 = 11
	NBT_TAG_Long_Array int8 = 12
)

type Tag interface {
	ID() byte
	Name() string
	NameLength() int16
}

type ListableTag interface {
	TagLength() int
	FindTags() []Tag
	FindTag(index int) Tag
	AddTag() bool
	DeleteTag(index int) bool
	UpdateTag(index int, tag Tag) bool
}

type defaultTag struct {
	id         int8
	nameLength int16
	name       string
}

func newDefaultTag(id int8, name string) defaultTag {
	return defaultTag{
		id:         id,
		name:       name,
		nameLength: int16(len(name)),
	}
}

func (tag defaultTag) ID() byte {
	return byte(tag.id)
}

func (tag defaultTag) NameLength() int16 {
	return tag.nameLength
}

func (tag defaultTag) Name() string {
	return tag.name
}

func (tag defaultTag) onHeadBytes(w io.Writer) (err error) {
	_, err = w.Write([]byte{
		byte(tag.id),
		byte(tag.nameLength << 8),
		byte(tag.nameLength),
	})
	return
}
