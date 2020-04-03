package net

import (
	"fmt"

	gonet "net"
)

type Listener struct {
	ln gonet.Listener
}

func Listen(address string) (*Listener, error) {
	ln, err := gonet.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Listener{ln: ln}, nil
}

func(listener *Listener) Accepter() {
	buf := make([]byte, 2048)
	for {
		conn, err := listener.ln.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		//reader := bufio.NewReader(conn)

		n, err := conn.Read(buf)

		fmt.Println(conn, buf[:n], n)
	}
}
