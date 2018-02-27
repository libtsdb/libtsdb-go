package graphite

import (
	"strconv"

	"github.com/libtsdb/libtsdb-go/libtsdb/common"
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
	"github.com/libtsdb/libtsdb-go/libtsdb/util/bytesutil"
)

// TODO: pickle format https://github.com/lomik/graphite-pickle

var _ common.Encoder = (*TextEncoder)(nil)

// TextEncoder encodes points in graphite text format and use tag
// TODO: text encoder that does not use tag
type TextEncoder struct {
	bytesutil.Buffer
}

// PickleEncoder encodes points in graphite pickle format
type PickleEncoder struct {
}

func NewTextEncoder() *TextEncoder {
	return &TextEncoder{}
}

// WritePointIntTagged keeps tag
// NOTE: graphite only supports float, you will get float when read
// NOTE: tag is only supported since 1.1.x
// `my.series;tag1=value1;tag2=value2 10 1519266078`
func (e *TextEncoder) WritePointIntTagged(p *pb.PointIntTagged) {
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, ';')
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, tag.K...)
		e.Buf = append(e.Buf, '=')
		e.Buf = append(e.Buf, tag.V...)
		e.Buf = append(e.Buf, ';')
	}
	e.Buf[len(e.Buf)-1] = ' '
	e.Buf = strconv.AppendInt(e.Buf, p.Point.V, 10)
	e.Buf = append(e.Buf, ' ')
	// TODO: time precision is second, not ms or ns like other tsdb
	e.Buf = strconv.AppendInt(e.Buf, p.Point.T, 10)
	e.Buf = append(e.Buf, '\n')
}

// NOTE: tag is only supported since 1.1.x
// `my.series;tag1=value1;tag2=value2 10.2 1519266078`
func (e *TextEncoder) WritePointDoubleTagged(p *pb.PointDoubleTagged) {
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, ';')
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, tag.K...)
		e.Buf = append(e.Buf, '=')
		e.Buf = append(e.Buf, tag.V...)
		e.Buf = append(e.Buf, ';')
	}
	e.Buf[len(e.Buf)-1] = ' '
	// TODO: most part are copy and pasted except this line ...
	e.Buf = strconv.AppendFloat(e.Buf, p.Point.V, 'f', -1, 64)
	e.Buf = append(e.Buf, ' ')
	// TODO: time precision is second, not ms or ns like other tsdb
	e.Buf = strconv.AppendInt(e.Buf, p.Point.T, 10)
	e.Buf = append(e.Buf, '\n')
}
