# libtsdb-go

Clients and Server implementation of multiple TSDB protocols in Go

## Motivation

- [xephonhq/xephon-b](https://github.com/xephonhq/xephon-b) a TSDB benchmark suites need clients for various TSDB to do benchmark
- [xephonhq/xephon-k](https://github.com/xephonhq/xephon-k) a TSDB w/ multiple backends needs servers to accept data from collectors
- [google/cAdvisor](https://github.com/google/cadvisor) a container metrics collector has too many storage specific code 
- Add tracing to existing go-base tsdb storage engines
  - https://github.com/prometheus/tsdb
  - https://github.com/influxdata/influxdb
  
## License

MIT

