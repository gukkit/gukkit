package types

type ByteArray []byte

func (b ByteArray) Encode(w Writer) (err error) {
	if err = VarInt(len(b)).Encode(w); err != nil {
		return
	}

	_, err = w.Write(b)
	return
}

func (b *ByteArray) Decode(r Reader) (n int, err error) {
	var len VarInt

	if n, err = len.Decode(r); err != nil {
		return
	}

	buf := make([]byte, len)

	if n, err = r.Read(buf); err != nil {
		return
	}

	*b = ByteArray(buf)
	return
}
