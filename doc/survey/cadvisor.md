# google cAdvisor Container monitoring

- [ ] TODO: already have a survey for agent https://github.com/benchhub/benchboard/issues/2 but not for storage
- https://github.com/google/cadvisor
- https://github.com/google/cadvisor/blob/master/storage/storage.go
  - https://github.com/google/cadvisor/tree/master/info has v1 and v2
  - [x] what's the difference between v1 and v2, v2 is cleaner
- statsd
  - `fmt.Sprintf("%s.%s.%s:%d|g", namespace, containerName, key, value)`
- redis
  - use JSON, just serialize `ContainerStats` entirely
  - `self.conn.Send("LPUSH", self.redisKey, seriesToFlush)` 
- kafka, use JSON, same as redis ...
- influxdb
  - influxdb.Point, measurement, tags etc.
- elasticsearch, use JSON, same as redis ...
- bigquery, put everything in a row

````go
package storage

import (
	"fmt"
	"sort"

	info "github.com/google/cadvisor/info/v1"
)

type StorageDriver interface {
	AddStats(ref info.ContainerReference, stats *info.ContainerStats) error

	// Close will clear the state of the storage driver. The elements
	// stored in the underlying storage may or may not be deleted depending
	// on the implementation of the storage driver.
	Close() error
}
````

````go
type Percentiles struct {
	// Indicates whether the stats are present or not.
	// If true, values below do not have any data.
	Present bool `json:"present"`
	// Average over the collected sample.
	Mean uint64 `json:"mean"`
	// Max seen over the collected sample.
	Max uint64 `json:"max"`
	// 50th percentile over the collected sample.
	Fifty uint64 `json:"fifty"`
	// 90th percentile over the collected sample.
	Ninety uint64 `json:"ninety"`
	// 95th percentile over the collected sample.
	NinetyFive uint64 `json:"ninetyfive"`
}

type ProcessInfo struct {
	User          string  `json:"user"`
	Pid           int     `json:"pid"`
	Ppid          int     `json:"parent_pid"`
	StartTime     string  `json:"start_time"`
	PercentCpu    float32 `json:"percent_cpu"`
	PercentMemory float32 `json:"percent_mem"`
	RSS           uint64  `json:"rss"`
	VirtualSize   uint64  `json:"virtual_size"`
	Status        string  `json:"status"`
	RunningTime   string  `json:"running_time"`
	CgroupPath    string  `json:"cgroup_path"`
	Cmd           string  `json:"cmd"`
}
````