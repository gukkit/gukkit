package session

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
	Initial     state = 0
	Handshaking state = 1
	Play        state = 2
)

//玩家会话
type Session struct {
	manager *Manager

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
		State:      Initial,
		Conn:       conn,
		Compressed: false,
	}
}

func (session *Session) NextState(nextState state) (err error) {
	if session.State != Initial {
		err = dupHandshakeErr
		return
	}

	session.State = nextState
	return
}

func (session *Session) SendPacket(packet packet.Packet) (err error) {
	// data, err := packet.Encode()
	// if err != nil {
	// 	return err
	// }

	// if err = session.Conn.AsyncWrite(data); err != nil {
	// 	return
	// }

	return
}

func (session *Session) RecvPacket(packet packet.Packet) {

}

func (session *Session) Disconnect() (err error) {
	return
}
