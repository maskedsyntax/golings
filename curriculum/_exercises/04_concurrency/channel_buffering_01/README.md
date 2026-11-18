# channel_buffering_01: A buffered queue

A **buffered channel** holds up to `n` values without a receiver:

```go
ch := make(chan string, 2)
ch <- "a" // doesn't block
ch <- "b" // doesn't block
ch <- "c" // would block: the buffer is full
```

`Queue` should return a channel that already contains every message, so the caller can
receive them later. It sends from the calling goroutine, with no helper goroutine. Right now it
blocks on the first send.

## Your task

Fix `Queue` by giving the channel the right buffer size. Don't add a goroutine.

Reference: https://gobyexample.com/channel-buffering
