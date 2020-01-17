# Prometheus Data Model

## TODO

- [ ] both the expose format and the internal (remote read/write) format

## Overview

Although Prometheus supports type like histogram, they are [flatted into untyped time series](https://prometheus.io/docs/concepts/metric_types/).
It is tagged series, name, tags, time and value

```text
metric_name [
  "{" label_name "=" `"` label_value `"` { "," label_name "=" `"` label_value `"` } [ "," ] "}"
] value [ timestamp ]
```

- value is float64
- time is unix millisecond
- [ ] is metric_name encoded as special label?

## Scalar value type

