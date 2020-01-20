package influxdb

import (
	"time"

	"github.com/libtsdb/libtsdb-go/database"
)

const (
	name      = "influxdb"
	precision = time.Nanosecond
)

var meta = database.Meta{
	Name:               name,
	TimePrecision:      precision,
	SupportTag:         true,
	SupportInt:         true,
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
