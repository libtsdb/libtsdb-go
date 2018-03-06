package common

import (
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
)

type Encoder interface {
	Len() int
	Bytes() []byte
	Reset()
	WritePointIntTagged(p *pb.PointIntTagged)
	WritePointDoubleTagged(p *pb.PointDoubleTagged)
}
