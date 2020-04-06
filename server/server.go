package server

import (
	"gukkit/net"
	"github.com/panjf2000/gnet"
)

type Server struct {
	//*net.Listener

	players	[]*Player
}

func NewServer(address string) (server *Server, err error){
	server = new(Server)

	// listener, err := net.Listen(server, address)
	// if err != nil {
	// 	return nil, err
	// }

	// server.Accepter()

	return server, nil
}

func(server *Server) Run(address string) error {
	return gnet.Serve(&net.NetworkEventHandler{}, address)
}

func(server *Server) Players() []*Player {
	return server.players
}

func(server *Server) Close() error {
	return nil
}

func(server *Server) OuputMotd() {

}
