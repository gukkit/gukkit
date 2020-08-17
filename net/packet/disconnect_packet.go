package packet

import "gukkit/net/data/types"

const (
	DisconnectPacketID = types.VarInt(0x1B)
)

type DisconnectPacket struct {
	Clientbound

	Reason types.Chat
}

func (pk DisconnectPacket) Encode(w Writer) (err error) {
	if err = DisconnectPacketID.Encode(w); err != nil {
		return
	}

	err = pk.Reason.Encode(w)
	return
}
