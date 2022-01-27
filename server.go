package gukkit

import (
	"fmt"
	"gukkit/internal"
	"gukkit/internal/transport"
)

const (
	Version  = ""
	Protocol = 49
)

type Server struct {
	motd      internal.Motd
	players   []*Player
	port      int
	maxPlayer int
}

func NewServer(address string) (svr *Server, err error) {
	svr = &Server{}

	return svr, nil
}

func (server *Server) Motd() internal.Motd {
	return server.motd
}

func (server *Server) Listen(address string) error {
	fmt.Println(Logo)
	return transport.Serve(server, address)
}

func (server *Server) Players() (players []*Player) {
	players = make([]*Player, 0, len(server.players))
	copy(server.players, players)

	return
}

func (server *Server) Shutdown() error {
	return nil
}
