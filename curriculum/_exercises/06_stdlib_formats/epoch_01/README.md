# epoch_01: Milliseconds since 1970

Many APIs send times as a number: seconds (or milliseconds) since 1970-01-01 00:00 UTC, the
*Unix epoch*.

```go
t.Unix()           // seconds
t.UnixMilli()      // milliseconds
time.Unix(sec, 0)  // and back
```

The functions here lose the milliseconds.

## Your task

Fix `ToUnixMillis` and `FromUnixMillis` so a time survives the round trip with millisecond
precision, including times before 1970. `FromUnixMillis` returns times in UTC.

Reference: https://gobyexample.com/epoch
