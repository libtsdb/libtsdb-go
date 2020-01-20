package database

// meta.go keeps a database of databases at runtime

import (
	"log"
	"sort"
	"sync"
	"time"

	"github.com/dyweb/gommon/errors"
)

var (
	metaMu    sync.Mutex
	metas     = make(map[string]Meta)
	emptyMeta = Meta{Name: "empty"}
)

// Meta describes a database's behavior
type Meta struct {
	Name          string
	Repo          string
	TimePrecision time.Duration
	SupportTag    bool
	SupportInt    bool
	SupportDouble bool
	// write multiple series in one payload
	SupportBatchSeries bool
	// merge points of same series into more compact format
	SupportBatchPoints bool
}

func RegisterMeta(db string, meta Meta) {
	metaMu.Lock()
	defer metaMu.Unlock()
	if _, dup := metas[db]; dup {
		log.Printf("RegisterMeta is called twice for %s", db)
	}
	metas[db] = meta
}

func GetDatabaseMeta(name string) (Meta, error) {
	metaMu.Lock()
	defer metaMu.Unlock()

	if m, ok := metas[name]; ok {
		return m, nil
	} else {
		return emptyMeta, errors.Errorf("database %s didn't register meta", name)
	}
}

func Databases() []string {
	metaMu.Lock()
	defer metaMu.Unlock()
	var list []string
	for db := range metas {
		list = append(list, db)
	}
	sort.Strings(list)
	return list
}
