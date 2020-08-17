package transport

import (
	"fmt"
	"gukkit/internal/session"

	"github.com/panjf2000/gnet"
	"github.com/valyala/bytebufferpool"
)

var (
	BufferPool bytebufferpool.Pool
)

type Transporter struct {
	*gnet.EventServer
}

func (t *Transporter) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	c.SetContext(OpenInitSession(c))
	return
}

func (t *Transporter) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	session := c.Context().(*Session)
	session.Reset()

	SessionPool.Put(session)
	return
}

func (t *Transporter) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	var (
		recvRaw = recvRawPacketPool.Get().(*recvRawPacket)
		sess    = c.Context().(*Session)
		err     error
	)
	defer recvRawPacketPool.Put(recvRaw)
	fmt.Println(frame)

	if err = recvRaw.Unpack(frame, sess.Compressed); err != nil {
		sess.Close()
		return
	}

	switch sess.State {
	case Handshaking:

	case Status:
		sSession := sess.(*session.StatusSession)

		sSession.Ping(nil)
	case Login:

	case Playing:

	}

	return
}

func (t *Transporter) RunService(address string) (err error) {
	options := []gnet.Option{
		gnet.WithMulticore(true),
		gnet.WithCodec(&transport.Codec{}),
	}

	err = gnet.Serve(&transport.Transporter{}, address, options...)
	return
}
