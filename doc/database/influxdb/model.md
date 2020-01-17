# InfluxDB Data Model

## TODO

- [ ] Its internal struct for communicate protocol and storage is different
- [ ] compare example with other tsdb

## Overview

InfluxDB has both name, tag, fields and time, while most TSDB only has name, tags, a single value and time.

```text
temperature,machine=unit42,type=assembly internal=32,external=100 1434055562000000035
---------- ----------------------------- ------------------------ -------------------
  name           tags                        fields                   time
```

- `name` is called `measurement`
- `tags` are called `tag set`
- `fields` are called `fields set`
- `time` is unix nano

When saved on disk, each series is identified by `name+tags+field`, e.g. `temperature,machine=unit42,type=assembly internal`.

## Scalar value type

- bool
- int
- unsigned int
- float
- string

## Protocol

Based on [models/point.go](https://github.com/influxdata/influxdb/blob/master/models/points.go)

```text
temperature,machine=unit42,type=assembly internal=32,external=100 1434055562000000035
```

- `Point` is an interface instead of a struct
  - `measurement` is what commonly called `series name`, i.e. `cpu`
  - `tags` is list of string key value pair
  - `fields` is `map[string]interface{}` because type of value is unknown 
- `point` is the default implementation, mapping to the line protocol
  - `key` is measurement + tags ` temperature,machine=unit42,type=assembly`
  - `fields` is field names + values `internal=32,external=100`
  - `ts` is timestamp in string `1434055562000000035`
     - using binary format would be much smaller, also text format can use Base64 VLQ like js source map, though it might not be suitable for timestamp

```go
// Tag represents a single key/value tag pair.
type Tag struct {
	Key   []byte
	Value []byte
}

// Tags represents a sorted list of tags.
type Tags []Tag

type Fields map[string]interface{}

// Point defines the values that will be written to the database.
type Point interface {
	// Name return the measurement name for the point.
	Name() []byte
	// Tags returns the tag set for the point.
	Tags() Tags
	// Fields returns the fields for the point.
	Fields() (Fields, error)
	// Time return the timestamp for the point.
	Time() time.Time
    // HashID returns a non-cryptographic checksum of the point's key.
	HashID() uint64
	// Key returns the key (measurement joined with tags) of the point.
	Key() []byte
}

// point is the default implementation of Point.
type point struct {
	time time.Time
	// text encoding of measurement and tags
	// key must always be stored sorted by tags, if the original line was not sorted,
	// we need to resort it
	key []byte
	// text encoding of field data
	fields []byte
	// text encoding of timestamp
	ts []byte
	// cached version of parsed fields from data
	cachedFields map[string]interface{}
	// cached version of parsed name from key
	cachedName string
	// cached version of parsed tags
	cachedTags Tags
	it fieldIterator
}
```

## Storage

Write has three steps

- convert from line protocol to internal format `tsdb.NewSeriesCollection` and `tsm1.CollectionToValues`
- write to wal, it use snappy to compress bytes
- update index `e.index.CreateSeriesListIfNotExists(collection)`
- write values `e.engine.WriteValues(values)` which actually writes to cache
- [ ] on disk format and index format, it seems tsm2 is different from tsm1?

```go
func (e *Engine) WritePoints(ctx context.Context, points []models.Point) error {
    collection, j := tsdb.NewSeriesCollection(points), 0
    // Filter by tags etc.
    // Convert the collection to values for adding to the WAL/Cache.
	values, err := tsm1.CollectionToValues(collection)
	// Add the write to the WAL to be replayed if there is a crash or shutdown.
	if _, err := e.wal.WriteMulti(ctx, values); err != nil {  }
    return e.writePointsLocked(ctx, collection, values)
}
```

- `CollectionToValues` split a point with multiple fields into different series
  - `[]Point` is converted to `map[string][]Value`
  - map key is `name+tags+field`, NOTE: there is just one field in key
  - `Value` contains both timestamp and the scalar value

```text
Point
temperature,machine=unit42,type=assembly internal=32,external=100 1434055562000000035
temperature,machine=unit42,type=assembly internal=38,external=82 1434055562000000036

Map
{
    "temperature,machine=unit42,type=assembly#!~#internal": [(1434055562000000035, 32), (1434055562000000036,38)]
    "temperature,machine=unit42,type=assembly#!~#external": [(1434055562000000035, 100), (1434055562000000036, 92)]
}
```

```go
const keyFieldSeparator = "#!~#"

// tsdb/tsm1/value.go
func CollectionToValues(collection *tsdb.SeriesCollection) (map[string][]Value, error) {
    for citer := collection.Iterator(); citer.Next(); {
            // reset global buf and append key
            // key is name with tags e.g. temperature,machine=unit42,type=assembly
            keyBuf = append(keyBuf[:0], citer.Key()...)
            keyBuf = append(keyBuf, keyFieldSeparator...)
            baseLen = len(keyBuf)
            p := citer.Point()
            iter := p.FieldIterator()
            t := p.Time().UnixNano()
    
            // Loop each field, each field becomes a new series
            for iter.Next() {
                keyBuf = append(keyBuf[:baseLen], iter.FieldKey()...)
                var v Value
                vs, ok := values[string(keyBuf)]
                values[string(keyBuf)] = append(vs, v)
            }
    }
}
```

````go
// Value represents a TSM-encoded value.
type Value interface {
	// UnixNano returns the timestamp of the value in nanoseconds since unix epoch.
	UnixNano() int64
	// Value returns the underlying value.
	Value() interface{}
	// Size returns the number of bytes necessary to represent the value and its timestamp.
	Size() int
	// String returns the string representation of the value and its timestamp.
	String() string
	// internalOnly is unexported to ensure implementations of Value
	// can only originate in this package.
	internalOnly()
}

type FloatValue struct {
	unixnano int64
	value    float64
}
// ... and other types
// StringValue represents a string value.
type StringValue struct {
	unixnano int64
	value    string
}
````