package main

import (
	"fmt"
	"gukkit"
)

type A struct {
	Name string
}

func main() {
	var a A

	a.Name = "ChinaHDJ"

	svr, _ := gukkit.NewServer(":25525")

	fmt.Println(svr.Listen(":25525"))
}
