# Protocol

- http://graphite.readthedocs.io/en/latest/feeding-carbon.html#feeding-in-your-data
  - http://graphite.readthedocs.io/en/latest/tools.html
  - https://github.com/influxdata/telegraf/blob/master/plugins/outputs/graphite/graphite.go

## TCP Plain text

`<metric path> <metric value> <metric timestamp>`

- http://graphite.readthedocs.io/en/latest/feeding-carbon.html#step-1-plan-a-naming-hierarchy 
- [ ] TODO: it seems it can also support UDP
- [ ] TODO: value is int? float?
- `nc localhost 2003`
- `local.random 4 1519266078`

> A tagged series is made up of a name and a set of tags, like “disk.used;datacenter=dc1;rack=a1;server=web01”

## TCP Pickle

Pickle is a python serialization format, though we can hand construct it ...

- https://github.com/lomik/graphite-pickle usd 