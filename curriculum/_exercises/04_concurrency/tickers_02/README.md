# tickers_02: A stoppable heartbeat

`Heartbeat` calls `beat` on every tick until `stop` is closed. It doesn't stop.

## Your task

Fix `Heartbeat` so it returns promptly once `stop` is closed, stops its ticker, and never
calls `beat` after it has returned.

Reference: https://gobyexample.com/tickers
