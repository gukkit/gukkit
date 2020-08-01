package login

import (
	"gukkit/net/data/types"
	"gukkit/net/packet"
)

const (
	DisconnectID         = types.VarInt(0)
	EncryptionRequestId  = types.VarInt(1)
	LoginSuccessID       = types.VarInt(2)
	SetCompressionID     = types.VarInt(3)
	LoginPluginRequestID = types.VarInt(4)
)

type DisconnectPacket struct {
	packet.Clientbound

	Reason types.Chat //掉线原因
}

func (pk *DisconnectPacket) Encode(w packet.Writer) (err error) {

	return
}

type EncryptionRequestPacket struct {
	packet.Clientbound

	ServerID    types.String
	PublicKey   types.ByteArray
	VerifyToken types.ByteArray
}

func (pk *EncryptionRequestPacket) Encode(w packet.Writer) (err error) {

	return
}

type LoginSuccessPacket struct {
	packet.Clientbound

	UUID     types.UUID
	Username types.String
}

type SetCompressionPacket struct {
	packet.Clientbound

	Threshold types.VarInt
}

type LoginPluginRequestPacket struct {
	packet.Clientbound

	MessageID types.VarInt
	Channel   types.Identifier
	Data      types.ByteArray
}
