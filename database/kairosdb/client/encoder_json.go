package client

import (
	"strconv"

	"github.com/libtsdb/libtsdb-go/tspb"
	"github.com/libtsdb/libtsdb-go/util/bytesutil"
)

// KairosDBJSONEncoder support mix of single point and series
type KairosDBJSONEncoder struct {
	bytesutil.Buffer
	finalized bool
}

func NewKairosDBJSONEncoder() *KairosDBJSONEncoder {
	e := &KairosDBJSONEncoder{}
	e.Reset()
	return e
}

func (e *KairosDBJSONEncoder) Reset() {
	e.Buffer.Reset()
	// start of json array
	e.Buf = append(e.Buf, '[')
	// FIXED: found via https://github.com/xephonhq/xephon-b/issues/36
	e.finalized = false
}

func (e *KairosDBJSONEncoder) Bytes() []byte {
	if e.finalized {
		return e.Buffer.Bytes()
	}
	// replace last extra comma with end of json array
	e.Buf[len(e.Buf)-1] = ']'
	e.finalized = true
	return e.Buffer.Bytes()
}

func (e *KairosDBJSONEncoder) WritePointIntTagged(p *tspb.PointIntTagged) {
	e.Buf = append(e.Buf, `{"name":"`...)
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, `",`...)
	e.Buf = append(e.Buf, `"timestamp":`...)
	e.Buf = strconv.AppendInt(e.Buf, p.Point.Time, 10)
	e.Buf = append(e.Buf, `,"value":`...)
	e.Buf = strconv.AppendInt(e.Buf, p.Point.Value, 10)
	e.Buf = append(e.Buf, `,"tags":{`...)
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, '"')
		e.Buf = append(e.Buf, tag.Key...)
		e.Buf = append(e.Buf, `":"`...)
		e.Buf = append(e.Buf, tag.Value...)
		e.Buf = append(e.Buf, `",`...)
	}
	e.Buf[len(e.Buf)-1] = '}'
	e.Buf = append(e.Buf, `},`...)
}

func (e *KairosDBJSONEncoder) WritePointDoubleTagged(p *tspb.PointDoubleTagged) {
	e.Buf = append(e.Buf, `{"name":"`...)
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, `",`...)
	e.Buf = append(e.Buf, `"timestamp":`...)
	e.Buf = strconv.AppendInt(e.Buf, p.Point.Time, 10)
	e.Buf = append(e.Buf, `,"value":`...)
	// NOTE: the only difference with write int
	e.Buf = strconv.AppendFloat(e.Buf, p.Point.Value, 'f', -1, 64)
	e.Buf = append(e.Buf, `,"tags":{`...)
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, '"')
		e.Buf = append(e.Buf, tag.Key...)
		e.Buf = append(e.Buf, `":"`...)
		e.Buf = append(e.Buf, tag.Value...)
		e.Buf = append(e.Buf, `",`...)
	}
	e.Buf[len(e.Buf)-1] = '}'
	e.Buf = append(e.Buf, `},`...)
}

func (e *KairosDBJSONEncoder) WriteSeriesIntTagged(p *tspb.SeriesIntTagged) {
	e.Buf = append(e.Buf, `{"name":"`...)
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, `","datapoints":[`...)
	for _, p := range p.Points {
		e.Buf = append(e.Buf, '[')
		e.Buf = strconv.AppendInt(e.Buf, p.Time, 10)
		e.Buf = append(e.Buf, ',')
		e.Buf = strconv.AppendInt(e.Buf, p.Value, 10)
		e.Buf = append(e.Buf, `],`...)
	}
	e.Buf[len(e.Buf)-1] = ']'
	e.Buf = append(e.Buf, `,"tags":{`...)
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, '"')
		e.Buf = append(e.Buf, tag.Key...)
		e.Buf = append(e.Buf, `":"`...)
		e.Buf = append(e.Buf, tag.Value...)
		e.Buf = append(e.Buf, `",`...)
	}
	e.Buf[len(e.Buf)-1] = '}'
	e.Buf = append(e.Buf, `},`...)
}

func (e *KairosDBJSONEncoder) WriteSeriesDoubleTagged(p *tspb.SeriesDoubleTagged) {
	e.Buf = append(e.Buf, `{"name":"`...)
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, `","datapoints":[`...)
	for _, p := range p.Points {
		e.Buf = append(e.Buf, '[')
		e.Buf = strconv.AppendInt(e.Buf, p.Time, 10)
		e.Buf = append(e.Buf, ',')
		e.Buf = strconv.AppendFloat(e.Buf, p.Value, 'f', -1, 64)
		e.Buf = append(e.Buf, `],`...)
	}
	e.Buf[len(e.Buf)-1] = ']'
	e.Buf = append(e.Buf, `,"tags":{`...)
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, '"')
		e.Buf = append(e.Buf, tag.Key...)
		e.Buf = append(e.Buf, `":"`...)
		e.Buf = append(e.Buf, tag.Value...)
		e.Buf = append(e.Buf, `",`...)
	}
	e.Buf[len(e.Buf)-1] = '}'
	e.Buf = append(e.Buf, `},`...)
}
