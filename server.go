package gukkit

import (
	"fmt"
	"errors"
	"gukkit/net/packet"
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
	*gnet.EventServer
	players	[]*Player
	port 			int
	maxPlayer int
}

func NewServer(address string) (server *Server, err error){
	server = new(Server)

	return server, nil
}

func(server *Server) Listen(address string) error {
	fmt.Println(Logo)
	return gnet.Serve(server, address)
}

func(server *Server) Players() []*Player {
	return server.players
}

func(server *Server) Close() error {
	return nil
}

func(server *Server) OuputMotd() {

}

func(server *Server) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	c.SetContext(&context{state: HandshakeStatus})
	return
}

func(server *Server) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	fmt.Println("[OnClosed]", c, err)
	return
}

func(server *Server) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	pk, err := server.RecvPacket(frame)
	fmt.Println("[React]", pk, err, frame)
	if err != nil {
		c.Close()
		return
	}

	context := c.Context().(*context);

	switch context.Value("state") {
	case HandshakeStatus:
		switch pk.ID() {
		case 0:

		}

	case ServerListStatus:

	case PlayingStatus:

	}
	return
}

func(server *Server) RecvPacket(pkBytes []byte) (pk packet.Packet, err error) {
	pkReader, err := packet.NewPacketReader(pkBytes)
	if err != nil {
		return nil, err
	}

	switch pkReader.ReadVarInt() {
	case 0x00:
		pk, err = packet.EncodeServerListPingPacket(pkReader)
	default:
		return nil, errors.New("can find this packet")
	}
	return
}