package handshaking

import (
	"gukkit/internal/packet"
	"gukkit/net/data/types"
)

type HandshakePacket struct {
	packet.Serverbound

	ID        types.VarInt
	Version   types.VarInt
	Address   types.String
	Port      types.UnsignedShort
	NextState types.VarInt //value 1 for status, 2 for login
}

func (pk *HandshakePacket) Decode(r packet.Reader) (err error) {
	if _, err = pk.ID.Decode(r); err != nil {
		return
	}

	if _, err = pk.Version.Decode(r); err != nil {
		return
	}

	if _, err = pk.Address.Decode(r); err != nil {
		return
	}

	if _, err = pk.Port.Decode(r); err != nil {
		return
	}

	if _, err = pk.NextState.Decode(r); err != nil {
		return
	}

	return
}
