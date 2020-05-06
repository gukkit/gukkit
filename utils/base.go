package utils

func DecodeVarInt(data []byte) (result int, bit int) {
	for { //读数据前的长度标记
		value := data[bit]
		result |= int(value & 0b11111111) << int32(7*bit)

		bit++
		if value&0b10000000 == 0 {
			break
		} else if bit > 5 {
			return 0, 0
		}
	}

	return
}
