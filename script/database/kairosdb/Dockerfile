FROM java:8-alpine

MAINTAINER at15 at15@dongyue.io

# Thanks to @zbintliff in https://github.com/kairosdb/kairosdb/issues/288
RUN apk upgrade libssl1.0 --update-cache && \
    apk add curl ca-certificates bash

RUN mkdir /opt; \
  cd /opt; \
  curl -L https://github.com/kairosdb/kairosdb/releases/download/v1.2.0/kairosdb-1.2.0-1.tar.gz | \
  tar zxvfp -

# Use Cassandra
COPY kairosdb.properties /opt/kairosdb/conf/kairosdb.properties

# KairosDB must start after Cassandra is running and thrift protocol is enabled
#COPY wait-for-it.sh /usr/bin/wait-for-it
COPY waitforit /usr/bin/waitforit
COPY kairosdb.sh /usr/bin/kairosdb

EXPOSE 4242 8080 2003 2004

ENTRYPOINT []
CMD [ "/usr/bin/kairosdb" ]
