#!/usr/bin/env bash

echo "waiting for cassandra to start"
# TODO： timeout is not included
#wait-for-it cassandra:9042
ls
/usr/bin/waitforit -w tcp://cassandra:9042
echo "cassandra started"
/opt/kairosdb/bin/kairosdb.sh run
