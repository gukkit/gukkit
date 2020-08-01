package packet

import (
	"errors"
	"io"
)

var (
	DecodeIDNotEqualErr = errors.New("decode packet id not equal err")
)

//发送给客户端的包
type Clientbound interface {
	Encode(w Writer) (err error)
}

//发送给服务端的包
type Serverbound interface {
	Decode(r Reader) (err error)
}

type Writer interface {
	io.Writer
	io.ByteWriter
}

type Reader interface {
	io.Reader
	io.ByteReader
}
