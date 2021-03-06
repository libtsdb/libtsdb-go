# Protocol

## Write Telnet

- https://kairosdb.github.io/docs/build/html/telnetapi/Putm.html
- `putm` requires millisecond, `put` support both second and millisecond
  - `putm` is NOT for putting multiple values ...

````text
putm <metric name> <time stamp> <value> <tag> <tag>... \n
putm load_value_test 1521355855810716 42 host=A
````

## Write HTTP

- http://localhost:8080/api/v1/datapoints
- it has to form for points
  - `[1359788400000, 123]` first is ts in **milliseconds**
  - `{"timestamp": 1359786400000, "value": 321}` should be from OpenTSDB
- https://kairosdb.github.io/docs/build/html/restapi/AddDataPoints.html

````json
[
  {
      "name": "archive_file_tracked",
      "datapoints": [[1359788400000, 123], [1359788300000, 13.2], [1359788410000, 23.1]],
      "tags": {
          "host": "server1",
          "data_center": "DC1"
      },
      "ttl": 300
  },
  {
      "name": "archive_file_search",
      "timestamp": 1359786400000,
      "value": 321,
      "tags": {
          "host": "server2"
      }
  }
]
````