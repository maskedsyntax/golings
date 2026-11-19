# tcp_server_01: Serving many clients

A TCP server accepts connections in a loop and handles each one:

```go
for {
	conn, err := ln.Accept()
	...
	go handle(conn)
}
```

This server uppercases each line a client sends. While one client is connected, nobody else
gets an answer. And closing the listener, which is how the program stops it, makes `Serve`
return an error.

## Your task

Fix `Serve`: serve clients concurrently, and return `nil` once the listener is closed.

Reference: https://gobyexample.com/tcp-server
