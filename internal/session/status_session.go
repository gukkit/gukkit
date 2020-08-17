package session

import (
	"encoding/json"

	"gukkit/internal/interfaces"
	"gukkit/internal/packet/status"
	"gukkit/net/data/types"
	"gukkit/net/packet"
	"unsafe"
)

type StatusSession struct {
	server interfaces.Server
	*Session
}

// return PingPongPacket
func (session *StatusSession) Ping(pk packet.Serverbound) (reply packet.Clientbound, err error) {
	reply = pk.(*status.PingPongPacket)
	return
}

// return ResponsePacket
func (session *StatusSession) Request(pk packet.Serverbound) (reply packet.Clientbound, err error) {
	result, err := json.Marshal(session.server.Motd())

	if err != nil {
		return
	}

	reply = status.ResponsePacket{
		Message: types.String(*(*string)(unsafe.Pointer(&result))),
	}

	return
}
