package event

type Event interface {
	Trigger() interface{}

	Cancel() error
}
