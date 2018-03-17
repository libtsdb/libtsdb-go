#!/usr/bin/env bash

# https://hub.docker.com/r/akumuli/akumuli/

docker run -p 8181:8181 -p 8282:8282 -p 8383:8383 -p 4242:4242 \
       -e nvolumes='16' -e volume_size='1GB' \
       akumuli/akumuli:skylake

#  curl -XGET "http://localhost:8181/api/stats"