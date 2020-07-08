# Databases

Deprecated: the content is being moved to [notes-on-tsdb](https://github.com/xephonhq/notes-on-tsdb).

This folder contains notes on implementation of TSDBs.
They should be merged into [awesome-time-series-database](https://github.com/xephonhq/awesome-time-series-database) eventually.

Required list in readme

- link to [awesome tsdb](https://github.com/xephonhq/awesome-time-series-database), which should contains all the basic meta
- link to required files within the folder
- a short overview
- a key take way from this database

Required files for describing each database.

Code walk through

- **read.md** Read path, link to source w/ commit hash.
- **write.md** Write path, link to source w/ commit hash.

API

- **protocol.md** Wire protocol format and transport, mainly about write because many TSDB have dedicated query language.
- **query-language.md** Query language.

Internal

- **model.md** General data model, what is a time series for this TSDB (yeah, this definition varies).
- **compression.md** Compression related algorithm or code.
- **query-execution.md** Query execution and optimization, especially for those with query language and distributed ones.
- **storage-engine.md** Only applies to TSDB w/ their own storage format, i.e. write opaque blob to local fs or object store.
- **schema.md** Only applies to TSDB w/ underlying database i.e. Cassandra, ElasticSearch
- **distributed.md** Only applies to distributed TSDB, replication model, consensus protocol. Including those built on top of distributed data store i.e. Cassandra, S3.

Operation

- **build.md** How to build from source locally
- **docker.md** How to run it using docker(-compose) locally
- **config.md** Config file example, how to config the system (underlying database, operating system) properly
- **k8s.md** How to run it on k8s, operator and special things about their operator e.g. local volume

## TODO

- [ ] publish the doc or just move it to awesome-tsdb?