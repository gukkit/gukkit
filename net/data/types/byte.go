package types

type Byte byte

func (b Byte) Encode(w Writer) (err error) {
	err = w.WriteByte(byte(b))
	return
}

func (b Byte) Decode(r Reader) (n int, err error) {
	bytee, err := r.ReadByte()
	if err != nil {
		return 0, err
	}

	b = Byte(bytee)
	n = 1
	return
}
