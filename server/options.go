package server

type Option func (server *Server) error

// func SetPortOption(port uint16) Option {
// 	return func(server *Server) error {
// 		server.port = port
// 		return nil
// 	}
// }

// func Motd(motd string) Option {
// 	return func(server *Server) error {
// 		server.motd = motd
// 	}
// }

// func MaxPlayers(playerCount uint16) Option {
// 	return func(server *Server) error {
// 		server.maxPlayers = playerCount
// 		return nil
// 	}
// }