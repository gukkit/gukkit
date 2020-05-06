package packet

const (
	ServerListPingID = 0
)

const (
	StatusPongPacketID = 1

)


var (
	StatusPacketSet = []Packet {
		ServerListPingID: ServerListPingPacket{},

	}
)

