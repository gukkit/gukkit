package ticker

import (
	"time"
)

const DefaultInterval = time.Duration(50)

type Ticker struct {
	Interval time.Duration
	tick     uint64
}

func (ticker *Ticker) doTick() {
	for ; ; time.Sleep(ticker.Interval * time.Millisecond) {

		ticker.tick++
	}
}
