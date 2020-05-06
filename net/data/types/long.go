package types

type Long int64

func(long Long) Encode() (b []byte, err error) {
	b = []byte {
		byte(long >> 56),
		byte(long >> 48),
		byte(long >> 40),
		byte(long >> 32),
		byte(long >> 24),
		byte(long >> 16),
		byte(long >> 8),
		byte(long),
	}

	return
}
