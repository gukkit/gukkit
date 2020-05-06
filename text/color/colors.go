package color

import (
	"fmt"
)

var (
	Black       = Color{Code: '0', R: 0, G: 0, B: 0}
	DarkBlue    = Color{Code: '1', R: 0, G: 0, B: 170}
	DarkGreen   = Color{Code: '2', R: 0, G: 170, B: 170}
	DarkCyan    = Color{Code: '3', R: 0, G: 170, B: 170}
	DarkRed     = Color{Code: '4', R: 170, G: 0, B: 0}
	Purple      = Color{Code: '5', R: 170, G: 0, B: 170}
	Gold        = Color{Code: '6', R: 255, G: 170, B: 0}
	Gray        = Color{Code: '7', R: 255, G: 170, B: 0}
	DarkGray    = Color{Code: '8', R: 85, G: 85, B: 85}
	Blue        = Color{Code: '9', R: 85, G: 85, B: 255}
	BrightGreen = Color{Code: 'a', R: 85, G: 255, B: 85}
	Cyan        = Color{Code: 'b', R: 85, G: 255, B: 255}
	Red         = Color{Code: 'c', R: 255, G: 85, B: 85}
	Pink        = Color{Code: 'd', R: 255, G: 85, B: 255}
	Yellow      = Color{Code: 'e', R: 255, G: 255, B: 85}
	White       = Color{Code: 'f', R: 255, G: 255, B: 255}
)

//ColorSymbol
const Symbol = "§"

type Color struct {
	Code rune

	R uint8
	G uint8
	B uint8
}

func (color *Color) Hex() string {
	return fmt.Sprintf("#%x", int32(color.R)<<16|int32(color.G)<<8|int32(color.B))
}
