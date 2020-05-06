package bytebufferpool

import "io"

type ByteBuffer struct {
	buf []byte
}

func (b *ByteBuffer) Len() int {
	return len(b.buf)
}

func (b *ByteBuffer) ReadFrom(r io.Reader) (int64, error) {
	p := b.buf
	nStart := int64(len(p))
	nMax := int64(cap(p))
	n := nStart
	if nMax == 0 {
		nMax = 64
		p = make([]byte, nMax)
	} else {
		p = p[:nMax]
	}
	for {
		if n == nMax {
			nMax *= 2
			bNew := make([]byte, nMax)
			copy(bNew, p)
			p = bNew
		}
		nn, err := r.Read(p[n:])
		n += int64(nn)
		if err != nil {
			b.buf = p[:n]
			n -= nStart
			if err == io.EOF {
				return n, nil
			}
			return n, err
		}
	}
}

func (b *ByteBuffer) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b.buf)
	return int64(n), err
}

func (b *ByteBuffer) Bytes() []byte {
	return b.buf
}

func (b *ByteBuffer) Write(p []byte) (int, error) {
	b.buf = append(b.buf, p...)
	return len(p), nil
}

func (b *ByteBuffer) WriteByte(c byte) error {
	b.buf = append(b.buf, c)
	return nil
}

func (b *ByteBuffer) WriteString(s string) (int, error) {
	b.buf = append(b.buf, s...)
	return len(s), nil
}

func (b *ByteBuffer) Set(p []byte) {
	b.buf = append(b.buf[:0], p...)
}

func (b *ByteBuffer) SetString(s string) {
	b.buf = append(b.buf[:0], s...)
}

func (b *ByteBuffer) String() string {
	return string(b.buf)
}

func (b *ByteBuffer) Reset() {
	b.buf = b.buf[:0]
}
