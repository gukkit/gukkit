package session

import (
	"errors"
	"gukkit/internal/interfaces"
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

type Session interface {
	State() state
	NextState(nextState state, server interfaces.Server) (err error)
	Compressed() bool
	UUID() string
	Username() string
	SendPacket(pk packet.Clientbound) (err error)
	Release()

	Write(b []byte) (n int, err error)
	WriteByte(b byte) (err error)
	Close() (err error)
}

//玩家会话
type session struct {
	buf      []byte
	uuid     string
	username string
	state    state
	Conn     gnet.Conn

	compressed bool
	PrivateKey string
	PublicKey  string
}

func OpenInitSession(conn gnet.Conn) Session {
	sess, _ := SessionPool.Get().(*session)

	if sess == nil {
		sess = &session{
			state:      Handshaking,
			Conn:       conn,
			compressed: false,
		}
	} else {
		sess.Conn = conn
	}

	return sess
}

func (session *session) Username() string {
	return session.username
}

func (session *session) UUID() string {
	return session.uuid
}

func (session *session) State() state {
	return session.state
}

func (session *session) Compressed() bool {
	return session.compressed
}

func (session *session) NextState(nextState state, server interfaces.Server) (err error) {
	if session.state != Handshaking {
		err = dupHandshakeErr
		return
	}

	if nextState == Status {
		statusSession := &StatusSession{
			session: session,
			server:  server,
		}

		session.Conn.SetContext(statusSession)
	}

	if nextState == Login {
		//todo
	}

	session.state = nextState
	return
}

func (session *session) Write(b []byte) (n int, err error) {
	session.buf = append(session.buf, b...)
	return len(b), nil
}

func (session *session) WriteByte(b byte) (err error) {
	session.buf = append(session.buf, b)
	return nil
}

func (session *session) SendPacket(pk packet.Clientbound) (err error) {
	if err = pk.Encode(session); err != nil {
		return err
	}

	err = session.Conn.AsyncWrite(session.buf)
	return
}

func (session *session) Close() (err error) {
	err = session.Conn.Close()
	return
}

func (session *session) Reset() {
	session.uuid = ""
	session.username = ""
	session.state = Handshaking
	session.Conn = nil
	session.compressed = false
	session.PrivateKey = ""
	session.PublicKey = ""
}

func (session *session) Release() {
	session.Reset()

	SessionPool.Put(session)
}
