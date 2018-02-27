# Protocol

There is no way to write to prometheus directly, but prometheus can write to external system

## Remote Write 

- https://github.com/prometheus/prometheus/tree/master/prompb
- it seems there is no push down to remote storage, promql has aggregation, but it only has label matcher

````proto
message Sample {
  double value    = 1;
  int64 timestamp = 2;
}

message TimeSeries {
  repeated Label labels   = 1;
  repeated Sample samples = 2;
}

message Label {
  string name  = 1;
  string value = 2;
}

message Labels {
  repeated Label labels = 1 [(gogoproto.nullable) = false];
}

// Matcher specifies a rule, which can match or set of labels or not.
message LabelMatcher {
  enum Type {
    EQ  = 0;
    NEQ = 1;
    RE  = 2;
    NRE = 3;
  }
  Type type    = 1;
  string name  = 2;
  string value = 3;
}
````

````proto
message WriteRequest {
  repeated prometheus.TimeSeries timeseries = 1;
}

message ReadRequest {
  repeated Query queries = 1;
}

message ReadResponse {
  // In same order as the request's queries.
  repeated QueryResult results = 1;
}

message Query {
  int64 start_timestamp_ms = 1;
  int64 end_timestamp_ms = 2;
  repeated prometheus.LabelMatcher matchers = 3;
}

message QueryResult {
  // Samples within a time series must be ordered by time.
  repeated prometheus.TimeSeries timeseries = 1;
}
````

There is [push gateway](https://github.com/prometheus/pushgateway), but it's a standalone server