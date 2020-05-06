package gukkit

import (
	"os"
)

//我都懒得注释了，懂的自然懂
type WhiteList interface {
	Add(name string) bool

	Remove(name string) bool

	Exist(name string) bool
}

type fileWhiteList struct {
	file 		*os.File
	existed []string
}

func(fwl *fileWhiteList) Add(name string) bool {

	return false
}
