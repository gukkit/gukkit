package gukkit

import (
	"fmt"
	"gukkit/net"
	"github.com/panjf2000/gnet"
)

const Logo = `
        GGGGGGGGGGGGG                  kkkkkkkk           kkkkkkkk             iiii          tttt
     GGG::::::::::::G                  k::::::k           k::::::k            i::::i      ttt:::t
   GG:::::::::::::::G                  k::::::k           k::::::k             iiii       t:::::t
  G:::::GGGGGGGG::::G                  k::::::k           k::::::k                        t:::::t
 G:::::G       GGGGGGuuuuuu    uuuuuu   k:::::k    kkkkkkk k:::::k    kkkkkkkiiiiiiittttttt:::::ttttttt
G:::::G              u::::u    u::::u   k:::::k   k:::::k  k:::::k   k:::::k i:::::it:::::::::::::::::t
G:::::G              u::::u    u::::u   k:::::k  k:::::k   k:::::k  k:::::k   i::::it:::::::::::::::::t
G:::::G    GGGGGGGGGGu::::u    u::::u   k:::::k k:::::k    k:::::k k:::::k    i::::itttttt:::::::tttttt
G:::::G    G::::::::Gu::::u    u::::u   k::::::k:::::k     k::::::k:::::k     i::::i      t:::::t
G:::::G    GGGGG::::Gu::::u    u::::u   k:::::::::::k      k:::::::::::k      i::::i      t:::::t
G:::::G        G::::Gu::::u    u::::u   k:::::::::::k      k:::::::::::k      i::::i      t:::::t
 G:::::G       G::::Gu:::::uuuu:::::u   k::::::k:::::k     k::::::k:::::k     i::::i      t:::::t    tttttt
  G:::::GGGGGGGG::::Gu:::::::::::::::uuk::::::k k:::::k   k::::::k k:::::k   i::::::i     t::::::tttt:::::t
   GG:::::::::::::::G u:::::::::::::::uk::::::k  k:::::k  k::::::k  k:::::k  i::::::i     tt::::::::::::::t
     GGG::::::GGG:::G  uu::::::::uu:::uk::::::k   k:::::k k::::::k   k:::::k i::::::i       tt:::::::::::tt
        GGGGGG   GGGG    uuuuuuuu  uuuukkkkkkkk    kkkkkkkkkkkkkkk    kkkkkkkiiiiiiii         ttttttttttt
`

type Server struct {
	receiver *net.NetworkPacketReceiver
	players	[]*Player
}

func NewServer(address string) (server *Server, err error){
	server = new(Server)

	return server, nil
}

func(server *Server) Listen(address string) error {
	fmt.Println(Logo)
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
