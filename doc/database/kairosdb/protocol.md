# Protocol

## Write Telnet

## Write HTTP

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