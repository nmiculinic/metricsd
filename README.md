# metricsd

Simple metrics aggregation service:

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
