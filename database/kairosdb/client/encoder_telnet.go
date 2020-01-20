package client

import (
	"strconv"

	"github.com/libtsdb/libtsdb-go/tspb"
	"github.com/libtsdb/libtsdb-go/util/bytesutil"
)

type KairosDBTelnetEncoder struct {
	bytesutil.Buffer
}

func NewKaiorsDBTelnetEncoder() *KairosDBTelnetEncoder {
	return &KairosDBTelnetEncoder{}
}

// putm <metric name> <time stamp> <value> <tag> <tag>... \n
func (e *KairosDBTelnetEncoder) WritePointIntTagged(p *tspb.PointIntTagged) {
	e.Buf = append(e.Buf, "putm "...)
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, ' ')
	e.Buf = strconv.AppendInt(e.Buf, p.Point.Time, 10)
	e.Buf = append(e.Buf, ' ')
	e.Buf = strconv.AppendInt(e.Buf, p.Point.Value, 10)
	e.Buf = append(e.Buf, ' ')
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, tag.Key...)
		e.Buf = append(e.Buf, '=')
		e.Buf = append(e.Buf, tag.Value...)
		e.Buf = append(e.Buf, ' ')
	}
	e.Buf[len(e.Buf)-1] = '\n'
}

func (e *KairosDBTelnetEncoder) WritePointDoubleTagged(p *tspb.PointDoubleTagged) {
	e.Buf = append(e.Buf, "putm "...)
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, ' ')
	e.Buf = strconv.AppendInt(e.Buf, p.Point.Time, 10)
	e.Buf = append(e.Buf, ' ')
	e.Buf = strconv.AppendFloat(e.Buf, p.Point.Value, 'f', -1, 64)
	e.Buf = append(e.Buf, ' ')
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, tag.Key...)
		e.Buf = append(e.Buf, '=')
		e.Buf = append(e.Buf, tag.Value...)
		e.Buf = append(e.Buf, ' ')
	}
	e.Buf[len(e.Buf)-1] = '\n'
}

func (e *KairosDBTelnetEncoder) WriteSeriesIntTagged(p *tspb.SeriesIntTagged) {
	var (
		header  []byte
		trailer []byte
	)
	header = append(header, "putm "...)
	header = append(header, p.Name...)
	header = append(header, ' ')
	trailer = append(trailer, ' ')
	for _, tag := range p.Tags {
		trailer = append(trailer, tag.Key...)
		trailer = append(trailer, '=')
		trailer = append(trailer, tag.Value...)
		trailer = append(trailer, ' ')
	}
	trailer[len(trailer)-1] = '\n'
	for _, p := range p.Points {
		e.Buf = append(e.Buf, header...)
		e.Buf = strconv.AppendInt(e.Buf, p.Time, 10)
		e.Buf = append(e.Buf, ' ')
		e.Buf = strconv.AppendInt(e.Buf, p.Value, 10)
		e.Buf = append(e.Buf, trailer...)
	}
}

func (e *KairosDBTelnetEncoder) WriteSeriesDoubleTagged(p *tspb.SeriesDoubleTagged) {
	var (
		header  []byte
		trailer []byte
	)
	header = append(header, "putm "...)
	header = append(header, p.Name...)
	header = append(header, ' ')
	trailer = append(trailer, ' ')
	for _, tag := range p.Tags {
		trailer = append(trailer, tag.Key...)
		trailer = append(trailer, '=')
		trailer = append(trailer, tag.Value...)
		trailer = append(trailer, ' ')
	}
	trailer[len(trailer)-1] = '\n'
	for _, p := range p.Points {
		e.Buf = append(e.Buf, header...)
		e.Buf = strconv.AppendInt(e.Buf, p.Time, 10)
		e.Buf = append(e.Buf, ' ')
		e.Buf = strconv.AppendFloat(e.Buf, p.Value, 'f', -1, 64)
		e.Buf = append(e.Buf, trailer...)
	}
}
