package interfaces

import "gukkit/internal"

type Server interface {
	Motd() internal.Motd
}
