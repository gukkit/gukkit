package gukkit

const (
	HandshakeStatus 	= 0		//初次连接	等待客户端发送ServerListPacket包确定下一状态
	ServerListStatus	= 1		//服务器列表状态
	PlayingStatus			= 2		//游戏内状态，即真正游戏内容
)

type context struct {
	state 		int
	createdAt int64
	pkCompressed bool //是否接收压缩的数据包
}
