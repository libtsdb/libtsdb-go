package akumuli

import (
	"strconv"

	"github.com/libtsdb/libtsdb-go/libtsdb/common"
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
	"github.com/libtsdb/libtsdb-go/libtsdb/util/bytesutil"
)

var _ common.Encoder = (*Encoder)(nil)

type Encoder struct {
	bytesutil.Buffer
}

func NewEncoder() *Encoder {
	return &Encoder{}
}

func (e *Encoder) WritePointIntTagged(p *pb.PointIntTagged) {
	e.Buf = append(e.Buf, '+')
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, ' ')
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, tag.K...)
		e.Buf = append(e.Buf, '=')
		e.Buf = append(e.Buf, tag.V...)
		e.Buf = append(e.Buf, ' ')
	}
	e.Buf[len(e.Buf)-1] = '\r'
	e.Buf = append(e.Buf, "\n:"...)
	// time in integer
	e.Buf = strconv.AppendInt(e.Buf, p.Point.T, 10)
	e.Buf = append(e.Buf, "\r\n+"...)
	// value in string
	e.Buf = strconv.AppendInt(e.Buf, p.Point.V, 10)
	e.Buf = append(e.Buf, "\r\n"...)
}

func (e *Encoder) WritePointDoubleTagged(p *pb.PointDoubleTagged) {
	e.Buf = append(e.Buf, '+')
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, ' ')
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, tag.K...)
		e.Buf = append(e.Buf, '=')
		e.Buf = append(e.Buf, tag.V...)
		e.Buf = append(e.Buf, ' ')
	}
	e.Buf[len(e.Buf)-1] = '\r'
	e.Buf = append(e.Buf, "\n:"...)
	// time in integer
	e.Buf = strconv.AppendInt(e.Buf, p.Point.T, 10)
	e.Buf = append(e.Buf, "\r\n+"...)
	// value in string
	e.Buf = strconv.AppendFloat(e.Buf, p.Point.V, 'f', -1, 64)
	e.Buf = append(e.Buf, "\r\n"...)
}

func (e *Encoder) WriteSeriesIntTagged(p *pb.SeriesIntTagged) {
	var header []byte
	header = append(header, '+')
	header = append(header, p.Name...)
	header = append(header, ' ')
	for _, tag := range p.Tags {
		header = append(header, tag.K...)
		header = append(header, '=')
		header = append(header, tag.V...)
		header = append(header, ' ')
	}
	header[len(header)-1] = '\r'
	header = append(header, "\n:"...)
	for i := range p.Points {
		e.Buf = append(e.Buf, header...)
		// time in integer
		e.Buf = strconv.AppendInt(e.Buf, p.Points[i].T, 10)
		e.Buf = append(e.Buf, "\r\n+"...)
		// value in string
		e.Buf = strconv.AppendInt(e.Buf, p.Points[i].V, 10)
		e.Buf = append(e.Buf, "\r\n"...)
	}
}

func (e *Encoder) WriteSeriesDoubleTagged(p *pb.SeriesDoubleTagged) {
	var header []byte
	header = append(header, '+')
	header = append(header, p.Name...)
	header = append(header, ' ')
	for _, tag := range p.Tags {
		header = append(header, tag.K...)
		header = append(header, '=')
		header = append(header, tag.V...)
		header = append(header, ' ')
	}
	header[len(header)-1] = '\r'
	header = append(header, "\n:"...)
	for i := range p.Points {
		e.Buf = append(e.Buf, header...)
		// time in integer
		e.Buf = strconv.AppendInt(e.Buf, p.Points[i].T, 10)
		e.Buf = append(e.Buf, "\r\n+"...)
		// value in string
		e.Buf = strconv.AppendFloat(e.Buf, p.Points[i].V, 'f', -1, 64)
		e.Buf = append(e.Buf, "\r\n"...)
	}
}
