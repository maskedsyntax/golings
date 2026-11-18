# json_03: Streaming JSON lines

Logs are often written as *JSON lines*: one JSON object per line.

```text
{"level":"info","msg":"started"}
{"level":"warn","msg":"disk 90% full"}
{"msg":"no level"}
```

A `json.Decoder` reads values one at a time from an `io.Reader`, so you never need the whole
file in memory.

## Your task

Implement `CountByLevel(r io.Reader) (map[string]int, error)`. It counts entries by their
`level` field, counting entries without a level as `"unknown"`. The example above gives
`{"info": 1, "warn": 1, "unknown": 1}`.

Malformed input is an error. Don't skip it. Blank lines are fine, and empty input gives an empty
(non-nil) map.

Reference: https://gobyexample.com/json
