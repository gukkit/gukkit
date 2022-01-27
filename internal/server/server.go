package server

import (
	"gukkit/entity"
	"gukkit/entity/player"
	"gukkit/net/packet"

	"time"
)

const DefaultInterval = time.Duration(50)

type Server struct {
	Interval   time.Duration
	isClosed   bool
	tick       uint64
	shutdown   chan struct{} //Channel for sending a "quit message" to the reader goroutine
	entities   []entity.Entity
	players    []*player.Player
	packetChan <-chan packet.Packet
}

func (server *Server) Loop() {
	for ; ; time.Sleep(server.Interval * time.Millisecond) {

		server.tick++
	}
}

func (server *Server) packet() {

}

func (server *Server) IsClosed() bool {
	select {
	case <-server.shutdown:
		return true
	default:
		return false
	}
}
