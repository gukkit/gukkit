package types

type UnsignedByte uint8

func (ub UnsignedByte) Encode(w Writer) (err error) {
	err = w.WriteByte(byte(ub))
	return
}

func (ub UnsignedByte) Decode(r Reader) (n int, err error) {
	data, err := r.ReadByte()
	if err != nil {
		return 0, err
	}

	n = 1
	ub = UnsignedByte(data)
	return
}
