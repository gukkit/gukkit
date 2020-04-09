package gukkit

import (
	"os"
	"io/ioutil"
)

type motdRender struct {
	maxCap	*int32
	online	*int32
	players *[]*Player

	favicon []byte
}

func(motd *motdRender) SetFaviconByBytes(bytes []byte) {

}

func(motd *motdRender) SetFaviconByFile(file *os.File) error {
	prefix := "data:image/png;base64,"

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	motd.favicon = append([]byte(prefix), data...)
	return nil
}
