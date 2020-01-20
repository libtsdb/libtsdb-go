package akumuli

import (
	"time"

	"github.com/libtsdb/libtsdb-go/database"
)

const (
	name      = "akumuli"
	precision = time.Nanosecond
)

var meta = database.Meta{
	Name:               name,
	TimePrecision:      time.Nanosecond,
	SupportTag:         true,
	SupportInt:         true, // TODO: the protocol support int, but is it stored as float?
	SupportDouble:      true,
	SupportBatchSeries: true,
	SupportBatchPoints: false,
}

func Meta() database.Meta {
	return meta
}

func init() {
	database.RegisterMeta(name, meta)
}
