# worker_pools_01: A fixed number of workers

`MapPool` applies `f` to every item concurrently. It starts one goroutine per item, which is
fine for ten items and a problem for ten million, or when `f` talks to a server that only
accepts a few connections at once.

A **worker pool** runs a fixed number of goroutines that take jobs from a shared channel:

```go
for w := 1; w <= 3; w++ {
	go worker(w, jobs, results)
}
```

## Your task

Rewrite `MapPool` so **at most `workers`** calls to `f` run at the same time, while still using
all of them. `results[i]` must be `f(items[i])`.

Reference: https://gobyexample.com/worker-pools
