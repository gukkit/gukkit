package types

type Type interface {
	Encode() (b []byte, err error) //编码成字节
	Decode(b []byte) (n int, err error)//从字节中解码
}

type TypeReader struct {
	raw				[]byte
	rawIndex 	int
}
