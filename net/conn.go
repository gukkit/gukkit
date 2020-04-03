package net

import (
	"bufio"
//	"errors"
//	"gukkit/net/packet"
	gonet "net"
)

type Conn struct {
	reader *bufio.Reader
	conn gonet.Conn
}

func newConn(conn gonet.Conn) *Conn {
	return &Conn {
		reader: bufio.NewReader(conn),
		conn: conn,
	}
}
