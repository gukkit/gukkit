package session

import (
	"errors"
	"gukkit/internal/packet"
	"sync"

	"github.com/panjf2000/gnet"
)

type state int32

var (
	dupHandshakeErr = errors.New("No duplication handshake")

	SessionPool sync.Pool
)

const (
	Handshaking state = 0
	Status      state = 1
	Login       state = 2
	Playing     state = 3
)

//玩家会话
type Session struct {
	buf      []byte
	UUID     string
	Username string
	State    state
	Conn     gnet.Conn

	Compressed bool
	PrivateKey string
	PublicKey  string
}

func OpenInitSession(conn gnet.Conn) (session *Session) {

	if session, _ = SessionPool.Get().(*Session); session == nil {
		session = &Session{
			State:      Handshaking,
			Conn:       conn,
			Compressed: false,
		}
	} else {
		session.Conn = conn
	}

	return
}

func (session *Session) NextState(nextState state) (err error) {
	if session.State != Handshaking {
		err = dupHandshakeErr
		return
	}

	session.State = nextState
	return
}

func (session *Session) Write(b []byte) (n int, err error) {
	session.buf = append(session.buf, b...)
	return len(b), nil
}

func (session *Session) WriteByte(b byte) (err error) {
	session.buf = append(session.buf, b)
	return nil
}

func (session *Session) SendPacket(pk packet.Clientbound) (err error) {
	if err = pk.Encode(session); err != nil {
		return err
	}

	err = session.Conn.AsyncWrite(session.buf)
	return
}

func (session *Session) Close() (err error) {
	err = session.Conn.Close()
	return
}

func (session *Session) Reset() {
	session.UUID = ""
	session.Username = ""
	session.State = Handshaking
	session.Conn = nil
	session.Compressed = false
	session.PrivateKey = ""
	session.PublicKey = ""
}
