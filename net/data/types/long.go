package types

type Long int64

func (long Long) Encode(w Writer) (err error) {
	b := []byte{
		byte(long >> 56),
		byte(long >> 48),
		byte(long >> 40),
		byte(long >> 32),
		byte(long >> 24),
		byte(long >> 16),
		byte(long >> 8),
		byte(long),
	}

	_, err = w.Write(b)
	return
}

func (long *Long) Decode(r Reader) (n int, err error) {
	data := make([]byte, 8)

	if n, err = r.Read(data); n != 8 || err != nil {
		err = readIntLenErr
		return
	}

	res := int64(data[0]) << 56
	res |= int64(data[1]) << 48
	res |= int64(data[2]) << 40
	res |= int64(data[3]) << 32
	res |= int64(data[4]) << 24
	res |= int64(data[5]) << 16
	res |= int64(data[6]) << 8
	res |= int64(data[7])

	*long = Long(res)
	return 8, nil
}
