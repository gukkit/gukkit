package types

type Angle int8

func (angle Angle) Encode(w Writer) (err error) {

	return
}

func (angle *Angle) Decode(r Reader) (n int, err error) {
	var data Byte

	if n, err = data.Decode(r); err != nil {
		return
	}

	*angle = Angle(data)
	return
}
