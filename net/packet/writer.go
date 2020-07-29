package packet

import "io"

//发送给客户端的包
type Clientbound interface {
	Encode(w Writer) (err error)
}

//发送给服务端的包
type Serverbound interface {
	Decode(r Reader) (err error)
}

type Writer interface {
	packet.
	io.Writer
	io.ByteWriter
}

type Reader interface {
	io.Reader
	io.ByteReader
}
