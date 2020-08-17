package packet

type PacketHandler func(inpk Serverbound) (outpk Clientbound, err error)
