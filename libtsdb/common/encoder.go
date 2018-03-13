package common

import (
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
)

// TODO: might move it to top level

type Encoder interface {
	Len() int
	Bytes() []byte
	Reset()
	WritePointIntTagged(p *pb.PointIntTagged)
	WritePointDoubleTagged(p *pb.PointDoubleTagged)
	WriteSeriesIntTagged(p *pb.SeriesIntTagged)
	WriteSeriesDoubleTagged(p *pb.SeriesDoubleTagged)
}
