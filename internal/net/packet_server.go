package net

import (
	"gukkit/internal/intfac"
	"gukkit/internal/net/session"
	"gukkit/net"
	"sync"

	"github.com/panjf2000/gnet"
)

var dataPacketPool *sync.Pool = &sync.Pool{
	New: func() interface{} {
		return &net.DataPacket{}
	},
}

type PacketServer struct {
	*gnet.EventServer

	server intfac.Server
	notify <-chan *net.DataPacket
}

func (server *PacketServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	c.SetContext(session.OpenInitSession(c))
	return
}

func (server *PacketServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	curSession := c.Context().(*session.Session)

	if curSession.State == session.Playing {
		//玩家离开了游戏
	}
	return
}

func (server *PacketServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	dataPacket := dataPacketPool.Get().(*net.DataPacket)
	defer dataPacketPool.Put(dataPacket)

	session := c.Context().(*session.Session)

	if err := dataPacket.Unpack(frame, session.Compressed); err != nil {
		c.Close()
		return
	}

	server.onReact(dataPacket, session)
	return
}

func (server *PacketServer) onReact(dataPacket *net.DataPacket, session *session.Session) {

}
