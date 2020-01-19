package protocol

import (
	"github.com/libtsdb/libtsdb-go/tspb"
)

type Encoder interface {
	Len() int
	Bytes() []byte
	Reset()
	WritePointIntTagged(p *tspb.PointIntTagged)
	WritePointDoubleTagged(p *tspb.PointDoubleTagged)
	WriteSeriesIntTagged(p *tspb.SeriesIntTagged)
	WriteSeriesDoubleTagged(p *tspb.SeriesDoubleTagged)
}
