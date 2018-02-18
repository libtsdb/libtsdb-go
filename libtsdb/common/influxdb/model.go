package influxdb

import "time"

// TODO: why influxdb is using all []byte instead of string, to avoid copy?

type Tag struct {
	Key   string
	Value string
}

type Fields map[string]interface{}

type Point struct {
	Time   time.Time
	Name   string
	Tags   []Tag
	Fields Fields
}

// BatchPoints is based on https://github.com/influxdata/influxdb/blob/master/client/v2/client.go
// TODO: we don't have access to influxdb cluster edition, so writeConsistency is useless,
// also retentionPolicy is not useful since we may not run the benchmark long enough to trigger it
type BatchPoints struct {
	Points           []*Point
	Database         string
	Precision        string
	RetentionPolicy  string
	WriteConsistency string
}

func NewPoint() *Point {
	return &Point{
		Time: time.Now(),
	}
}

// temperature,machine=unit42,type=assembly internal=32,external=100 1434055562000000035
func (p *Point) toLineProtocol() {

}
