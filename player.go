package gukkit

import (
	"gukkit/net"
)

type Player struct {
	conn net.Conn
}

func(player *Player) SendMessage() error {
	return nil
}

