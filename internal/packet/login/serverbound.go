package login

import (
	"gukkit/internal/packet"
	"gukkit/net"
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

func (pk *LoginStartPacket) Decode(p *net.DataPacket) (err error) {
	if err = p.ParseVarInt(&pk.ID); err != nil {
		return
	}

	err = p.ParseString(&pk.Name)
	return
}

type EncryptionResponsePacket struct {
	packet.Serverbound

	ID           types.VarInt
	SharedSecret types.ByteArray
	VerifyToken  types.ByteArray
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

type LoginPluginResponsePacket struct {
	packet.Serverbound

	ID         types.VarInt
	MessageID  types.VarInt
	Successful types.Boolean
	Data       types.ByteArray //Optional
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
