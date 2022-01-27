package net

import (
	"gukkit/net"
	"sync"
)

var dataPool = dataPacketPool{
	pool: sync.Pool{
		New: func() interface{} {
			return &net.DataPacket{}
		},
	},
}

type dataPacketPool struct {
	pool sync.Pool
}

func (p *dataPacketPool) Unpack() {

}

func (p *dataPacketPool) Get() *net.DataPacket {
	return p.pool.Get().(*net.DataPacket)
}

func (p *dataPacketPool) Put(pk *net.DataPacket) {

	p.pool.Put(pk)
}
