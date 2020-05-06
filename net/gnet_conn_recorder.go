package net

type connType	int32

const (
	Conn_Type_Handshake connType = 1
	Conn_Type_StatusQuo connType = 2
	Conn_Type_Playing		connType = 3
)

type connRecorder struct {
	ConnType		connType
	Compressed	bool
}

func newDefaultConnRecorder() *connRecorder {
	return &connRecorder{ConnType: Conn_Type_Handshake, Compressed: false}
}