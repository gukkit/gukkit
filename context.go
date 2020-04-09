package gukkit

import (
	"time"
)

const (
	HandshakeStatus 	= 0		//初次连接	等待客户端发送ServerListPacket包确定下一状态
	ServerListStatus	= 1		//服务器列表状态
	PlayingStatus			= 2		//游戏内状态，即真正游戏内容
)

type context struct {
	state 		int
	createdAt int64
}

func(ctx *context) Deadline() (deadline time.Time, ok bool) {
	return
}

func(ctx *context) Done() <-chan struct{} {
	return nil
}

func(ctx *context) Err() error {
	return nil
}

func(ctx *context) Value(key interface{}) interface{} {
	switch key {
	case "state":
	case "createdAt":

	}

	return nil
}
