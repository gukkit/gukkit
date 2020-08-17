package gukkit

import (
	"bytes"
	"gukkit/internal/transport"
	"gukkit/net/packet"
	"gukkit/text"
)

type Player struct {
	server *Server

	session  *transport.Session
	username string
}

func (player *Player) Username() string {
	return player.username
}

func (player *Player) SetTitle(text text.Chat) (err error) {
	var buf bytes.Buffer

	pk := &packet.TitlePacket{
		ActionType: packet.TitleActionID,
	}

	if err = pk.Encode(&buf); err != nil {
		return
	}

	err = player.session.SendPacket(pk)
	return
}

func (player *Player) Kick(reason text.Chat) (err error) {
	pk := packet.DisconnectPacket{}

	err = player.session.SendPacket(pk)
	return
}
