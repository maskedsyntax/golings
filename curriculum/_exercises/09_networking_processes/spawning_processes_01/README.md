# spawning_processes_01: Running a command

`os/exec` runs other programs:

```go
cmd := exec.CommandContext(ctx, "git", "status")
var out bytes.Buffer
cmd.Stdout = &out
err := cmd.Run()
```

`Run` returns an error both when the program can't start **and** when it exits with a
non-zero status. Those are very different situations.

## Your task

Implement `Run(ctx, name, args...) (stdout string, code int, err error)`:

- On a normal exit, return its standard output and exit status (0 or otherwise) with a nil
  error.
- If the program can't be started, return an error.
- If `ctx` ends first, the process is killed. Return an error wrapping `ctx.Err()`.

The tests use their own test binary as the program, so this works on every OS.

Reference: https://gobyexample.com/spawning-processes
