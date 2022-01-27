package event

type Event interface {
	Trigger() interface{}
	Done() <-chan struct{}
	Cancel() error
}
