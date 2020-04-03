package packet

type Packet interface {
	ID() int
}

type ServerListPingPacket struct {
	Version		int
	Address 	string
	Port			int
	NextState int
}

func NewServerListPingPacket(reader *PacketDataReader) (*ServerListPingPacket, error) {

	return nil, nil
}

func(pk *ServerListPingPacket) ID() int {
	return ServerListPingID
}

func ParsePacket(id int, data []byte) (pk *Packet, err error){
	switch id {
	case ServerListPingID:

	}

	return
}
