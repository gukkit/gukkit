package net

import (
	"gukkit/internal/net/session"
	"gukkit/net/packet"

	"github.com/panjf2000/gnet"
)

type NetworkListener struct {
	*gnet.EventServer
	statePacketNofity chan packet.Packet
	gamePacketNotify  chan packet.Packet
}

func (listener *NetworkListener) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	c.SetContext(session.OpenInitSession(c))
	return
}

func (listener *NetworkListener) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	curSession := c.Context().(*session.Session)
	if curSession.State == session.Play {
		//玩家离开了游戏
	}
	return
}

func (listener *NetworkListener) React(data []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	dataPacket := dataPool.Get()
	defer dataPool.Put(dataPacket)

	session := c.Context().(*session.Session)

	if err := dataPacket.Unpack(data, session.Compressed); err != nil {
		c.Close()
		return
	}

	// pk, err := parsePK(dataPacket)
	// if err != nil {
	// 	c.Close()
	// 	return
	// }

	// listener.packetNotify <- pk

	return
}

func (listener *NetworkListener) notifyChannel() <-chan packet.Packet {
	return nil
}
