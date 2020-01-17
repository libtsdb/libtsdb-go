# Protocol

There is no way to write to prometheus directly, but prometheus can write to external system

````proto
message Sample {
  double value    = 1;
  int64 timestamp = 2;
}

// TimeSeries represents samples and labels for a single time series.
message TimeSeries {
  repeated Label labels   = 1 [(gogoproto.nullable) = false];
  repeated Sample samples = 2 [(gogoproto.nullable) = false];
}

message Label {
  string name  = 1;
  string value = 2;
}

message Labels {
  repeated Label labels = 1 [(gogoproto.nullable) = false];
}

// Chunk represents a TSDB chunk.
// Time range [min, max] is inclusive.
message Chunk {
  int64 min_time_ms = 1;
  int64 max_time_ms = 2;

  // We require this to match chunkenc.Encoding.
  enum Encoding {
    UNKNOWN = 0;
    XOR     = 1;
  }
  Encoding type  = 3;
  bytes data     = 4;
}

// ChunkedSeries represents single, encoded time series.
message ChunkedSeries {
  // Labels should be sorted.
  repeated Label labels = 1 [(gogoproto.nullable) = false];
  // Chunks will be in start time order and may overlap.
  repeated Chunk chunks = 2 [(gogoproto.nullable) = false];
}
````

## Remote Write 

- https://github.com/prometheus/prometheus/tree/master/prompb
- it seems there is no push down to remote storage, promql has aggregation, but it only has label matcher

````proto
message WriteRequest {
  repeated prometheus.TimeSeries timeseries = 1;
}
````

There is [push gateway](https://github.com/prometheus/pushgateway), but it's a standalone server that expose metric.
Should be able to test other system like thanos etc. Or write a server using tsdb directly.