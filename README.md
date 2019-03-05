# metricsd

Simple metrics aggregation service:


# Client

Use https://github.com/gogo/letmegrpc

Install prereq:
* `GO111MODULE=off go get  github.com/gogo/protobuf/...`
* ...

* `letmegrpc --addr=localhost:5555 --port=8080 service.proto `
* Go to:
    * http://localhost:8080/MetricsService/ReportProcessMeasurement
    * http://localhost:8080/MetricsService/ReportNodeMeasurement
    * http://localhost:8080/MetricsService/NodeAverages
    * http://localhost:8080/MetricsService/ProcessesAverages
    * http://localhost:8080/MetricsService/ProcessAverage

# TODO

* [*] gitlab CI + auto registry push
* [ ] gitlab CI for database
* [ ] implement basic metrics pushing
* [ ] Add golden file test
* [ ] k8s/helm manifests
* [ ] Add metrics quering
* [ ] Add golden file test
* [ ] Add pod local memcache

# Deployment

* Any http2 caching proxy can be put in front of this service
* Backend requires timescale database. Scaling + HA is similar to postgres database,
    sync/async replication + read only slave replicas
