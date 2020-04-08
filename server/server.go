package server

import (
	"gukkit/net"
	"github.com/panjf2000/gnet"
)

type Server struct {
	receiver *net.NetworkPacketReceiver
	players	[]*Player
}

func NewServer(address string) (server *Server, err error){
	server = new(Server)

	return server, nil
}

func(server *Server) Listen(address string) error {
	return gnet.Serve(server.receiver, address)
}

func(server *Server) Players() []*Player {
	return server.players
}

func(server *Server) Close() error {
	return nil
}

func(server *Server) OuputMotd() {

}
