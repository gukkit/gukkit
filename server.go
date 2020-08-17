package gukkit

import (
	"fmt"
	"gukkit/internal"
	"gukkit/internal/transport"

	"github.com/panjf2000/gnet"
)

const (
	Version  = ""
	Protocol = 49
)

type Server struct {
	internal.Motd
	players   []*Player
	port      int
	maxPlayer int
}

func NewServer(address string) (svr *Server, err error) {
	svr = &Server{}

	return svr, nil
}

func (server *Server) Listen(address string) error {
	fmt.Println(Logo)
	options := []gnet.Option{
		gnet.WithMulticore(true),
		gnet.WithCodec(&transport.Codec{}),
	}

	t := &transport.Transporter{}
	return t.RunService(address)
}

func (server *Server) Players() (players []*Player) {
	players = make([]*Player, 0, len(server.players))
	copy(server.players, players)

	return
}

func (server *Server) Shutdown() error {
	return nil
}
