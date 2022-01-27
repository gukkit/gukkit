package nbt

var (
	NoneTagName = ""
)

type EndTag struct {
	defaultTag
}

type ByteTag struct {
	defaultTag
	Payload byte
}

type ShortTag struct {
	defaultTag
	Payload int16
}

type IntTag struct {
	defaultTag
	Payload int32
}

type LongTag struct {
	defaultTag
	Payload int64
}

type FloatTag struct {
	defaultTag
	Payload float32
}

type DoubleTag struct {
	defaultTag
	Payload float64
}

type ByteArrayTag struct {
	defaultTag
	Payload []byte
}

type StringTag struct {
	defaultTag
	Payload string
}

type ListTag struct {
	defaultTag
	Payload []Tag
}

type CompoundTag struct {
	defaultTag
	Payload []Tag
}

type IntArrayTag struct {
	defaultTag
	Payload []int32
}

type LongArrayTag struct {
	defaultTag
	Payload []int64
}

func NewByteTag(name string, value byte) *ByteTag {
	return &ByteTag{
		defaultTag: newDefaultTag(NBT_TAG_Byte, name),
		Payload:    value,
	}
}

func NewShortTag(name string, value int16) *ShortTag {
	return &ShortTag{
		defaultTag: newDefaultTag(NBT_TAG_Short, name),
		Payload:    value,
	}
}

func NewIntTag(name string, value int32) *IntTag {
	return &IntTag{
		defaultTag: newDefaultTag(NBT_TAG_Int, name),
		Payload:    value,
	}
}

func NewLongTag(name string, value int64) *LongTag {
	return &LongTag{
		defaultTag: newDefaultTag(NBT_TAG_Long, name),
		Payload:    value,
	}
}

func NewFloatTag(name string, value float32) *FloatTag {
	return &FloatTag{
		defaultTag: newDefaultTag(NBT_TAG_Float, name),
		Payload:    value,
	}
}

func NewDoubleTag(name string, value float64) *DoubleTag {
	return &DoubleTag{
		defaultTag: newDefaultTag(NBT_TAG_Double, name),
		Payload:    value,
	}
}

func NewByteArrayTag(name string, value []byte) *ByteArrayTag {
	return &ByteArrayTag{
		defaultTag: newDefaultTag(NBT_TAG_Byte_Array, name),
		Payload:    value,
	}
}

func NewStringTag(name, value string) *StringTag {
	return &StringTag{
		defaultTag: newDefaultTag(NBT_TAG_String, name),
		Payload:    value,
	}
}

func NewListTag(name string, value []Tag) *ListTag {
	return &ListTag{
		defaultTag: newDefaultTag(NBT_TAG_List, name),
		Payload:    value,
	}
}

func NewCompoundTag(name string, value []Tag) *CompoundTag {
	return &CompoundTag{
		defaultTag: newDefaultTag(NBT_TAG_Compound, name),
		Payload:    value,
	}
}

func NewIntArrayTag(name string, value []int32) *IntArrayTag {
	return &IntArrayTag{
		defaultTag: newDefaultTag(NBT_TAG_Int_Array, name),
		Payload:    value,
	}
}

func NewLongArrayTag(name string, value []int64) *LongArrayTag {
	return &LongArrayTag{
		defaultTag: newDefaultTag(NBT_TAG_Long_Array, name),
		Payload:    value,
	}
}
