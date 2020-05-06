package intfac

type Server interface {
	Listen(addr string) error

	SetMaxPlayer(max int)

	Motd() Motd
}