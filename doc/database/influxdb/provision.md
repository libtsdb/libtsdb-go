# Provision

- https://docs.influxdata.com/influxdb/v1.4/query_language/data_exploration/

````bash
docker exec -it f953 influxdb
influx
create database libtsdbtest
use libtsdbtest
show measurements;
select * from temperature;
````