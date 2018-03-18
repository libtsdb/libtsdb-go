# libtsdb-go

Clients and Server implementation of multiple TSDB protocols in Go

## Supported time series databases

| Database     | client write  | client read   | server write | server read |
| ------------ |:------:| :-----:| :----------: | :--: |
| Akumuli  | TCP/RESP | N | N | N |
| InfluxDB | HTTP/Line | N | N | N |
| Graphite | TCP/Line | N | N | N |
| KairosDB | HTTP/JSON  TCP/Line | N | N | N |
| Xephon-K | HTTP/JSON (TODO) GRPC | N | N | N |
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
the client (both http and tcp) return struct that fits the [interface](https://github.com/xephonhq/xephon-b/tree/master/pkg/metrics)
to avoid a wrapper around them.
http client use `net/http/httptrace` and give more detailed data.

For a single request

- request start time
- latency 
- series, points written
- raw data size
- actual payload size

Accumulated results

- [ ] TODO:

## Roadmap

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

