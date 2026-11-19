# logging_02: Structured logs

`log/slog` writes *structured* logs: key/value pairs that machines can parse.

```go
logger.Info("started", "port", 8080)
// {"time":"...","level":"INFO","msg":"started","port":8080}
```

Timestamps make output impossible to compare exactly in tests, so this logger leaves them
out.

## Your task

Implement `NewJSONLogger(w io.Writer, level slog.Level) *slog.Logger`. It:

- writes JSON lines to `w`,
- drops the `time` field,
- ignores records below `level`,
- adds `"service":"api"` to every record.

```go
NewJSONLogger(w, slog.LevelInfo).Info("started", "port", 8080)
// {"level":"INFO","msg":"started","service":"api","port":8080}
```

Reference: https://gobyexample.com/logging
