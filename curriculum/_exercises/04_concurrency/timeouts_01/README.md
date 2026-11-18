# timeouts_01: Giving up

`select` with `time.After` puts a limit on how long you wait:

```go
select {
case res := <-ch:
	fmt.Println(res)
case <-time.After(time.Second):
	fmt.Println("timeout")
}
```

## Your task

Implement `WithTimeout(fetch func() string, d time.Duration) (string, error)`:

- If `fetch` returns within `d`, return its result and `nil`.
- Otherwise return `""` and `ErrTimeout` as soon as `d` has passed. Don't wait for `fetch`.

A slow `fetch` must not leak a goroutine: once it does finish, its goroutine has to be able
to exit.

Reference: https://gobyexample.com/timeouts
