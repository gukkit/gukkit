package packet

type Packet interface {
	ID() int

	// Decode() (*Packet, error)

	// Encode() ([]byte, error)
}

type ServerBound interface {

}

type PacketEncoder struct {

}

type PacketDecoder struct {

}
