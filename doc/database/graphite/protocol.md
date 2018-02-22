# Protocol

- http://graphite.readthedocs.io/en/latest/feeding-carbon.html#feeding-in-your-data
  - http://graphite.readthedocs.io/en/latest/tools.html
  - https://github.com/influxdata/telegraf/blob/master/plugins/outputs/graphite/graphite.go

## TCP Plain text

`<metric path> <metric value> <metric timestamp>`

- [ ] TODO: it seems it can also support UDP
- `nc localhost 2003`
- `local.random 4 1519266078`