# tcp_server_02: Idle timeouts

Network servers must not wait forever for a client that has gone quiet. Deadlines make
blocking network calls fail after a point in time:

```go
conn.SetReadDeadline(time.Now().Add(30 * time.Second))
n, err := conn.Read(buf) // fails with a timeout error after 30s of silence
```

This server answers `PING` with `PONG`. It should hang up on any client that stays silent for
longer than `Idle`, but active clients must be able to talk for as long as they like.

## Your task

Fix `handle` so it closes a connection after `Idle` with no data from the client, and only then.

Reference: https://gobyexample.com/tcp-server
