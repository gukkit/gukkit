package net

import (
	"gukkit/net/packet"
	"github.com/panjf2000/gnet"
)

type Conn struct {
	pkWriter *packet.PacketWriter

	conn gnet.Conn //TCPConn
}

func newConn(conn gnet.Conn) *Conn {
	return &Conn {
		conn: conn,
		//pkWriter: packet.NewPacketWriter(conn),
	}
}

func(conn *Conn) WritePacket(pk packet.Packet) error {
	return nil
}
