package packet

import (
	"errors"
)

type ServerListPingPacket struct {
	Version		int32
	Address 	string
	Port			uint16
	NextState int32
}

func EncodeServerListPingPacket(reader *PacketReader) (pk *ServerListPingPacket, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = errors.New(rec.(string))
			return
		}
	}()

	pk = new(ServerListPingPacket)
	pk.Version 	 = reader.ReadVarInt()
	pk.Address 	 = reader.ReadString()
	pk.Port 	 	 = reader.ReadUnsignedShort()
	pk.NextState = reader.ReadVarInt()

	return pk, nil
}

func(pk *ServerListPingPacket) ID() int {
	return ServerListPingID
}
