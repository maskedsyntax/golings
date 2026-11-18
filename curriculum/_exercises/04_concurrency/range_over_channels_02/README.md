# range_over_channels_02: Batching

Writing one row at a time to a database is slow, so programs often **batch** a stream of
items.

## Your task

Implement `Batch(in <-chan string, size int) <-chan []string`:

- Group the items from `in` into slices of `size`, in order.
- When `in` is closed, send the last partial batch (if there is one) and then close the
  output channel.
- Return the output channel immediately and do the work in a goroutine.

```text
in: a b c d e   (then closed),  size 2   →   [a b] [c d] [e]   (then closed)
```

Reference: https://gobyexample.com/range-over-channels
