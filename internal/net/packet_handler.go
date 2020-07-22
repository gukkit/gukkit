package net

import (
	"gukkit/internal/event"
	"gukkit/internal/net/session"
	"gukkit/net/packet"
)

type PacketHandler struct {
	pipeline *event.EventPipeline
}

func (handler *PacketHandler) Handle(session *session.Session, pk packet.Packet) {

	return
}
