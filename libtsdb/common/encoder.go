package common

import (
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
	"io"
)

type Encoder interface {
	Len() int
	Bytes() []byte
	ReadCloser() io.ReadCloser
	Reset()
	WritePointIntTagged(p *pb.PointIntTagged)
	WritePointDoubleTagged(p *pb.PointDoubleTagged)
}
