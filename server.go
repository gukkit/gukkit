package gukkit

import (
	"fmt"
	//	"gukkit/net"

	"gukkit/internal/ticker"
	"gukkit/internal/transport"

	"github.com/panjf2000/gnet"
)

type Server interface {
	Listen(address string) error
}

type server struct {
	players   []*Player
	port      int
	maxPlayer int

	ticker *ticker.Ticker
}

func NewServer(address string) (svr Server, err error) {
	svr = &server{}

	return svr, nil
}

func (server *server) Listen(address string) error {
	fmt.Println(Logo)
	options := []gnet.Option{
		gnet.WithMulticore(true),
		gnet.WithCodec(&transport.Codec{}),
	}

	return gnet.Serve(&transport.Transporter{}, address, options...)
}

func (server *server) Players() []*Player {
	return server.players
}

func (server *server) Close() error {
	return nil
}
