# libtsdb-go

Clients and Server implementation of multiple TSDB protocols in Go

- status: Under major rewrite

## Supported time series databases

| Database     | client write  | client read   | server write | server read |
| ------------ |:------:| :-----:| :----------: | :--: |
| Akumuli  | TCP/RESP | N | N | N |
| InfluxDB | HTTP/Line | N | N | N |
| Graphite | TCP/Line | N | N | N |
| KairosDB | HTTP/JSON  TCP/Line | N | N | N |
| Xephon-K | (TODO) TCP HTTP/JSON GRPC | N | N | N |
| OpenTSDB | HTTP/JSON (NA) | N | N | N |
| Heroic | HTTP/JSON (NA) | N | N | N |

## Motivation

- [xephonhq/xephon-b](https://github.com/xephonhq/xephon-b) a TSDB benchmark suites need clients for various TSDB to do benchmark
- [xephonhq/xephon-k](https://github.com/xephonhq/xephon-k) a TSDB w/ multiple backends needs servers to accept data from collectors
- [google/cAdvisor](https://github.com/google/cadvisor) a container metrics collector has too many storage specific code 
- Add tracing to existing go-base tsdb storage engines
  - https://github.com/prometheus/tsdb
  - https://github.com/influxdata/influxdb

## Metrics

To be used with benchmark like [xephon-b](https://github.com/xephonhq/xephon-b), 
the client (both http and tcp) return struct that meet `libtsdb.Trace` interface
to avoid user add extra wrapper to get latency etc.
http client uses `net/http/httptrace` and can give more detailed data.

For a single request

- request start time
- request end time (for http, this include reading response)
- series, points written
- raw size (data + meta)
- raw meta size (series name, tags)
- payload size

Accumulated results

- points written
- raw size (data + meta)
- raw meta size (series name, tags)
- payload size

## Roadmap

- [ ] archive current branch to archive/xxx
- [ ] use go mod and fix dependencies
- [ ] move survey to awesome tsdb
- [ ] compression etc. (same as libtsdb-rs from here)
- client write
  - [x] simple text line protocol, InfluxDB, Graphite, Akumuli
  - [x] OpenTSDB(ish) JSON, KairosDB, Heroic, Xephon-K
    - Xephon-K no longer use JSOn, OpenTSDB and Heroic don't have handy docker images
  - [x] GRPC, Xephon-K
  - [ ] Thrift, Gorilla
- client read
  - [ ] InfluxDB
  - [ ] OpenTSDB(ish) JSON, KairosDB, Heroic, Xephon-K
  - [ ] Prometheus?
- server write

## License

MIT

