# metricsd

Simple metrics aggregation service:


# Client

For local testing gRPC visual client is provided

It's located under `cmd/demo-client` and after running you should be about
to see its results under following paths:

* Go to:
    * http://localhost:8080/MetricsService/ReportProcessMeasurement
    * http://localhost:8080/MetricsService/ReportNodeMeasurement
    * http://localhost:8080/MetricsService/NodeAverages
    * http://localhost:8080/MetricsService/ProcessesAverages
    * http://localhost:8080/MetricsService/ProcessAverage

# TODO

* [x] gitlab CI + auto registry push
* [x] gitlab CI for database
* [x] implement basic metrics pushing
* [ ] Add golden file test
* [ ] k8s/helm manifests
* [ ] Add metrics quering
* [ ] Add golden file test
* [ ] Add pod local memcache
* [ ] Optimize insertions with prepared queries

# Deployment

* Any http2 caching proxy can be put in front of this service
* Backend requires timescale database. Scaling + HA is similar to postgres database,
    sync/async replication + read only slave replicas
