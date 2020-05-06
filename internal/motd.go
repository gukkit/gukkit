package internal

import (
	"gukkit"
	"os"
	"io/ioutil"
)

var defaultMotdTitle = "Gukkit Server for Minecraft By iHDJ"

type MotdRender struct {
	maxCap	*int32
	online	*int32
	players *[]*gukkit.Player

	title		string
	favicon []byte
}

type MotdJsonResponse struct {
	Favicon	string
}

func(motd *MotdRender) SetFaviconByFile(file *os.File) error {
	prefix := "data:image/png;base64,"

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	motd.favicon = append([]byte(prefix), data...)
	return nil
}

func(motd *MotdRender) Render() (out []byte) {

}

