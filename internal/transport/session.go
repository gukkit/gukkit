package transport

import (
	"errors"
	"gukkit/internal/packet"

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
	UUID     string
	Username string
	State    state
	Conn     gnet.Conn

	Compressed bool
	PrivateKey string
	PublicKey  string
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

func (session *Session) SendPacket(pk packet.Clientbound) (err error) {
	buffer := BufferPool.Get()
	buffer.Reset()
	defer BufferPool.Put(buffer)

	if err = pk.Encode(buffer); err != nil {
		return err
	}

	err = session.Conn.AsyncWrite(buffer.Bytes())
	return
}

func (session *Session) Reset(conn gnet.Conn) {
	session.UUID = ""
	session.Username = ""
	session.State = Handshaking
	session.Conn = conn
	session.Compressed = false
	session.PrivateKey = ""
	session.PublicKey = ""
}
