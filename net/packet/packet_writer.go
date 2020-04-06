package packet

import (
	"io"
)

type PacketWriter struct {
	receiver io.Writer
	buf			 []byte
}

func(pw *PacketWriter) Write(data []byte) {
	pw.buf = append(pw.buf, data...)
}

func(pw *PacketWriter) WriteByte(b byte) {
	pw.buf = append(pw.buf, b)
}

func(pw *PacketWriter) Flush() (err error) {
	if len(pw.buf) == 0 {
		return nil
	}
	bufLen, n, m := len(pw.buf), 0, 0

	for bufLen >= n {
		if m, err = pw.receiver.Write(pw.buf[n:bufLen]); err == nil {
			n += m
			continue
		}
		break
	}
	return
}

func(pw *PacketWriter) WriteVarInt(value int32) {
	var data []byte

	num := uint32(value)
	for {
		b := num & 0x7F
		num >>= 7
		if num != 0 {
			b |= 0x80
		}
		data = append(data, byte(b))
		if num == 0 {
			break
		}
	}

	pw.Write(data)
}

func(pw *PacketWriter) WriteVarLong(value int64) {
	var data []byte

	num := uint64(value)
	for {
		b := num & 0x7F
		num >>= 7
		if num != 0 {
			b |= 0x80
		}
		data = append(data, byte(b))
		if num == 0 {
			break
		}
	}

	pw.Write(data)
}

var (
	Ture 	= []byte{1}
	False = []byte{0x00}
)

func(pw *PacketWriter) WriteBoolean(b bool) {
	if b {
		pw.receiver.Write(Ture)
		return
	}

	pw.receiver.Write(False)
	return
}

func(pw *PacketWriter) WriteString(str string) {
	pw.WriteVarInt(int32(len(str)))
	pw.Write([]byte(str))
}

func(pw *PacketWriter) WriteUnsignedByte(ub uint8) {
	pw.WriteByte(byte(ub))
}

func(pw *PacketWriter) WriteUnsignedShort(us uint16) {
	pw.Write([]byte{
		byte(us >> 8),
		byte(us),
	})
}

func(pw *PacketWriter) WriteInt(n int32) {
	pw.Write([]byte{byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n)})
}

func(pw *PacketWriter) WriteLong(n int64) {
	pw.Write([]byte{
		byte(n >> 56), byte(n >> 48), byte(n >> 40), byte(n >> 32),
		byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n),
	})
}
