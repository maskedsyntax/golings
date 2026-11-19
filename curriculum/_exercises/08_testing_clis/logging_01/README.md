# logging_01: Logger prefixes and flags

`log.New(out, prefix, flags)` creates a logger. The flags choose what goes before each message:

```go
log.LstdFlags   // 2009/01/23 01:23:23
log.Lshortfile  // main.go:23:
log.Lmsgprefix  // put the prefix right before the message, not at the start of the line
```

## Your task

Fix both constructors:

- `NewLogger(w)`: `logger.Print("started")` writes `app: started` and nothing else.
- `NewDebugLogger(w)`: writes `<file>:<line>: debug: <message>`, for example
  `logger_test.go:31: debug: cache miss`.

Reference: https://gobyexample.com/logging
