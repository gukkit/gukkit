package packet

import (
	"bytes"
	"compress/zlib"
	"errors"
	"io/ioutil"

	"gukkit/net/data/types"
	"gukkit/utils"
)

var (
	RawDataInComplete = errors.New("raw data in complete")
)

type DataPacket struct {
	rawData  []byte //从网络流中截取到的数据
	rawIndex int    //数组指针索引

	TotalLen types.Int       //整包长度
	DataLen  types.Int       //包数据长度
	ID       types.Int       //包的ID varint类型
	Data     types.ByteArray //包数据

	Compressed bool //是否是压缩的
}

func (pk *DataPacket) Pack() (raw []byte, err error) {
	if pk.Compressed {

	} else {

	}

	return
}

//从rawData中剪裁出数据包，若剪裁不出数据包则报 error
func (pk *DataPacket) Unpack(rawData []byte, compressed bool) (err error) {
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

func (dataPk *DataPacket) readVarInt() (types.Int, error) {
	result, bit := utils.DecodeVarInt(dataPk.rawData[dataPk.rawIndex:])
	if bit == 0 {
		return 0, RawDataInComplete
	}

	dataPk.rawIndex += bit
	return types.Int(result), nil
}

func (pk *DataPacket) unCompressZlib(data []byte) (unpack []byte) {
	r, _ := zlib.NewReader(bytes.NewReader(data))

	unpack, _ = ioutil.ReadAll(r)
	return
}

func (pk *DataPacket) unpackWithCompressed() (err error) {
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

func (pk *DataPacket) unpackWithoutCompressed() (err error) {
	pk.ID, err = pk.readVarInt()
	if err != nil {
		return
	}

	if len(pk.rawData)-1 > pk.rawIndex {
		pk.Data = pk.rawData[pk.rawIndex:]
		pk.DataLen = types.Int(len(pk.Data))
	}
	return
}
