package login

import (
	"gukkit/internal/packet"
	"gukkit/net/data/types"
)

const (
	LoginStartID          = types.VarInt(0)
	EncryptionResponseID  = types.VarInt(1)
	LoginPluginResponseID = types.VarInt(2)
)

type LoginStartPacket struct {
	packet.Serverbound

	ID   types.VarInt
	Name types.String
}

type EncryptionResponsePacket struct {
	packet.Serverbound

	ID           types.VarInt
	SharedSecret types.ByteArray
	VerifyToken  types.ByteArray
}

type LoginPluginResponsePacket struct {
	packet.Serverbound

	ID         types.VarInt
	MessageID  types.VarInt
	Successful types.Boolean
	Data       types.ByteArray //Optional
}

func (pk *LoginStartPacket) Decode(r packet.Reader) (err error) {
	if _, err = pk.ID.Decode(r); err != nil {
		return
	}

	if _, err = pk.Name.Decode(r); err != nil {
		return
	}

	return
}

func (pk *EncryptionResponsePacket) Decode(r packet.Reader) (err error) {
	if _, err = pk.ID.Decode(r); err != nil {
		return
	}

	if _, err = pk.SharedSecret.Decode(r); err != nil {
		return
	}

	if _, err = pk.VerifyToken.Decode(r); err != nil {
		return
	}

	return
}

func (pk *LoginPluginResponsePacket) Decode(r packet.Reader) (err error) {
	if _, err = pk.ID.Decode(r); err != nil {
		return
	}

	if _, err = pk.MessageID.Decode(r); err != nil {
		return
	}

	if _, err = pk.Successful.Decode(r); err != nil {
		return
	}

	if _, err = pk.Data.Decode(r); err != nil {
		return
	}

	return
}
