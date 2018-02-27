#!/usr/bin/env bash

# NOTE: we don't map 80 in guest to host 80
docker run \
 -p 8090:80\
 -p 2003-2004:2003-2004\
 -p 2023-2024:2023-2024\
 -p 8125:8125/udp\
 -p 8126:8126\
 graphiteapp/graphite-statsd