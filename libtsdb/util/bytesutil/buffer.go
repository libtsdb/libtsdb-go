package bytesutil

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Buffer is a dump struct that wraps around a byte slice, it does not handle grow like bytes.Buffer
// but it can be used directly with Append* style function which requires a byte slice directly
type Buffer struct {
	Buf []byte
}

func (b *Buffer) Reset() {
	b.Buf = b.Buf[:0]
}

func (b *Buffer) Len() int {
	return len(b.Buf)
}

func (b *Buffer) Cap() int {
	return cap(b.Buf)
}

func (b *Buffer) Bytes() []byte {
	return b.Buf
}

func (b *Buffer) ReadCloser() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(b.Buf))
}
