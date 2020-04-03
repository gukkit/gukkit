package main

import (
	"fmt"
	"gukkit/net"
)

func main() {
	listener, err := net.Listen(":25525")
	if err != nil {
		fmt.Println(err)
		return
	}

	listener.Accepter()
}
