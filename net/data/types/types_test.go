package types

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPositionType(t *testing.T) {
	var buf bytes.Buffer
	pos := &Position{
		X: 33554431,
		Y: 2047,
		Z: -33554431,
	}

	var p2 *Position = &Position{}

	pos.Encode(&buf)
	p2.Decode(&buf)

	fmt.Println(p2)
}

func TestTypes(t *testing.T) {
	//s1 = String("ChinaHDJ")
	var str String = String("ChinaHDJ")
	var ss String
	var buf bytes.Buffer

	fmt.Println(str.Encode(&buf))

	ss.Decode(&buf)
	fmt.Println(ss)
}
