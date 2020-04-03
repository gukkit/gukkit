package packet

import (
	"github.com/google/uuid"
)

type (
	Boolean bool

	Byte int8

	UnsignedByte uint8

	Short int16

	UnsignedShort uint16

	Int int32

	Long int64

	Float float32

	Double float64

	String string

	Chat = string

	Identifier = String

	VarInt int32

	VarLong int64

	Position struct {
		X, Y, Z int
	}

	Angle int8

	UUID uuid.UUID

	Nbt struct {
		V interface{}
	}

	ByteArray []byte
)
