package player

import (
	"bytes"
	"gukkit/internal/transport"
	"gukkit/net/packet"
	"gukkit/text"
)

type Player struct {
	session *transport.Session
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
