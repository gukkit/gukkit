package gukkit

import "gukkit/text"

var (
	defaultMotd = "Welcome Join Gukkit Server"
)

type Option func(server *Server)

type Config struct {
	MOTD text.Chat
	Port uint16
}
