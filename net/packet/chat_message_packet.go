package packet

import "gukkit/net/data/types"

const (
	ChatMessagePacketID = types.VarInt(0x0F)

	ChatBox       = types.Byte(0)
	SystemMessage = types.Byte(1)
	GameInfo      = types.Byte(2) //above hotbar
)

type ChatMessagePacket struct {
	Clientbound

	Chat     types.Chat
	Position types.Byte //ChatBox | SystemMessage | GameInfo
}

func (pk *ChatMessagePacket) Encode(w Writer) (err error) {
	if err = ChatMessagePacketID.Encode(w); err != nil {
		return
	}

	if err = pk.Chat.Encode(w); err != nil {
		return
	}

	if err = pk.Position.Encode(w); err != nil {
		return
	}

	return
}
