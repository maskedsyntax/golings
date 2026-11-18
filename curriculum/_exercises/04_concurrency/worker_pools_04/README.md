# worker_pools_04: Fan-in

When several workers each produce a stream of results, a **fan-in** merges those streams into
one.

## Your task

Implement `Merge(chans ...<-chan int) <-chan int`:

- Every value from every input appears on the output exactly once. Order across inputs
  doesn't matter.
- The output is closed once **all** inputs are closed. With no inputs it closes right away.
- `Merge` returns immediately; the copying happens in goroutines.

Reference: https://gobyexample.com/worker-pools
