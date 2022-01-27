package gukkit

import (
	"bytes"
	"gukkit/internal/session"
	"gukkit/net/packet"
	"gukkit/text"
)

type Player struct {
	server *Server

	x float64
	y float64
	z float64

	username string
	session  session.Session
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
