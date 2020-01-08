# Data Model

We need to figure out a internal data representation, and how to convert into database specific structs

- issue https://github.com/libtsdb/libtsdb-go/issues/1

- [InfluxDB](influxdb/model.md) 
  - protocol `name: string, tags: []{k: []byte, v:[]byte}, fields: map[string]interface{}`
  - storage `key` = `name` + `tags` + `field`, `value` = `ts` + `int`|`float`|`bool`|`string`

Prometheus

Google

- https://github.com/googleapis/googleapis/blob/master/google/monitoring/v3/metric.proto
