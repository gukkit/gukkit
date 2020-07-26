package types

import "math"

type Float float32

func (f Float) Encode(w Writer) (err error) {
	err = Long(math.Float32bits(float32(f))).Encode(w)
	return
}

func (d *Float) Decode(r Reader) (n int, err error) {
	var long Long

	if n, err = long.Decode(r); err != nil {
		return
	}

	*d = Float(math.Float32frombits(uint32(long)))
	return
}

type Double float64

func (d Double) Encode(w Writer) (err error) {
	err = Long(math.Float64bits(float64(d))).Encode(w)
	return
}

func (d *Double) Decode(r Reader) (n int, err error) {
	var long Long

	if n, err = long.Decode(r); err != nil {
		return
	}

	*d = Double(math.Float64frombits(uint64(long)))
	return
}
