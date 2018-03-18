package kairosdb

import (
	"strconv"

	"github.com/libtsdb/libtsdb-go/libtsdb/common"
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
	"github.com/libtsdb/libtsdb-go/libtsdb/util/bytesutil"
)

var _ common.Encoder = (*TelnetEncoder)(nil)

type TelnetEncoder struct {
	bytesutil.Buffer
}

func NewTelnetEncoder() *TelnetEncoder {
	return &TelnetEncoder{}
}

// putm <metric name> <time stamp> <value> <tag> <tag>... \n
func (e *TelnetEncoder) WritePointIntTagged(p *pb.PointIntTagged) {
	e.Buf = append(e.Buf, "putm "...)
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, ' ')
	e.Buf = strconv.AppendInt(e.Buf, p.Point.T, 10)
	e.Buf = append(e.Buf, ' ')
	e.Buf = strconv.AppendInt(e.Buf, p.Point.V, 10)
	e.Buf = append(e.Buf, ' ')
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, tag.K...)
		e.Buf = append(e.Buf, '=')
		e.Buf = append(e.Buf, tag.V...)
		e.Buf = append(e.Buf, ' ')
	}
	e.Buf[len(e.Buf)-1] = '\n'
}

func (e *TelnetEncoder) WritePointDoubleTagged(p *pb.PointDoubleTagged) {
	e.Buf = append(e.Buf, "putm "...)
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, ' ')
	e.Buf = strconv.AppendInt(e.Buf, p.Point.T, 10)
	e.Buf = append(e.Buf, ' ')
	e.Buf = strconv.AppendFloat(e.Buf, p.Point.V, 'f', -1, 64)
	e.Buf = append(e.Buf, ' ')
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, tag.K...)
		e.Buf = append(e.Buf, '=')
		e.Buf = append(e.Buf, tag.V...)
		e.Buf = append(e.Buf, ' ')
	}
	e.Buf[len(e.Buf)-1] = '\n'
}

func (e *TelnetEncoder) WriteSeriesIntTagged(p *pb.SeriesIntTagged) {
	var (
		header  []byte
		trailer []byte
	)
	header = append(header, "putm "...)
	header = append(header, p.Name...)
	header = append(header, ' ')
	trailer = append(trailer, ' ')
	for _, tag := range p.Tags {
		trailer = append(trailer, tag.K...)
		trailer = append(trailer, '=')
		trailer = append(trailer, tag.V...)
		trailer = append(trailer, ' ')
	}
	trailer[len(trailer)-1] = '\n'
	for i := range p.Points {
		e.Buf = append(e.Buf, header...)
		e.Buf = strconv.AppendInt(e.Buf, p.Points[i].T, 10)
		e.Buf = append(e.Buf, ' ')
		e.Buf = strconv.AppendInt(e.Buf, p.Points[i].V, 10)
		e.Buf = append(e.Buf, trailer...)
	}
}

func (e *TelnetEncoder) WriteSeriesDoubleTagged(p *pb.SeriesDoubleTagged) {
	var (
		header  []byte
		trailer []byte
	)
	header = append(header, "putm "...)
	header = append(header, p.Name...)
	header = append(header, ' ')
	trailer = append(trailer, ' ')
	for _, tag := range p.Tags {
		trailer = append(trailer, tag.K...)
		trailer = append(trailer, '=')
		trailer = append(trailer, tag.V...)
		trailer = append(trailer, ' ')
	}
	trailer[len(trailer)-1] = '\n'
	for i := range p.Points {
		e.Buf = append(e.Buf, header...)
		e.Buf = strconv.AppendInt(e.Buf, p.Points[i].T, 10)
		e.Buf = append(e.Buf, ' ')
		e.Buf = strconv.AppendFloat(e.Buf, p.Points[i].V, 'f', -1, 64)
		e.Buf = append(e.Buf, trailer...)
	}
}
