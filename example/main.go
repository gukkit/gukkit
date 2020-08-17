package main

import (
	"fmt"
	"gukkit"
)

//this from go.mod, but you can set github/ihdj/gukkit

func main() {
	svr, _ := gukkit.NewServer(":25525")

	fmt.Println(svr.Listen(":25525"))
}
