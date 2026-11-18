# channel_synchronization_01: A done signal

A goroutine can tell others "I'm finished" through a channel. Here the signal is a
**close**: receiving from a closed channel never blocks, so every waiter wakes up.

```go
done := Start(work)
<-done // waits until work has finished
```

Right now `<-done` returns before `work` has even started.

## Your task

Fix `Start` so its channel is closed only after `task` has returned.

Reference: https://gobyexample.com/channel-synchronization
