# recover_01: Turning a panic into an error

`recover` stops a panic and lets the program continue. It only works inside a **deferred**
function:

```go
defer func() {
	if r := recover(); r != nil {
		fmt.Println("Recovered. Error:\n", r)
	}
}()
```

## Your task

Implement `Safely(f func()) error`. It calls `f`. If `f` panics, `Safely` returns an error
whose message is `recovered: <panic value>`; otherwise it returns `nil`. `Safely` itself must
never panic.

Reference: https://gobyexample.com/recover
