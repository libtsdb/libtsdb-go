# InfluxDB Data Model

- A major difference is InfluxDB has the concept of field, which means multiple value can be associated with a timestamp,
while most other TSDB are just one timestamp with one value, it seems they are changing towards this direction in the new
[ifql](https://github.com/influxdata/ifql) and use `join` instead
- Also it supports `bool` and `string` which is rare is most TSDB but pretty common in most column databases, i.e. Druid

Its internal struct for communicate protocol and storage is different

## Protocol

From https://github.com/influxdata/influxdb/blob/master/models/points.go

- line protocol example `temperature,machine=unit42,type=assembly internal=32,external=100 1434055562000000035`
- `Point` is a interface instead of a struct
  - `measurement` is what commonly called `series name`, i.e. `cpu`
- `point` is the default implementation, mapping to the line protocol
  - `key` is measurement + tags ` temperature,machine=unit42,type=assembly`
  - `fields` is field names + values `internal=32,external=100`
  - `ts` is timestamp in string `1434055562000000035`
     - using binary format would be much smaller, also text format can use Base64 VLQ like js source map, though it might not be suitable for timestamp

````go
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
````

## Storage

- in `func (e *Engine) WritePoints(points []models.Point)` `Point` is changed into `Value`, which is timestamp and value, like most TSDB
- series name is `measurement` + `tags` + `field`, see [write-path](write-path.md)
- https://github.com/influxdata/influxdb/blob/master/tsdb/engine/tsm1/encoding.go#L97-L113

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

// IntegerValue represents an int64 value.
type IntegerValue struct {
	unixnano int64
	value    int64
}

// BooleanValue represents a boolean value.
type BooleanValue struct {
	unixnano int64
	value    bool
}

// UnsignedValue represents an int64 value.
type UnsignedValue struct {
	unixnano int64
	value    uint64
}

// StringValue represents a string value.
type StringValue struct {
	unixnano int64
	value    string
}
````