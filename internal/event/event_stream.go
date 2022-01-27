package event

import "sync"

var (
	regMutex      sync.Mutex
	RegistedEvent map[string][]Event
)

func Register() {
	regMutex.Lock()

	regMutex.Unlock()
}

type Event struct {
}

type EventStream struct {
}
