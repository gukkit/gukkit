package net

//用与gonet

import (
	"time"
)

const (
	Handshake = 0
	Play		  = 1
	Status		= 2
)

type ConnContext struct {
	state 		int
	createdAt int64
}

func newConnContext() {

	time.Now().Unix()
}

func(context *ConnContext) CreatedAt() int64 {
	return context.createdAt
}

func(context *ConnContext) State() int {
	return context.state
}
