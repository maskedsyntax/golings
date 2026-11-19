# signals_01: Reacting to signals

Servers react to Unix signals: `SIGHUP` conventionally means "reload your configuration", and
`SIGINT` (Ctrl+C) or `SIGTERM` means "shut down". In Go they arrive on a channel registered with
`signal.Notify`.

Here the logic takes the channel as a parameter, so tests can send signals without really
signalling the process, and it works on every OS.

## Your task

Fix `Supervise(sigs, reload, shutdown) os.Signal`:

- `syscall.SIGHUP` calls `reload`, as many times as it arrives.
- `os.Interrupt` or `syscall.SIGTERM` calls `shutdown` once and returns that signal.
- Other signals are ignored.
- If `sigs` is closed, return `nil` without calling `shutdown`.

Reference: https://gobyexample.com/signals
