# Protocol

- https://github.com/akumuli/Akumuli/wiki/Writing-data-using-the-TCP-API

## TCP Write

non grouped series

- `:` for integer timestamp (unix nano)
- `+` for string value and string time

````text
+balancers.cpuload host=machine1 region=NW
:1418224205000000000
+22.0
````

bulk NOTE: **it is grouping series with same tag**

> Akumuli assumes that measurements with the same set of tags came from the same object. 
These measurements will differ only by metric names. 
E.g. "mem.usage host=machine1 region=NW" and "cpu.user host=machine1 region=NW" will be considered originating from the same host. 
That host is described by the set of tags - "host=machine1 region=NW" and metric name can be seen as a column name. 
Usually, it is preferable to write these metrics together and Akumuli has special message format for this case.

this is very similar to InfluxDB's style

````text
+cpu.real|cpu.user|cpu.sys host=machine1 region=NW
+20141210T074343
*3
+3.12
+8.11
+12.6
````

````text
+cpu.sys host=machine1 region=NW
*3
+20141210T074343
+3.12
+20141210T074344
+8.11
+20141210T074345
+12.6
````

- [ ] opened a issue to ask if it is possible to support https://github.com/akumuli/Akumuli/issues/260

## HTTP Read

- output format, csv or resp

````json
{
    "select": "balancers.cpuload",
		 "range": {
        "from": "20120102T123000.000000",
        "to":   "20190102T123010.000000"
    }
}
````

````text
+balancers.cpuload host=machine1 region=NW
+20141210T151005.000000000
+22
````
