# Protocol

## Write HTTP

- https://spotify.github.io/heroic/#!/docs/api/post-write
- very similar to KairosDB
- [ ] TODO: only one series in one request?

````json
{
  "series": {"key": "foo", "tags": {"site": "lon", "host": "www.example.com"}},
  "data": {"type": "points", "data": [[1300000000000, 42.0], [1300001000000, 84.0]]}
}
````