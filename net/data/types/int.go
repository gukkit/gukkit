package types

type Int int32

func(v *Int) Encode() ([]byte, error) {
	i := *v

	return []byte{byte(i >> 24), byte(i >> 16), byte(i >> 8), byte(i)}, nil
}
