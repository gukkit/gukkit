package types

import (
	"errors"
)

var (
	VarIntTooBigErr = errors.New("VarInt is too big")
)

type VarInt int32

func (varInt VarInt) Encode(w Writer) (err error) {
	var value VarInt

	for {
		value = varInt & 0x7F

		if varInt >>= 7; varInt != 0 {
			value |= 0x80
		}

		err = w.WriteByte(byte(value))

		if varInt == 0 {
			break
		}
	}

	return
}

func (varInt *VarInt) Decode(r Reader) (n int, err error) {
	var result int32

	for { //读数据前的长度标记
		value, err := r.ReadByte()
		if err != nil {
			return n, err
		}
		result |= int32(value&0b11111111) << int32(7*n)

		n++
		if value&0b10000000 == 0 {
			break
		} else if n > 5 {
			return 0, VarIntTooBigErr
		}
	}

	*varInt = VarInt(result)
	return
}
