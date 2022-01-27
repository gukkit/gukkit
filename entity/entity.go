package entity

type Entity interface {
	X() int
	Y() int
	Z() int
	// Level() *level.Level

	LastX() int
	LastY() int
	LastZ() int
}
