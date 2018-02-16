# k8s Heapster

- https://kubernetes.io/docs/tasks/debug-application-cluster/resource-usage-monitoring/
  - gathered from kublet, which is using cadvisor https://github.com/kubernetes/kubernetes/tree/master/pkg/kubelet/cadvisor
  - https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/cadvisor/cadvisor_linux.go
    - it creates cadvisor manager directly, does NOT start a new container https://github.com/google/cadvisor/blob/master/manager/manager.go

- [docs/sink-configuration](https://github.com/kubernetes/heapster/blob/master/docs/sink-configuration.md)
  - some support metrics, InfluxDB, OpenTSDB, Honeycomb supports events
    - [ ] what are events?
- https://github.com/kubernetes/heapster/tree/master/metrics/sinks
  - graphite 
    - https://github.com/marpaia/graphite-golang which says it should not be used for production ....
  - opentsdb
    - https://github.com/bluebreezecf/opentsdb-goclient ... which is last updated on 2016 ...
  - riemann
    - https://github.com/riemann/riemann/tree/master/src/riemann actually has many clients ...
    - https://github.com/riemann/riemann-go-client
  - librato, commercial
  - statsd, dead
  - hawkular, Red Hat, dead, they are now working on jagear
  - wavefront, bought by vmware
- [docs/storage-schema.md](https://github.com/kubernetes/heapster/blob/master/docs/storage-schema.md)