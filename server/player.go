package server

import (
	"gukkit/net"
)

type Player struct {
	*net.Conn
}
