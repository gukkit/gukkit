package main

import (
	"fmt"
	//"fmt"
	//"gukkit/net"
	"gukkit/server"
)


func main() {
	svr, _ := server.NewServer(":25525")

	fmt.Println(svr.Listen(":25525"))
}
