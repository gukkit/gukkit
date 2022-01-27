package packet

import "gukkit/net/data/types"

type SetDifficultyPacket struct {
	ID            types.VarInt
	NewDifficulty types.VarInt //0: peaceful, 1: easy, 2: normal, 3: hard .
}
