package nbt

import (
	"io"
	"math"
)

var (
	fixedEndTagBinary = []byte{0x00}
)

func (tag *EndTag) Encode(w io.Writer) {
	w.Write(fixedEndTagBinary)
	return
}

func (tag *ByteTag) Encode(w io.Writer) (err error) {
	if err = tag.onHeadBytes(w); err != nil {
		return
	}

	_, err = w.Write([]byte{tag.Payload})
	return
}

func (tag *ShortTag) Encode(w io.Writer) (err error) {
	if err = tag.onHeadBytes(w); err != nil {
		return
	}

	_, err = w.Write([]byte{
		byte(tag.Payload >> 8),
		byte(tag.Payload),
	})
	return
}

func (tag *IntTag) Encode(w io.Writer) (err error) {
	if err = tag.onHeadBytes(w); err != nil {
		return
	}

	_, err = w.Write([]byte{
		byte(tag.Payload >> 24),
		byte(tag.Payload >> 16),
		byte(tag.Payload >> 8),
		byte(tag.Payload),
	})
	return
}

func (tag *LongTag) Encode(w io.Writer) (err error) {
	if err = tag.onHeadBytes(w); err != nil {
		return
	}

	_, err = w.Write([]byte{
		byte(tag.Payload >> 56),
		byte(tag.Payload >> 48),
		byte(tag.Payload >> 40),
		byte(tag.Payload >> 32),
		byte(tag.Payload >> 24),
		byte(tag.Payload >> 16),
		byte(tag.Payload >> 8),
		byte(tag.Payload),
	})
	return
}

func (tag *FloatTag) Encode(w io.Writer) (err error) {
	if err = tag.onHeadBytes(w); err != nil {
		return
	}

	pwd := math.Float32bits(tag.Payload)
	_, err = w.Write([]byte{
		byte(pwd >> 24),
		byte(pwd >> 16),
		byte(pwd >> 8),
		byte(pwd),
	})
	return
}

func (tag *DoubleTag) Encode(w io.Writer) (err error) {
	if err = tag.onHeadBytes(w); err != nil {
		return
	}

	pwd := math.Float64bits(tag.Payload)
	_, err = w.Write([]byte{
		byte(pwd >> 54), byte(pwd >> 48), byte(pwd >> 40), byte(pwd >> 32),
		byte(pwd >> 24), byte(pwd >> 16), byte(pwd >> 8), byte(pwd),
	})
	return
}

func (tag *ByteArrayTag) Encode(w io.Writer) (err error) {
	if err = tag.onHeadBytes(w); err != nil {
		return
	}

	size := len(tag.Payload)

	w.Write(append(
		[]byte{
			byte(size >> 24),
			byte(size >> 16),
			byte(size >> 8),
			byte(size),
		},
		tag.Payload...,
	))

	return
}

func (tag *StringTag) Encode(w io.Writer) (err error) {
	if err = tag.onHeadBytes(w); err != nil {
		return
	}

	size := len(tag.Payload)

	w.Write(append(
		[]byte{
			byte(size >> 24),
			byte(size >> 16),
			byte(size >> 8),
			byte(size),
		},
		tag.Payload...,
	))

	return
}

func (tag *ListTag) Encode(w io.Writer) (err error) {
	if err = tag.onHeadBytes(w); err != nil {
		return
	}

	return
}
