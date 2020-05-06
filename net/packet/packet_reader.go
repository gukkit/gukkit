package packet

import (
	//"bytes"
	//"fmt"
	"errors"
	"encoding/binary"
	"math"
)

type PacketReader struct {
	data []byte
	idx  int
}

func NewPacketReader(data []byte) (pkReader *PacketReader, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			pkReader = nil
			err = errors.New(rec.(string))
			return
		}
	}()

	pkReader = &PacketReader{data: data}
	pkLen := int(pkReader.ReadVarInt())
	pkReader.data = pkReader.ReadBytes(pkLen)
	pkReader.idx = 0
	return
}

func(reader *PacketReader) ReadBoolean() bool {
	switch reader.ReadByte() {
	case 0x00:
		return false;
	case 0x01:
		return true;
	default:
		panic("read bool but the byte not is bool")
	}
}

func(reader *PacketReader) ReadByte() (b byte) {
	if reader.idx >= len(reader.data) {
		panic("no byte any more")
	}

	b = reader.data[reader.idx]
	reader.idx++
	return
}

func(reader *PacketReader) ReadUnsignedByte() uint8 {
	return uint8(reader.ReadByte())
}

func(reader *PacketReader) ReadShort() int16 {
	left, right := reader.ReadByte(), reader.ReadByte()

	return int16(left) << 8 | int16(right)
}

func(reader *PacketReader) ReadUnsignedShort() uint16 {
	left, right := reader.ReadByte(), reader.ReadByte()

	return uint16(int16(left) << 8 | int16(right))
}

func(reader *PacketReader) ReadInt() int32 {
	bs := reader.ReadBytes(4)

	return int32(bs[0]) << 24 | int32(bs[1]) << 16 | int32(bs[2]) << 8 | int32(bs[3])
}

func(reader *PacketReader) ReadLong() (result int64) {
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

func(reader *PacketReader) ReadFloat() float32 {
	bits := binary.LittleEndian.Uint32(reader.ReadBytes(4))
	return math.Float32frombits(bits)
}

func(reader *PacketReader) ReadDouble() float64 {
	bits := binary.LittleEndian.Uint64(reader.ReadBytes(4))
	return math.Float64frombits(bits)
}

func(reader *PacketReader) ReadString() string {
	len := reader.ReadVarInt() //string len

	return string(reader.ReadBytes(int(len)))
}

func(reader *PacketReader) ReadChat() {

}

func(reader *PacketReader) ReadIdentifier() {

}

func(reader *PacketReader) ReadVarInt() (result int32) {
	for i := 0; ;i++ { //读数据前的长度标记
		value := reader.ReadByte()

		result |= int32(value & 0b11111111) << int(7*i)

		if value&0b10000000 == 0 {
			break
		} else if i > 5 {
			panic("VarInt is too big")
		}
	}

	return result
}

func(reader *PacketReader) ReadVarLong() int64 {
	return 0
}

func(reader *PacketReader) ReadEntityMetadata() {

}

func(reader *PacketReader) ReadSlot() {

}

func(reader *PacketReader) ReadNbtTag() {

}

func(reader *PacketReader) ReadPosition() {

}

func(reader *PacketReader) ReadAngle() {

}

func(reader *PacketReader) ReadUUID() {

}

func(reader *PacketReader) ReadOptionalX() {

}

func(reader *PacketReader) ReadArrayOfX() {

}

func(reader *PacketReader) ReadXEnum() {

}

func(reader *PacketReader) ReadByteArray() {

}

func(reader *PacketReader) ReadBytes(n int) (b []byte) {
	if reader.idx + n > len(reader.data) {
		panic("no byte enough more")
	}

	b = reader.data[reader.idx:reader.idx+n]
	reader.idx += n
	return
}






