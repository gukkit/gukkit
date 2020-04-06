package net

import (
	"fmt"
	"time"
	"gukkit/net/packet"
	"github.com/panjf2000/gnet"
)

type NetworkEventHandler struct {

}

func(handler *NetworkEventHandler) OnInitComplete(server gnet.Server) (action gnet.Action) {
	fmt.Println("[OnInitComplete]", server)
	return gnet.None
}

func(handler *NetworkEventHandler) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {

	return
}

func(handler *NetworkEventHandler) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	fmt.Println("[OnClosed]", c, err)
	return
}

func(handler *NetworkEventHandler) PreWrite() {

}

func(handler *NetworkEventHandler) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {

	pk, err := packet.RecvPacket(frame)
	if err != nil {
		c.Close()
	}

	fmt.Println("[React]", pk, frame)
	//pk := packet.NewPacketReader(frame)
	return
}

func(handler *NetworkEventHandler) Tick() (delay time.Duration, action gnet.Action) {
	return
}
