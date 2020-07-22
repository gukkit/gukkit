package packet

import (
	"gukkit/net/data/types"
)

type ParticlePacket struct {
	ID            types.Int
	LongDistance  types.Boolean
	X             types.Double
	Y             types.Double
	Z             types.Double
	OffsetX       types.Float
	OffsetY       types.Float
	OffsetZ       types.Float
	ParticleData  types.Float
	ParticleCount types.Int
	Data          []byte
}

func (pk *ParticlePacket) Encode(w Writer) (err error) {
	encoders := []types.TypeEncoder{
		pk.ID,
		pk.LongDistance,
		// pk.X,
		// pk.Y,
		// pk.Z,
		// pk.OffsetX,
		// pk.OffsetY,
		// pk.OffsetZ,
		// pk.ParticleData,
		pk.ParticleCount,
	}

	for _, encoder := range encoders {
		if err = encoder.Encode(w); err != nil {
			return
		}
	}

	return
}
