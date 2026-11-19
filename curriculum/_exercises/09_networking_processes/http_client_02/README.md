# http_client_02: POSTing JSON with a deadline

`client.Get` and `client.Post` are shortcuts. For anything more (a context, headers), build
the request yourself:

```go
req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
req.Header.Set("Content-Type", "application/json")
resp, err := client.Do(req)
```

## Your task

Implement `CreateUser(ctx, client, baseURL, u) (id string, err error)`:

- POST `u` as JSON to `baseURL + "/users"` with `Content-Type: application/json`.
- Success is `201 Created` with a body like `{"id": "u-42"}`. Return that id.
- Any other status is an error.
- Respect `ctx`: if it expires, return promptly with an error wrapping
  `context.DeadlineExceeded`.

Reference: https://gobyexample.com/http-client
