package packet

type Packet interface {
	// Decode() (*Packet, error)
	// Encode() ([]byte, error)
}

type ServerBound interface {
}

type PacketEncoder struct {
}

type PacketDecoder struct {
}
