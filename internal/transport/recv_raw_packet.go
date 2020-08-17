package transport

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"sync"

	"gukkit/net/data/types"
	"gukkit/utils"
)

// var (
// 	RawDataInComplete = errors.New("raw data in complete")
// )

var (
	recvRawPacketPool = sync.Pool{
		New: func() interface{} {
			return &recvRawPacket{}
		},
	}
)

type recvRawPacket struct {
	rawData  []byte //从网络流中截取到的数据
	rawIndex int    //数组指针索引

	TotalLen types.VarInt    //整包长度
	DataLen  types.VarInt    //包数据长度
	ID       types.VarInt    //包的ID varint类型
	Data     types.ByteArray //包数据
	dataIdx  int

	Compressed bool //是否是压缩的
}

func (pk *recvRawPacket) Read(b []byte) (n int, err error) {

	return
}

func (pk *recvRawPacket) ReadByte() (b byte, err error) {

	return
}

//从rawData中剪裁出数据包，若剪裁不出数据包则报 error
func (pk *recvRawPacket) Unpack(rawData []byte, compressed bool) (err error) {
	pk.rawData = rawData
	pk.rawIndex = 0

	if pk.TotalLen, err = pk.readVarInt(); err != nil {
		return
	}

	if compressed {
		err = pk.unpackWithCompressed()
	} else {
		//非压缩拆包
		err = pk.unpackWithoutCompressed()
	}

	pk.rawData = nil
	pk.rawIndex = 0
	return
}

func (pk *recvRawPacket) readVarInt() (types.VarInt, error) {
	result, bit := utils.DecodeVarInt(pk.rawData[pk.rawIndex:])
	if bit == 0 {
		return 0, RawDataInComplete
	}

	pk.rawIndex += bit
	return types.VarInt(result), nil
}

func (pk *recvRawPacket) unCompressZlib(data []byte) (unpack []byte) {
	r, _ := zlib.NewReader(bytes.NewReader(data))

	unpack, _ = ioutil.ReadAll(r)
	return
}

func (pk *recvRawPacket) unpackWithCompressed() (err error) {
	pk.DataLen, err = pk.readVarInt()
	if err != nil {
		return
	}

	end := pk.rawIndex + int(pk.DataLen)
	zlibData := pk.rawData[pk.rawIndex:end]
	pk.rawData = pk.unCompressZlib(zlibData)
	pk.rawIndex = 0

	err = pk.unpackWithoutCompressed()
	return
}

func (pk *recvRawPacket) unpackWithoutCompressed() (err error) {
	pk.ID, err = pk.readVarInt()
	if err != nil {
		return
	}

	if len(pk.rawData)-1 > pk.rawIndex {
		pk.Data = pk.rawData[pk.rawIndex:]
		pk.DataLen = types.VarInt(len(pk.Data))
	}
	return
}
