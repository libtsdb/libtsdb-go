# Protocol

- http://druid.io/docs/0.11.0/ingestion/data-formats.html

## Write HTTP

- [ ] TODO: http://localhost:8200/v1/post/metrics it suggest using its own client, no doc for public api?

````json
{"timestamp": "2013-08-31T01:02:33Z", "page": "Gypsy Danger", "language" : "en", "user" : "nuclear", "unpatrolled" : "true", "newPage" : "true", "robot": "false", "anonymous": "false", "namespace":"article", "continent":"North America", "country":"United States", "region":"Bay Area", "city":"San Francisco", "added": 57, "deleted": 200, "delta": -143}
````