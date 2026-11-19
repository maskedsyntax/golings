# execing_processes_01: Preparing an exec

`syscall.Exec` replaces the current process with another program. It needs three things,
each with a common pitfall:

- the **path** to the program: `exec` doesn't search `$PATH`, so look it up first;
- **argv**, which includes the program's own name as `argv[0]`;
- the full **environment** as `KEY=value` strings.

This exercise prepares those three on any OS. The next one performs the exec.

## Your task

Implement `Prepare(tool, args, extra, environ, lookPath) (path string, argv, env []string, err error)`:

- `path`: `lookPath(tool)`. Return an error that wraps its error.
- `argv`: `tool` followed by `args`.
- `env`: `environ`, with any variable named in `extra` replaced **where it is**. Variables in
  `extra` that aren't already in `environ` go at the end, sorted by name.

Reference: https://gobyexample.com/execing-processes
