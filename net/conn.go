package net

import (
	"bufio"
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
