# Design

Draft

## directory layout

- common, models, protobuf etc.
- storage, wrapper of existing storage engine or compatible implementation (not likely to work on it for a while, unless all wrapper)
  - graphite
  - tsm
- server, server protocol implementation, decode into internal representation, used for accepting requests from old clients during migration. can also be used from tsdb that don't have push API (i.e. prometheus)
- client, client wrapper or protocol compatible implementation (if the official client is small), used for unified interface of client side, so swap client just need to change config (maybe build flag as well), mainly used for benchmark