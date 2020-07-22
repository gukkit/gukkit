package types

import "errors"

type Int int32

var readIntLenErr = errors.New("reader result byte len < 4")

func (i Int) Encode(w Writer) (err error) {
	_, err = w.Write([]byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)})

	return
}

func (i Int) Decode(r Reader) (n int, err error) {
	data := make([]byte, 4)
	n, err = r.Read(data)
	if n != 4 {
		err = readIntLenErr
		return
	}

	i = Int(int(data[0])<<24 | int(data[1])<<16 | int(data[2])<<16 | int(data[3]))

	return 4, nil
}
