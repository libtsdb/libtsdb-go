package akumuli

import (
	"time"

	"github.com/libtsdb/libtsdb-go/libtsdb"
)

const (
	name      = "akumuli"
	precision = time.Nanosecond
)

var meta = libtsdb.Meta{
	Name:               name,
	TimePrecision:      precision,
	SupportTag:         true,
	SupportInt:         true, // TODO: the protocol support int, but is it stored as float?
	SupportDouble:      true,
	SupportBatchSeries: true,
	SupportBatchPoints: false,
}

func Meta() libtsdb.Meta {
	return meta
}

func init() {
	libtsdb.RegisterMeta(name, meta)
}
