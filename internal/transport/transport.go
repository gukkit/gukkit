package transport

import (
	"gukkit/internal/interfaces"
	"gukkit/internal/session"

	"github.com/panjf2000/gnet"
	"github.com/valyala/bytebufferpool"
)

var (
	BufferPool bytebufferpool.Pool
)

type Transporter struct {
	*gnet.EventServer

	reactor *Reactor
}

func Serve(server interfaces.Server, address string) (err error) {
	transporter := Transporter{
		reactor: &Reactor{
			server: server,
		},
	}

	err = transporter.RunService(address)
	return
}

func (t *Transporter) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	c.SetContext(session.OpenInitSession(c))
	return
}

func (t *Transporter) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	session := c.Context().(session.Session)
	session.Release()
	return
}

func (t *Transporter) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	var (
		recvRaw = recvRawPacketPool.Get().(*recvRawPacket)
		sess    = c.Context().(session.Session)
		err     error
	)

	defer recvRawPacketPool.Put(recvRaw)

	if err = recvRaw.Unpack(frame, sess.Compressed()); err != nil {
		sess.Close()
		return
	}

	t.reactor.react(sess, recvRaw)
	return
}

func (t *Transporter) RunService(address string) (err error) {
	options := []gnet.Option{
		gnet.WithMulticore(true),
		gnet.WithCodec(&Codec{}),
	}

	err = gnet.Serve(&Transporter{}, address, options...)
	return
}
