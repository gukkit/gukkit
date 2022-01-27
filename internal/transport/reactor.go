package transport

import (
	"errors"
	"gukkit/internal/interfaces"
	"gukkit/internal/packet"
	"gukkit/internal/packet/handshaking"
	"gukkit/internal/packet/status"
	"gukkit/internal/session"
)

var (
	packetIDErr         = errors.New("packet id is not equal")
	unknownNextStatusID = errors.New("unknown next status id")
)

type Reactor struct {
	server interfaces.Server
}

func (reactor *Reactor) react(sess session.Session, recv *recvRawPacket) {
	var (
		reply packet.Clientbound
		err   error
	)

	switch sess.State() {
	case session.Handshaking:
		reply, err = reactor.handleHandshake(sess, recv)
	case session.Status:
		reply, err = reactor.handleStatus(sess.(*session.StatusSession), recv)
	case session.Login:

	case session.Playing:

	}

	if err != nil {
		sess.Close()
		return
	}

	if reply != nil {
		if err = sess.SendPacket(reply); err != nil {
			sess.Close()
			return
		}
	}
}

func (reactor *Reactor) handleHandshake(sess session.Session, recv *recvRawPacket) (reply packet.Clientbound, err error) {
	var pk handshaking.HandshakePacket

	if recv.ID != handshaking.HandshakePacketID {
		return nil, packetIDErr
	}

	if err = recv.ConvPacket(&pk); err != nil {
		return
	}

	switch pk.NextState {
	case 1: //conv status session
		sess.NextState(session.Status, reactor.server)

	case 2: //conv login session
		sess.NextState(session.Login, reactor.server)
	default:
		err = unknownNextStatusID
	}

	return
}

func (reactor *Reactor) handleStatus(sess *session.StatusSession, recv *recvRawPacket) (reply packet.Clientbound, err error) {

	switch recv.ID {
	case status.PingPongID:
		var pk status.PingPongPacket
		if err = recv.ConvPacket(&pk); err != nil {
			return
		}

		reply, err = sess.Ping(&pk)
	case status.RequestID:
		var pk status.RequestPacket
		if err = recv.ConvPacket(&pk); err != nil {
			return
		}

		reply, err = sess.Request(&pk)
	}

	return nil, errors.New("no handler")
}

func (reactor *Reactor) handleLogin(recv *recvRawPacket) (reply packet.Clientbound, err error) {

	return
}
