# http_client_01: Status codes and closing bodies

`net/http` makes requests easy:

```go
resp, err := client.Get(url)
if err != nil { ... }        // network failure, bad URL...
defer resp.Body.Close()      // always
if resp.StatusCode != http.StatusOK { ... }
```

Two rules trip people up. An HTTP error status is **not** a Go error, and every body must
be **closed**.

`FetchTitle` gets a JSON document and returns its `title`. For a 404 it returns `""` and no
error, and it never closes the body.

## Your task

Fix `FetchTitle`:

- Return an error that mentions the status for any response other than 200 OK.
- Always close the response body.

Reference: https://gobyexample.com/http-client
