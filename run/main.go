package main

import (
	"fmt"
	//"fmt"
	//"gukkit/net"
	"gukkit/server"
)
//就是个main方法，方便你们以后编译
func main() {
	// listener, err := net.Listen(":25525")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// listener.Accepter()

	svr, _ := server.NewServer(":25525")

	fmt.Println(svr.Run(":25525"))

}
