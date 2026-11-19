# Troubleshooting

Start with:

```sh
golings doctor
```

It checks the Go toolchain, the race detector, the temp directory, the curriculum and your
workspace, and says what's wrong.

## Go is missing or too old (exit code 4)

golings uses the `go` command on your `PATH` to build and test your code, and it needs Go 1.24
or newer. Install Go from <https://go.dev/dl/>, open a new terminal, and check with
`go version`. golings never downloads a toolchain itself; it runs with `GOTOOLCHAIN=local`.

## "the race detector is not available"

A few concurrency exercises also run the race detector. It needs cgo on some platforms, and
isn't supported on every OS and architecture. When it isn't available, that check shows as
**unsupported**. It is skipped, never counted as passed, and the exercise's other checks
still run. To enable it, install a C compiler (gcc or clang) and make sure `go env CGO_ENABLED`
prints `1`.

## "the checks could not listen on the loopback network"

The networking exercises start servers on `127.0.0.1`. A firewall, VPN client, sandbox or
endpoint-security tool can block that. golings reports it as an environment problem, not a
failed exercise. Allow local connections for Go test binaries (they run from your temp
directory) and try again.

## The first check is slow, or hits the time limit

The first check of each exercise compiles it, and a cold build cache can take a while,
especially on Windows with antivirus scanning new executables. Compilation gets extra time
beyond the exercise's limit, so later runs are quick. If a slow machine still hits the limit,
raise it with `--timeout`, for example `golings verify --timeout 60s`.

A real time-limit failure usually means an infinite loop, a deadlock or a channel operation
that never completes. The message says so.

## "not inside a golings workspace" (exit code 3)

Run golings from inside your workspace (any subdirectory works), or point at it with
`--workspace path/to/golings-workspace`. Create one with `golings init`.

## Your editor shows errors in exercises

That's expected: most starters don't compile until you fix them. The workspace has its own
`go.mod` so gopls understands each exercise folder. The canonical tests aren't in the
workspace, so your editor can't run them; use `golings` instead.

## Progress looks wrong, or state.json is corrupt

Progress lives in `.golings/state.json` and is written atomically. If it was damaged anyway
(by hand editing, or a full disk), golings refuses to guess and tells you. To set the
damaged file aside and start fresh:

```sh
golings doctor --recover-state
```

Your exercise files aren't affected. Run `golings verify --all` to mark everything that
passes as complete again.

## An exercise shows "modified" or "recheck"

- **modified**: you passed it, then edited the file. Verify it again.
- **recheck**: a newer golings changed the exercise's checks. Your code is unchanged, but it
  hasn't been verified against the new checks. `golings next` takes you there.

## After updating golings

New golings releases can add or change exercises. Bring your workspace up to date with:

```sh
golings upgrade --dry-run   # see what would change
golings upgrade
```

It adds new exercises, moves renamed ones together with your files, and refreshes READMEs and
data files. It **never rewrites your code**, and it backs up your progress before changing
it. A workspace made by a *newer* golings is refused; install the newer version instead.

## Exercises marked unsupported on Windows

`execing_processes_02` and `signals_02` use Unix system calls (`exec` and real signals), so
they're checked on Linux and macOS only. On Windows they're skipped automatically and don't
count toward your total.

## Reporting a bug

Include the output of `golings version` and `golings doctor`, the exercise ID, and what you
expected. Please don't include your solutions to exercises in public issues.
