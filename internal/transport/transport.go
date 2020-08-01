package transport

import (
	"github.com/panjf2000/gnet"
	"github.com/valyala/bytebufferpool"
)

var (
	BufferPool bytebufferpool.Pool
)

type Transporter struct {
	*gnet.EventServer

	pkNode *RecvPacketNode
}

func (t *Transporter) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	c.SetContext(OpenInitSession(c))
	return
}

func (t *Transporter) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	//session := c.Context().(*Session)

	return
}

func (t *Transporter) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {

	return
}
