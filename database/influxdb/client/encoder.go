package client

import (
	"strconv"

	"github.com/libtsdb/libtsdb-go/tspb"
	"github.com/libtsdb/libtsdb-go/util/bytesutil"
)

const (
	defaultField = "v"
)

// InfluxDBEncoder encodes points with tags into InfluxDB's line protocol, it ONLY has a single field (default to v)
// DefaultField is used when encoding single field series, which is the case for most TSDB but not InfluxDB
// other tsdb (name, tags, value, ts)     : cpu.usage host=i7szx,dc=us-east 1.0 1359788400000
// influxdb (name, tags, field=value, ts) : cpu.usage host=i7szx,dc=us-east v=1.0 1359788400000
// ref https://github.com/influxdata/influxdb/blob/master/models/points.go#L2384 appendField
type InfluxDBEncoder struct {
	bytesutil.Buffer
	DefaultField string
}

func NewInfluxDBEncoder() *InfluxDBEncoder {
	return &InfluxDBEncoder{
		DefaultField: defaultField,
	}
}

// temperature,machine=unit42,type=assembly internal=32,external=100 1434055562000000035
func (e *InfluxDBEncoder) WritePointIntTagged(p *tspb.PointIntTagged) {
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, ',')
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, tag.Key...)
		e.Buf = append(e.Buf, '=')
		e.Buf = append(e.Buf, tag.Value...)
		e.Buf = append(e.Buf, ',')
	}
	e.Buf[len(e.Buf)-1] = ' '
	e.Buf = append(e.Buf, e.DefaultField...)
	e.Buf = append(e.Buf, '=')
	e.Buf = strconv.AppendInt(e.Buf, p.Point.Value, 10)
	e.Buf = append(e.Buf, ' ')
	e.Buf = strconv.AppendInt(e.Buf, p.Point.Time, 10)
	e.Buf = append(e.Buf, '\n')
}

func (e *InfluxDBEncoder) WritePointDoubleTagged(p *tspb.PointDoubleTagged) {
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, ',')
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, tag.Key...)
		e.Buf = append(e.Buf, '=')
		e.Buf = append(e.Buf, tag.Value...)
		e.Buf = append(e.Buf, ',')
	}
	e.Buf[len(e.Buf)-1] = ' '
	e.Buf = append(e.Buf, e.DefaultField...)
	e.Buf = append(e.Buf, '=')
	// TODO: most part are copy and pasted except this line ...
	e.Buf = strconv.AppendFloat(e.Buf, p.Point.Value, 'f', -1, 64)
	e.Buf = append(e.Buf, ' ')
	e.Buf = strconv.AppendInt(e.Buf, p.Point.Time, 10)
	e.Buf = append(e.Buf, '\n')
}

func (e *InfluxDBEncoder) WriteSeriesIntTagged(p *tspb.SeriesIntTagged) {
	// NOTE: InfluxDB does not support ingest multiple points in one line
	// first write the key, measurement + tags + default field
	var header []byte
	header = append(header, p.Name...)
	header = append(header, ',')
	for _, tag := range p.Tags {
		header = append(header, tag.Key...)
		header = append(header, '=')
		header = append(header, tag.Value...)
		header = append(header, ',')
	}
	header[len(header)-1] = ' '
	header = append(header, e.DefaultField...)
	header = append(header, '=')
	// then duplicate the header in every line
	for i := range p.Points {
		e.Buf = append(e.Buf, header...)
		e.Buf = strconv.AppendInt(e.Buf, p.Points[i].Value, 10)
		e.Buf = append(e.Buf, ' ')
		e.Buf = strconv.AppendInt(e.Buf, p.Points[i].Time, 10)
		e.Buf = append(e.Buf, '\n')
	}
}

func (e *InfluxDBEncoder) WriteSeriesDoubleTagged(p *tspb.SeriesDoubleTagged) {
	// NOTE: InfluxDB does not support ingest multiple points in one line
	// first write the key, measurement + tags + default field
	var header []byte
	header = append(header, p.Name...)
	header = append(header, ',')
	for _, tag := range p.Tags {
		header = append(header, tag.Key...)
		header = append(header, '=')
		header = append(header, tag.Value...)
		header = append(header, ',')
	}
	header[len(header)-1] = ' '
	header = append(header, e.DefaultField...)
	header = append(header, '=')
	// then duplicate the header in every line
	for i := range p.Points {
		e.Buf = append(e.Buf, header...)
		e.Buf = strconv.AppendFloat(e.Buf, p.Points[i].Value, 'f', -1, 64)
		e.Buf = append(e.Buf, ' ')
		e.Buf = strconv.AppendInt(e.Buf, p.Points[i].Time, 10)
		e.Buf = append(e.Buf, '\n')
	}
}
