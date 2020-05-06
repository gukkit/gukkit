package session

import (
	"errors"
	"github.com/panjf2000/gnet"
)

type state	int32

var (
	dupHandshakeErr = errors.New("No duplication handshake")
)

const (
	Initial 	state = 0
	StatusQuo state = 1
	Playing		state = 2
)

//玩家会话
type Session struct {
	UUID 	string
	State	state
	Conn	gnet.Conn

	Compressed	bool
}

func OpenInitSession(conn gnet.Conn) *Session {
	if conn == nil {
		return nil
	}

	return &Session {
		State: Initial,
		Conn: conn,
		Compressed: false,
	}
}

func (session *Session) Handshake(nextState state) (err error) {
	if session.State != Initial {
		err = dupHandshakeErr
		return
	}

	session.State = nextState
	return
}


