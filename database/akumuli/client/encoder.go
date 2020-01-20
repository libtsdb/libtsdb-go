package client

import (
	"strconv"

	"github.com/libtsdb/libtsdb-go/tspb"
	"github.com/libtsdb/libtsdb-go/util/bytesutil"
)

type AkumuliEncoder struct {
	bytesutil.Buffer
}

func NewAkumuliEncoder() *AkumuliEncoder {
	return &AkumuliEncoder{}
}

func (e *AkumuliEncoder) WritePointIntTagged(p *tspb.PointIntTagged) {
	e.Buf = append(e.Buf, '+')
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, ' ')
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, tag.Key...)
		e.Buf = append(e.Buf, '=')
		e.Buf = append(e.Buf, tag.Value...)
		e.Buf = append(e.Buf, ' ')
	}
	e.Buf[len(e.Buf)-1] = '\r'
	e.Buf = append(e.Buf, "\n:"...)
	// time in integer
	e.Buf = strconv.AppendInt(e.Buf, p.Point.Time, 10)
	e.Buf = append(e.Buf, "\r\n+"...)
	// value in string
	e.Buf = strconv.AppendInt(e.Buf, p.Point.Value, 10)
	e.Buf = append(e.Buf, "\r\n"...)
}

func (e *AkumuliEncoder) WritePointDoubleTagged(p *tspb.PointDoubleTagged) {
	e.Buf = append(e.Buf, '+')
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, ' ')
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, tag.Key...)
		e.Buf = append(e.Buf, '=')
		e.Buf = append(e.Buf, tag.Value...)
		e.Buf = append(e.Buf, ' ')
	}
	e.Buf[len(e.Buf)-1] = '\r'
	e.Buf = append(e.Buf, "\n:"...)
	// time in integer
	e.Buf = strconv.AppendInt(e.Buf, p.Point.Time, 10)
	e.Buf = append(e.Buf, "\r\n+"...)
	// value in string
	e.Buf = strconv.AppendFloat(e.Buf, p.Point.Value, 'f', -1, 64)
	e.Buf = append(e.Buf, "\r\n"...)
}

func (e *AkumuliEncoder) WriteSeriesIntTagged(p *tspb.SeriesIntTagged) {
	var header []byte
	header = append(header, '+')
	header = append(header, p.Name...)
	header = append(header, ' ')
	for _, tag := range p.Tags {
		header = append(header, tag.Key...)
		header = append(header, '=')
		header = append(header, tag.Value...)
		header = append(header, ' ')
	}
	header[len(header)-1] = '\r'
	header = append(header, "\n:"...)
	for i := range p.Points {
		e.Buf = append(e.Buf, header...)
		// time in integer
		e.Buf = strconv.AppendInt(e.Buf, p.Points[i].Time, 10)
		e.Buf = append(e.Buf, "\r\n+"...)
		// value in string
		e.Buf = strconv.AppendInt(e.Buf, p.Points[i].Value, 10)
		e.Buf = append(e.Buf, "\r\n"...)
	}
}

func (e *AkumuliEncoder) WriteSeriesDoubleTagged(p *tspb.SeriesDoubleTagged) {
	var header []byte
	header = append(header, '+')
	header = append(header, p.Name...)
	header = append(header, ' ')
	for _, tag := range p.Tags {
		header = append(header, tag.Key...)
		header = append(header, '=')
		header = append(header, tag.Value...)
		header = append(header, ' ')
	}
	header[len(header)-1] = '\r'
	header = append(header, "\n:"...)
	for i := range p.Points {
		e.Buf = append(e.Buf, header...)
		// time in integer
		e.Buf = strconv.AppendInt(e.Buf, p.Points[i].Time, 10)
		e.Buf = append(e.Buf, "\r\n+"...)
		// value in string
		e.Buf = strconv.AppendFloat(e.Buf, p.Points[i].Value, 'f', -1, 64)
		e.Buf = append(e.Buf, "\r\n"...)
	}
}
