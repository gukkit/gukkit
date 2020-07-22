package types

type UnsignedShort uint16

func (short UnsignedShort) Encode(w Writer) (err error) {
	err = w.WriteByte(byte(short << 8))
	err = w.WriteByte(byte(short))
	return
}

func (short UnsignedShort) Decode(r Reader) (n int, err error) {
	var (
		s1 byte //stream short 1
		s2 byte //stream short 2
	)

	s1, err = r.ReadByte()
	s2, err = r.ReadByte()
	if err != nil {
		return 0, err
	}

	short = UnsignedShort(uint16(s1)<<8 | uint16(s2))
	return
}
