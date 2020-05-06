package net

import (
	"sync"
	"gukkit/net"
	"gukkit/internal/intfac"
	"gukkit/internal/net/session"
	"github.com/panjf2000/gnet"
)

var dataPacketPool *sync.Pool = &sync.Pool {
	New: func() interface{} {
		return &net.DataPacket{}
	},
}

type PacketHandler struct {
	*gnet.EventServer

	server	intfac.Server
	notify 	<-chan *net.DataPacket
}

func (handler *PacketHandler) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	c.SetContext(session.OpenInitSession(c))
	return
}

func (handler *PacketHandler) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	curSession := getSession(c)

	if curSession.State == session.Playing {
		//玩家离开了游戏
	}
	return
}

func (handler *PacketHandler) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	dataPacket := dataPacketPool.Get().(*net.DataPacket)
	defer dataPacketPool.Put(dataPacket)

	session := c.Context().(*session.Session)

	if err := dataPacket.Unpack(frame, session.Compressed); err != nil {
		c.Close()
		return
	}

	handler.onReact(dataPacket, session)
	return
}

func (handler *PacketHandler) onReact(dataPacket *net.DataPacket, session *session.Session) {

}

func getSession(c gnet.Conn) *session.Session {
	return c.Context().(*session.Session)
}
