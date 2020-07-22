package types

import "errors"

var (
	boolTrue  byte = 1
	boolFalse byte = 0

	boolReadErr = errors.New("read byte is not a boolean type")
)

type Boolean bool

func (boolean Boolean) Encode(w Writer) (err error) {
	if boolean == true {
		err = w.WriteByte(boolTrue)
	} else {
		err = w.WriteByte(boolFalse)
	}

	return
}

func (boolean Boolean) Decode(r Reader) (n int, err error) {
	data, err := r.ReadByte()
	if err != nil {
		return 0, err
	}

	switch data {
	case boolTrue:
		boolean = true
	case boolFalse:
		boolean = false
	default:
		n = 0
		err = boolReadErr
		return
	}

	n = 1
	return
}
