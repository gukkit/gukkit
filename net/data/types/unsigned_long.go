package types

type UnsignedLong uint64

func (ulong UnsignedLong) Encode(w Writer) (err error) {
	b := []byte{
		byte(ulong >> 56),
		byte(ulong >> 48),
		byte(ulong >> 40),
		byte(ulong >> 32),
		byte(ulong >> 24),
		byte(ulong >> 16),
		byte(ulong >> 8),
		byte(ulong),
	}

	_, err = w.Write(b)
	return
}

func (ulong *UnsignedLong) Decode(r Reader) (n int, err error) {
	data := make([]byte, 8)

	if n, err = r.Read(data); n != 8 || err != nil {
		err = readIntLenErr
		return
	}

	*ulong = UnsignedLong(uint64(data[0])<<56 | uint64(data[1])<<48 | uint64(data[2])<<40 | uint64(data[3])<<32 | uint64(data[4])<<24 | uint64(data[5])<<16 | uint64(data[6])<<8 | uint64(data[7]))
	return 8, nil
}
