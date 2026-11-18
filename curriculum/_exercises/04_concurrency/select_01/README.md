# select_01: Whichever is ready

`select` waits on several channel operations and runs the case for whichever is ready
first:

```go
select {
case m := <-a:
	fmt.Println("from a:", m)
case m := <-b:
	fmt.Println("from b:", m)
}
```

`Receive` takes `n` messages from two channels. It alternates strictly: a, b, a, b… When `b`
has several messages before `a` has any, it gets stuck.

## Your task

Rewrite `Receive` with `select` so it takes each message from whichever channel is ready.

Reference: https://gobyexample.com/select
