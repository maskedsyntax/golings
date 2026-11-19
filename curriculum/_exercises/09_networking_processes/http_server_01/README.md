# http_server_01: Routing with ServeMux

A handler is anything with `ServeHTTP(w http.ResponseWriter, r *http.Request)`. A
`http.ServeMux` routes requests to handlers by method and path:

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /items/{id}", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "item %s\n", r.PathValue("id"))
})
```

## Your task

Implement `NewMux() http.Handler`:

| request | response |
|---|---|
| `GET /hello/{name}` | `200`, body `Hello, <name>!\n` |
| `POST /echo` | `200`, body = the request body |
| any other method on `/echo` | `405 Method Not Allowed` |
| anything else | `404 Not Found` |

The tests call your handler directly with `httptest.NewRecorder`, so no network is needed.

Reference: https://gobyexample.com/http-server
