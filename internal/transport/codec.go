package transport

import (
	"errors"
	"gukkit/utils"

	"github.com/panjf2000/gnet"
)

var (
	RawDataInComplete = errors.New("raw data in complete")
	EmptyContent      = errors.New("raw data is empty")
)

//防止TCP粘包
type Codec struct {
}

func (codec *Codec) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
	return nil, nil
}

//这Codec层面上将数据包的数据解开，
func (codec *Codec) Decode(c gnet.Conn) ([]byte, error) {
	buf := c.Read()
	if len(buf) == 0 {
		return nil, EmptyContent
	}

	totalLen, n := utils.DecodeVarInt(buf)
	if n == 0 || (len(buf)-n) < int(totalLen) {
		return nil, RawDataInComplete
	}

	c.ShiftN(totalLen + n)
	return buf[:totalLen+n], nil
}
