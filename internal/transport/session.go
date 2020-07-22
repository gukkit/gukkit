package transport

import (
	"errors"
	"gukkit/net/packet"

	"github.com/panjf2000/gnet"
)

type state int32

var (
	dupHandshakeErr = errors.New("No duplication handshake")
)

const (
	Handshaking state = 0
	Status      state = 1
	Login       state = 2
	Playing     state = 3
)

//玩家会话
type Session struct {
	UUID  string
	State state
	Conn  gnet.Conn

	Compressed bool
}

func OpenInitSession(conn gnet.Conn) *Session {
	if conn == nil {
		return nil
	}

	return &Session{
		State:      Handshaking,
		Conn:       conn,
		Compressed: false,
	}
}

func (session *Session) NextState(nextState state) (err error) {
	if session.State != Handshaking {
		err = dupHandshakeErr
		return
	}

	session.State = nextState
	return
}

func (session *Session) SendPacket(packet packet.Packet) (err error) {
	data, err := packet.Encode()
	if err != nil {
		return err
	}

	if err = session.Conn.AsyncWrite(data); err != nil {
		return
	}

	return
}

func (session *Session) RecvPacket(packet packet.Packet) {

}
