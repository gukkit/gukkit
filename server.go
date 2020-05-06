package gukkit

import (
	"fmt"
//	"gukkit/net"
	"gukkit/internal/net"
	"github.com/panjf2000/gnet"


)




type Server interface {
	Listen(address string) error
}

type server struct {

	players	[]*Player
	port 			int
	maxPlayer int
}

func NewServer(address string) (server Server, err error){
	server = new(server)

	//net.PacketHandker{}

	return server, nil
}

func(server *server) Listen(address string) error {
	fmt.Println(Logo)
	return gnet.Serve(&net.PacketHandler{}, address, gnet.WithMulticore(true), gnet.WithCodec(net.DefaultPackCodec))
}

func(server *server) Players() []*Player {
	return server.players
}

func(server *server) Close() error {
	return nil
}
