package types

import (
	"io"
)

type Type interface {
	TypeEncoder
	TypeDecoder
}

type TypeEncoder interface {
	Encode(w Writer) (err error)
}

type TypeDecoder interface {
	Decode(r Reader) (n int, err error)
}

type TypeReader struct {
	raw      []byte
	rawIndex int
}

type Writer interface {
	io.Writer
	io.ByteWriter
}

type Reader interface {
	io.Reader
	io.ByteReader
}
