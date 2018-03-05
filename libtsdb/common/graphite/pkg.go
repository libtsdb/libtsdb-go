package graphite

import (
	"github.com/libtsdb/libtsdb-go/libtsdb/common"
	"time"
)

const (
	name      = "graphite"
	precision = time.Second
)

var meta = common.Meta{
	Name:          name,
	TimePrecision: precision,
	SupportTag:    true,
	SupportInt:    false,
	SupportDouble: true,
}

func Meta() common.Meta {
	return meta
}

func init() {
	common.RegisterMeta(name, meta)
}
