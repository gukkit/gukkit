package net

import (
	"gukkit/internal/net/session"
	"gukkit/net"
	"gukkit/net/packet"
)

type PacketParser struct {
}

func (parser *PacketParser) Parse(sess *session.Session, data *net.DataPacket) (pk packet.Packet, err error) {
	switch sess.State {
	case session.Initial, session.Handshaking:

	}
	return
}

type PlayPacketParser struct {
}

func (parser *PlayPacketParser) Parse(sess *session.Session, data *net.DataPacket) (pk packet.Packet, err error) {
	return
}
