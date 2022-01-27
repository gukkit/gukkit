package net

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"

	"gukkit/net/data/types"
)

type DataPacket struct {
	//ConnType connType //当前连接的状态，用于分析具体包内容

	buf *bytes.Buffer

	rawData  []byte //从网络流中截取到的数据
	rawIndex int    //数组指针索引

	TotalLen types.VarInt //整包长度
	DataLen  types.VarInt //包数据长度
	ID       types.VarInt //包的ID varint类型
	Data     []byte       //包数据

	Compressed bool //是否是压缩的
}

func (pk *DataPacket) Pack() (raw []byte, err error) {

	return
}

//从rawData中剪裁出数据包，若剪裁不出数据包则报 error
func (pk *DataPacket) Unpack(rawData []byte, compressed bool) (err error) {
	pk.rawData = rawData
	pk.rawIndex = 0

	pk.TotalLen, err = pk.readVarInt()
	if err != nil {
		return
	}

	if compressed {
		err = pk.unpackWithCompressed()
	} else {
		err = pk.unpackWithoutCompressed()
	}

	pk.rawData = nil
	pk.rawIndex = 0
	return
}

func (dataPk *DataPacket) readVarInt() (varInt types.VarInt, err error) {
	n, err := varInt.Decode(dataPk.buf)
	if err != nil {
		return 0, err
	}

	if n == 0 {
		return 0, nil //RawDataInComplete
	}

	dataPk.rawIndex += n
	return
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
		pk.DataLen = types.VarInt(len(pk.Data))
	}
	return
}

func (pk *DataPacket) ParseAngle(t *types.Angle) (err error) {
	_, err = t.Decode(pk.buf)
	return
}

func (pk *DataPacket) ParseBoolean(t *types.Boolean) (err error) {
	_, err = t.Decode(pk.buf)
	return
}

func (pk *DataPacket) ParseByteArray(t *types.ByteArray) (err error) {
	_, err = t.Decode(pk.buf)
	return
}

func (pk *DataPacket) ParseByte(t *types.Byte) (err error) {
	_, err = t.Decode(pk.buf)
	return
}

func (pk *DataPacket) ParseFloat(t *types.Float) (err error) {
	_, err = t.Decode(pk.buf)
	return
}

func (pk *DataPacket) ParseChat(t *types.Chat) (err error) {
	_, err = t.Decode(pk.buf)
	return
}

func (pk *DataPacket) ParseIdentifier(t *types.Identifier) (err error) {
	// _, err = t.
	return
}

func (pk *DataPacket) ParseVarInt(t *types.VarInt) (err error) {
	_, err = t.Decode(pk.buf)
	return
}

func (pk *DataPacket) ParseInt(t *types.Int) (err error) {
	_, err = t.Decode(pk.buf)
	return
}

func (pk *DataPacket) ParseLong(t *types.Long) (err error) {
	_, err = t.Decode(pk.buf)
	return
}

func (pk *DataPacket) ParsePosition(t *types.Position) (err error) {
	_, err = t.Decode(pk.buf)
	return
}

func (pk *DataPacket) ParseShort(t *types.Short) (err error) {
	_, err = t.Decode(pk.buf)
	return
}

func (pk *DataPacket) ParseString(t *types.String) (err error) {
	_, err = t.Decode(pk.buf)
	return
}

func (pk *DataPacket) ParseUnsignedByte(t *types.UnsignedByte) (err error) {
	_, err = t.Decode(pk.buf)
	return
}

func (pk *DataPacket) ParseUnsignedShort(t *types.UnsignedShort) (err error) {
	_, err = t.Decode(pk.buf)
	return
}

func (pk *DataPacket) ParseUnsignedLong(t *types.UnsignedLong) (err error) {
	_, err = t.Decode(pk.buf)
	return
}
