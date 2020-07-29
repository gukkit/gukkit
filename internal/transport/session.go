package transport

import (
	"bytes"
	"errors"
	"gukkit/internal/packet"
	"sync"

	"github.com/panjf2000/gnet"
)

type state int32

var (
	dupHandshakeErr = errors.New("No duplication handshake")

	BufferPool = &sync.Pool{
		New: func() interface{} {
			buf := make([]byte, 128)

			return bytes.NewBuffer(buf)
		},
	}
)

const (
	Handshaking state = 0
	Status      state = 1
	Login       state = 2
	Playing     state = 3
)

//玩家会话
type Session struct {
	PrivateKey string

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

func (session *Session) SendPacket(pk packet.Clientbound) (err error) {
	buffer := BufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer BufferPool.Put(buffer)

	if err = pk.Encode(buffer); err != nil {
		return err
	}

	err = session.Conn.AsyncWrite(buffer.Bytes())
	return
}
