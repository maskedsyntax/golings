# closing_channels_01: No more values

**Closing** a channel tells receivers that no more values will be sent:

```go
close(jobs)
v, ok := <-jobs // ok is false once jobs is closed and drained
```

`Generate` sends every number, but the consumer in the tests waits for a "no more values"
signal that never comes.

## Your task

Fix `Generate` so the channel is closed after the last value has been sent.

Reference: https://gobyexample.com/closing-channels
