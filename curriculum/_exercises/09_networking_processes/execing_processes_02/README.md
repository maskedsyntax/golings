# execing_processes_02: Replacing the process

Unlike spawning, `syscall.Exec` **replaces** the running program with another one: same
process ID, no parent waiting for a child. It only returns if it fails.

```go
path, err := exec.LookPath("ls")
err = syscall.Exec(path, []string{"ls", "-a"}, os.Environ())
```

`Exec(tool, args...)` looks up `tool` and execs it with `args`, but the new program receives
the wrong arguments.

## Your task

Fix `Exec` so the new program sees `os.Args[1:] == args`.

This exercise uses the Unix `exec` system call, so it is checked on Linux and macOS only.

Reference: https://gobyexample.com/execing-processes
