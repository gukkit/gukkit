package gukkit

import (
	"gukkit/internal/net/session"
)

type Player struct {
	session  *session.Session
	username string
}

func (player *Player) Username() string {
	return player.username
}
