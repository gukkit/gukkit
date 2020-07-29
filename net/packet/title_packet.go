package packet

import (
	"gukkit/net/data/types"
)

type actionType types.VarInt

const (
	TitlePacketID = types.VarInt(0x50)

	TitleActionID        = actionType(0)
	SubtitleActionID     = actionType(1)
	SetActionBarID       = actionType(2)
	SetTimesAndDisplayID = actionType(3)
	HideActionID         = actionType(4)
	ResetActionID        = actionType(5)
)

type TitlePacket struct {
	Clientbound

	ActionType actionType
	Text       types.Chat

	FadeIn  types.Int
	Stay    types.Int
	FadeOut types.Int
}

func (pk *TitlePacket) Encode(w Writer) (err error) {
	if err = TitlePacketID.Encode(w); err != nil {
		return
	}

	if err = types.VarInt(pk.ActionType).Encode(w); err != nil {
		return
	}

	switch pk.ActionType {
	case TitleActionID, SubtitleActionID, SetActionBarID:
		err = pk.Text.Encode(w)
	case SetTimesAndDisplayID:
		err = pk.FadeIn.Encode(w)
		err = pk.Stay.Encode(w)
		err = pk.FadeOut.Encode(w)

	case HideActionID, ResetActionID:
		break
	}

	return
}
