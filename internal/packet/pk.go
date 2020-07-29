package packet

import (
	pk "gukkit/net/packet"
	"io"
)

//发送给客户端的包
type Clientbound interface {
	Encode(w pk.Writer) (err error)
}

//发送给服务端的包
type Serverbound interface {
	Decode(r pk.Reader) (err error)
}

type Writer interface {
	io.Writer
	io.ByteWriter
}

type Reader interface {
	io.Reader
	io.ByteReader
}
