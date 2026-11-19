# http_server_02: Graceful shutdown

A server that stops abruptly cuts off users mid-request. `http.Server` has two ways to stop:

- `Close`: drops every connection immediately.
- `Shutdown(ctx)`: stops accepting new connections, waits for in-flight requests to finish,
  and gives up when `ctx` is done.

After either one, `Serve` returns `http.ErrServerClosed`.

## Your task

Fix `Serve(ctx, ln, h, grace)`:

- Serve `h` on `ln` until `ctx` is cancelled.
- Then shut down gracefully, letting in-flight requests finish (waiting at most `grace`).
- Return `nil` after a clean shutdown, or the error if serving or shutting down fails.

The tests listen on 127.0.0.1.

Reference: https://gobyexample.com/http-server
