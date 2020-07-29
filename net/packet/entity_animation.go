package packet

import "gukkit/net/data/types"

const (
	EntityAnimationPacketID = types.VarInt(0x06)

	SwingMainArmAnimation        = types.UnsignedByte(0) //摆动双臂动画
	TakeDamageAnimation          = types.UnsignedByte(1) //受伤动画
	LeaveBedAnimation            = types.UnsignedByte(2) //起床动画
	SwingOffhandAnimation        = types.UnsignedByte(3) //副手摆动动画
	CriticalEffectAnimation      = types.UnsignedByte(4)
	MagicCriticalEffectAnimation = types.UnsignedByte(5)
)

type EntityAnimationPacket struct {
	Clientbound

	EntityID  types.VarInt
	Animation types.UnsignedByte
}
