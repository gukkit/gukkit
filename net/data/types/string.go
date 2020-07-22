package types

type String string

func (str String) Encode(w Writer) (err error) {
	if err = VarInt(len(str)).Encode(w); err != nil {
		return err
	}

	_, err = w.Write([]byte(str))
	return
}

func (str *String) Decode(r Reader) (n int, err error) {
	var strLen VarInt

	if n, err = strLen.Decode(r); err != nil {
		return
	}

	buf := make([]byte, strLen)

	if n, err = r.Read(buf); err != nil {
		return
	}

	*str = String(buf)

	return 0, nil
}
