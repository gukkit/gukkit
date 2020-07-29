package main

import "fmt"

//this from go.mod, but you can set github/ihdj/gukkit

func main() {
	str := "ChinaHDJ"

	defer fmt.Println(str)

	str = "HDJ"

	//svr, _ := gukkit.NewServer(":25525")

	//fmt.Println(svr.Listen(":25525"))
}
