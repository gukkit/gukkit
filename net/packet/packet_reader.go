package packet

import (
	//"bytes"
	//"errors"

	"encoding/binary"
	"math"
)

type PacketDataReader struct {
	data []byte
	idx  int
}

func(reader *PacketDataReader) ReadBoolean() bool {
	switch reader.ReadByte() {
	case 0x00:
		return false;
	case 0x01:
		return true;
	default:
		panic("read bool but the byte not is bool")
	}
}

func(reader *PacketDataReader) ReadByte() (b byte) {
	if reader.idx >= len(reader.data) {
		panic("no byte any more")
	}

	b = reader.data[reader.idx]
	reader.idx++
	return
}

func(reader *PacketDataReader) ReadUnsignedByte() uint8 {
	return uint8(reader.ReadByte())
}

func(reader *PacketDataReader) ReadShort() int16 {
	left, right := reader.ReadByte(), reader.ReadByte()

	return int16(left) << 8 | int16(right)
}

func(reader *PacketDataReader) ReadUnsignedShort() uint16 {
	left, right := reader.ReadByte(), reader.ReadByte()

	return uint16(int16(left) << 8 | int16(right))
}

func(reader *PacketDataReader) ReadInt() int32 {
	bs := reader.ReadBytes(4)

	return int32(bs[0]) << 24 | int32(bs[1]) << 16 | int32(bs[2]) << 8 | int32(bs[3])
}

func(reader *PacketDataReader) ReadLong() (result int64) {
	bs := reader.ReadBytes(8)

	for i := 0; i < 7; i++ {
		result += int64(bs[i]) << ((7-i) * 8)
	}

	// result += int64(bs[0]) << 56
	// result += int64(bs[1]) << 48
	// result += int64(bs[2]) << 40
	// result += int64(bs[3]) << 32
	// result += int64(bs[4]) << 24
	// result += int64(bs[5]) << 16
	// result += int64(bs[6]) << 8
	// result += int64(bs[7])
	return
}

func(reader *PacketDataReader) ReadFloat() float32 {
	bits := binary.LittleEndian.Uint32(reader.ReadBytes(4))
	return math.Float32frombits(bits)
}

func(reader *PacketDataReader) ReadDouble() float64 {
	bits := binary.LittleEndian.Uint64(reader.ReadBytes(4))
	return math.Float64frombits(bits)
}

func(reader *PacketDataReader) ReadString() string {
	len := reader.ReadVarInt() //string len

	return string(reader.ReadBytes(len))
}

func(reader *PacketDataReader) ReadChat() {

}

func(reader *PacketDataReader) ReadIdentifier() {

}

func(reader *PacketDataReader) ReadVarInt() (result int) {
	for i := 0; ;i++ { //读数据前的长度标记
		value := reader.ReadByte()

		result |= int(value & 0b11111111) << int(7*i)

		if value&0b10000000 == 0 {
			break
		} else if i > 5 {
			panic("VarInt is too big")
		}
	}

	return result
}

func(reader *PacketDataReader) ReadVarLong() int64 {
	return 0
}

func(reader *PacketDataReader) ReadEntityMetadata() {

}

func(reader *PacketDataReader) ReadSlot() {

}

func(reader *PacketDataReader) ReadNbtTag() {

}

func(reader *PacketDataReader) ReadPosition() {

}

func(reader *PacketDataReader) ReadAngle() {

}

func(reader *PacketDataReader) ReadUUID() {

}

func(reader *PacketDataReader) ReadOptionalX() {

}

func(reader *PacketDataReader) ReadArrayOfX() {

}

func(reader *PacketDataReader) ReadXEnum() {

}

func(reader *PacketDataReader) ReadByteArray() {

}


func(reader *PacketDataReader) ReadBytes(n int) (b []byte) {
	if reader.idx + n > reader.idx {
		panic("no byte enough more")
	}

	b = reader.data[reader.idx:reader.idx+n]
	reader.idx += n
	return
}






