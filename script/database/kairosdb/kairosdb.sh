#!/usr/bin/env bash

echo "waiting for cassandra to start"
# TODO： timeout is not included
#wait-for-it cassandra:9042
# using new waitforit https://github.com/benchhub/benchhub/issues/20
/usr/bin/waitforit -w tcp://cassandra:9042
echo "cassandra started"
/opt/kairosdb/bin/kairosdb.sh run
