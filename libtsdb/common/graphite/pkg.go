package graphite

import (
	"github.com/libtsdb/libtsdb-go/libtsdb"
	"time"
)

const (
	name      = "graphite"
	precision = time.Second
)

var meta = libtsdb.Meta{
	Name:               name,
	TimePrecision:      precision,
	SupportTag:         true,
	SupportInt:         false,
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
