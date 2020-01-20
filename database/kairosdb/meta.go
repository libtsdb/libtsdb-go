package kairosdb

import (
	"time"

	"github.com/libtsdb/libtsdb-go/database"
)

const (
	name      = "kairosdb"
	precision = time.Millisecond
)

var meta = database.Meta{
	Name:               name,
	TimePrecision:      precision,
	SupportTag:         true,
	SupportInt:         true,
	SupportDouble:      true,
	SupportBatchSeries: true,
	SupportBatchPoints: true,
}

func Meta() database.Meta {
	return meta
}

func init() {
	database.RegisterMeta(name, meta)
}
