package packet

import (
	"fmt"
//	"net"
	"io"
	"bufio"
	"errors"
)

func RecvPacket(pkBytes []byte) (pk Packet, err error) {
	pkReader, err := NewPacketReader(pkBytes)
	if err != nil {
		return nil, err
	}

	switch pkReader.ReadVarInt() {
	case 0x00:
		pk, err = NewServerListPingPacket(pkReader)
	default:
		return nil, errors.New("can find this packet")
	}
	return
}

func SendPacket(writer io.Writer, pk Packet) error {
	return nil
}

func readVarInt(reader *bufio.Reader) (result int, err error) {
	for i := 0; i < 5; i++ {
		b, err := reader.ReadByte()
		if err != nil {
			return 0, fmt.Errorf("read len of packet fail: %v", err)
		}
		result |= int(b&0x7F) << uint(7*i)
		if b & 0x80 == 0 {
			break
		}
	}
	return
}

func readVarLong(reader bufio.Reader) (result int, err error) {
	for i := 0; i < 10; i++ {
		b, err := reader.ReadByte()
		if err != nil {
			return 0, fmt.Errorf("read len of packet fail: %v", err)
		}
		result |= int(b&0x7F) << uint(7*i)
		if b & 0x80 == 0 {
			break
		}
	}
	return
}
