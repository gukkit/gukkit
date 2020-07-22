package types

import "fmt"

type Position struct {
	X int
	Y int
	Z int
}

func (pos Position) Encode(w Writer) (err error) {
	fmt.Println(pos.Y & 0xFFF)
	err = UnsignedLong((pos.X&0x3FFFFFF)<<38 | (pos.Z&0x3FFFFFF)<<12 | pos.Y&0xFFF).Encode(w)
	return
}

func (pos *Position) Decode(r Reader) (n int, err error) {
	var long UnsignedLong

	if n, err = long.Decode(r); err != nil {
		return
	}

	pos.X = int(long>>38) & 0x3FFFFFF
	pos.Y = int(long & 0xFFF)
	pos.Z = int(long<<26>>38) & 0x3FFFFFF

	if pos.X >= 1<<25 {
		pos.X -= 1 << 26
	}

	if pos.Y >= 1<<11 {
		pos.Y -= 1 << 12
	}

	if pos.Z >= 1<<25 {
		pos.Z -= 1 << 26
	}
	return
}
