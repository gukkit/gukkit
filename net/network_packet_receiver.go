package net

import (
	"fmt"
//	"time"

	"errors"
	"gukkit/net/packet"
	"github.com/panjf2000/gnet"
)

type NetworkPacketReceiver struct {
	gnet.EventServer

	queue []*packet.Packet
}

func(receiver *NetworkPacketReceiver) OnInitComplete(svr gnet.Server) (action gnet.Action) {
	return
}

func(receiver *NetworkPacketReceiver) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	return
}

func(receiver *NetworkPacketReceiver) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	return
}

func(receiver *NetworkPacketReceiver) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	pk, err := receiver.RecvPacket(frame)
	if err != nil {
		c.Close()
		return
	}

	if context := c.Context().(ConnContext); context.State() == Handshake {
		if pk.ID() == 0 {

		}
	}

	fmt.Println("[React]", pk, frame)
	return
}

func(receiver *NetworkPacketReceiver) RecvPacket(pkBytes []byte) (pk packet.Packet, err error) {
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
