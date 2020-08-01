package transport

import (
	"gukkit/internal/packet/status"
)

type StatusViewer struct {
	session *Session
}

func (viewer *StatusViewer) Response(req *status.RequestPacket) (err error) {
	return
}

func (viewer *StatusViewer) Pong(pong *status.PingPongPacket) (err error) {
	err = viewer.session.SendPacket(pong)
	return
}
