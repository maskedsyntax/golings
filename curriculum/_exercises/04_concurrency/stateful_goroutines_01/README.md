# stateful_goroutines_01: One goroutine owns the state

Instead of locking shared state, you can give it to **one goroutine** and have everyone else
send it requests over channels. This is Go's "share memory by communicating":

```go
type readOp struct {
	key  string
	resp chan int
}
```

The owner loops over a `select`, answers reads, applies writes, and is the only code that
ever touches the map.

## Your task

Implement `NewStore`, `Get` and `Set` using the request types provided. Don't use a mutex.

Reference: https://gobyexample.com/stateful-goroutines
