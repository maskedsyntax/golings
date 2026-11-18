# url_parsing_01: Taking a URL apart

`net/url` parses URLs into their parts:

```text
https://api.example.com:8443/v1/items?token=abc&x=1
scheme  hostname        port path     query
```

## Your task

Implement `Parse(raw string) (Endpoint, error)`:

- `Scheme` is `http` or `https`. Anything else is an error, and so is a URL with no host.
- `Host` is the host name **without** the port. For `http://[::1]:9000/` that's `::1`.
- `Port` is the explicit port, or else `"80"` for http and `"443"` for https.
- `Path` is the path (`""` if there is none).
- `Token` is the `token` query parameter (`""` if absent).

Reference: https://gobyexample.com/url-parsing
