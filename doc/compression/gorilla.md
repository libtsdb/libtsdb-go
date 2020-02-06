# Facebook Gorilla

## Paper

http://www.vldb.org/pvldb/vol8/p1816-teller.pdf

4.1 Time series compression

- timestamps and values are compressed separately using information about previous values
  - but they are put into same byte stream

4.1.1 Compressing time stamps 

- first value is aligned to two hour window
- second value is delta with first value, size is 14 bits because, 14 bits is 16384 seconds, 4.5h
- use a dictionary, the range of dictionary is determined by sample

4.1.2 Compressing values

- first XOR w/ previous value
- variable length encoding
 
## Beringei

https://github.com/facebookarchive/beringei

## Prometheus

It is defined in Append and is modified from go-tsz

- [prometheus/promtheus/tsdb/chunkenc/xor.go](https://github.com/prometheus/prometheus/blob/7cf09b0395125280eb3e2d44b603349aeecacec1/tsdb/chunkenc/xor.go#L137)
- it has the same problem of not able to read same chunk twice because it is using the bit stream code from go-tsz, see [fork](https://github.com/at15/prometheus/commit/018948da1dbf8533fd63328744ac55e2e3dbce3c)
  - [ ] but I remember it is using mmap, so how come the logic didn't break anything ...

> // The code in this file was largely written by Damian Gryski as part of
  // https://github.com/dgryski/go-tsz and published under the license below.
  // It was modified to accommodate reading from byte slices without modifying
  // the underlying bytes, which would panic when reading from mmap'd
  // read-only byte slices.

## go-tsz

- [x] it seems the bit reader destroy the underlying bit slices when read, it shifts bytes when read is no aligned with byte boundary
  - [test](https://github.com/at15/go-tsz/commit/c469915e5694a0541965f396edc14e6a88bc9bb7)
  
## InfluxDB

- they use the XOR for float64 value, for time it uses delta RLE, simple8b or full precision

## Victoriametrics

- https://github.com/VictoriaMetrics/VictoriaMetrics/tree/master/lib/encoding well need to see how they use it ...
 -[VictoriaMetrics: achieving better compression than Gorilla for time series data](https://medium.com/faun/victoriametrics-achieving-better-compression-for-time-series-data-than-gorilla-317bc1f95932)

## m3db

- [M3TSZ](https://github.com/m3db/m3/blob/b27738bb35578fff396a67dfc2797f972f203e5f/docs/m3db/architecture/engine.md#time-series-compression-m3tsz)
- time is double delta but it allows changing time unit [m3tsz/timestamp_encoder.go](https://github.com/m3db/m3/blob/master/src/dbnode/encoding/m3tsz/timestamp_encoder.go)

## timescaledb

- [Release](https://github.com/timescale/timescaledb/releases/tag/1.5.0)
- [PR](https://github.com/timescale/timescaledb/pull/1434)
- [Blog: Building columnar compression in a row-oriented database](https://blog.timescale.com/blog/building-columnar-compression-in-a-row-oriented-database/)
- [double delta](https://github.com/timescale/timescaledb/blob/master/tsl/src/compression/deltadelta.c#L371)
  - the double delta value is zig zag encoded, not using gorilla's dictionary way.
- [xor](https://github.com/timescale/timescaledb/blob/master/tsl/src/compression/gorilla.c#L394)