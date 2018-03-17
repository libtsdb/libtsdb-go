# libtsdb-go

Clients and Server implementation of multiple TSDB protocols in Go

## Supported time series databases

| Database     | client write  | client read   | server write | server read |
| ------------ |:------:| :-----:| :----------: | :--: |
| Akumuli  | TCP/Line | N | N | N |
| InfluxDB | HTTP/Line | N | N | N |
| Graphite | TCP/Line | N | N | N |
| KairosDB | HTTP/JSON | N | N | N |
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

## Roadmap

- client write
  - [x] simple text line protocol, InfluxDB, Graphite
  - [ ] OpenTSDB(ish) JSON, KairosDB, Heroic, Xephon-K
  - [ ] GRPC, Xephon-K
  - [ ] Thrift, Gorilla
- client read
  - [ ] InfluxDB
  - [ ] OpenTSDB(ish) JSON, KairosDB, Heroic, Xephon-K
  - [ ] Prometheus?
- server write

## License

MIT

