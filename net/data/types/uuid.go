package types

import (
	"io"

	"github.com/google/uuid"
)

type UUID uuid.UUID

func (u UUID) Encode(w Writer) (err error) {
	_, err = w.Write(u[:])
	return
}

func (u *UUID) Decode(r Reader) (n int, err error) {
	n, err = io.ReadFull(r, (*u)[:])
	return
}
