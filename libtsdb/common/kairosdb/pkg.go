package kairosdb

import (
	"time"

	"github.com/libtsdb/libtsdb-go/libtsdb"
)

const (
	name      = "kairosdb"
	precision = time.Millisecond
)

var meta = libtsdb.Meta{
	Name:               name,
	TimePrecision:      precision,
	SupportTag:         true,
	SupportInt:         true,
	SupportDouble:      true,
	SupportBatchSeries: true,
	SupportBatchPoints: true,
}

func Meta() libtsdb.Meta {
	return meta
}

func init() {
	libtsdb.RegisterMeta(name, meta)
}
