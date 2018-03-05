package influxdb

import (
	"github.com/libtsdb/libtsdb-go/libtsdb/common"
	"time"
)

const (
	name      = "influxdb"
	precision = time.Nanosecond
)

var meta = common.Meta{
	Name:          name,
	TimePrecision: precision,
	SupportTag:    true,
	SupportInt:    true,
	SupportDouble: true,
}

func Meta() common.Meta {
	return meta
}

func init() {
	common.RegisterMeta(name, meta)
}
