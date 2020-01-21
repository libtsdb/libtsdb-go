# Protocol

The protocol is http based.

## Write HTTP Line protocol

https://v2.docs.influxdata.com/v2.0/write-data/

```text
mem,host=host1 used_percent=23.43234543 1556892576842902000
cpu,host=host1 usage_user=3.8234,usage_system=4.23874 1556892726597397000
mem,host=host1 used_percent=21.83599203 1556892777007291000
```

## Read HTTP JSON

````bash
curl -G 'http://localhost:8086/query?pretty=true' --data-urlencode "db=mydb" --data-urlencode "q=SELECT \"value\" FROM \"cpu_load_short\" WHERE \"region\"='us-west'"
````

````json
{
    "results": [
        {
            "statement_id": 0,
            "series": [
                {
                    "name": "cpu_load_short",
                    "columns": [
                        "time",
                        "value"
                    ],
                    "values": [
                        [
                            "2015-01-29T21:55:43.702900257Z",
                            2
                        ],
                        [
                            "2015-01-29T21:55:43.702900257Z",
                            0.55
                        ],
                        [
                            "2015-06-11T20:46:02Z",
                            0.64
                        ]
                    ]
                }
            ]
        }
    ]
}
````