# golings

Learn Go by fixing small, broken Go programs until their checks pass, in the spirit of
[Rustlings](https://github.com/rust-lang/rustlings) and [Ziglings](https://ziglings.org/),
following the topics of [Go by Example](https://gobyexample.com/).

Everything runs locally with the Go toolchain you already have: no accounts, no network,
no other runtimes.

> **Status:** feature-complete: 220 exercises covering all 85 Go by Example topics, from hello
> world to signals and process exit codes, with watch mode, hints, solutions and workspace
> upgrades. Heading for 1.0.

## Requirements

- Go 1.24 or newer on your `PATH`.

## Install

Either download an archive for your platform from the
[releases page](https://github.com/maskedsyntax/golings/releases) and put `golings` on your
`PATH`, or build it with Go:

```sh
go install github.com/maskedsyntax/golings/cmd/golings@latest
```

`go install` puts the binary in `$(go env GOPATH)/bin`, which should be on your `PATH`. From a
checkout of this repository, `go install ./cmd/golings` does the same.

## Quick start

```sh
golings init               # creates ./golings-workspace
cd golings-workspace
golings                    # watch and check the current exercise
```

Edit the file it names, save, and golings re-checks it. When it passes, press `n` for the
next exercise.

## Commands

| Command | What it does |
|---|---|
| `golings` / `golings watch` | Watch the current exercise and re-check it on every save. Keys: `h` hint, `n` next, `p` previous, `l` list, `r` re-run, `q` quit (`--auto-advance` moves on after a pass) |
| `golings init [path]` | Create a workspace; refuses to touch non-empty directories |
| `golings list` | Exercises and their status (`--module`, `--topic`, `--difficulty`, `--status`, `--json`) |
| `golings verify [id]` / `--all` | Check one exercise, or all of them |
| `golings run <id>` | Check one exercise and show its program output |
| `golings next` | Select the next unfinished exercise |
| `golings hint [id]` | Reveal the next hint (`--level N` for a specific one) |
| `golings reset <id>` | Restore the starter files; asks first and backs up your version to `.golings/backups/` |
| `golings solution <id>` | Show the reference solution (asks first, or `--show`; `--out DIR` writes files) |
| `golings progress` | Completion statistics (`--json`) |
| `golings upgrade` | After updating golings: add new exercises and carry your work over (`--dry-run` first; your code is never rewritten) |
| `golings doctor` | Diagnose the toolchain and workspace (`--recover-state` repairs a corrupt progress file) |
| `golings version` | CLI and curriculum versions |

Exercises can be named by full ID (`golings verify slices_02`), by any unambiguous prefix, or
by number (`golings verify 3`). An ambiguous prefix lists the matches instead of guessing.

Exit codes: `0` ok, `1` exercise not passing, `2` usage error, `3` workspace problem,
`4` Go toolchain problem, `5` infrastructure error or time limit, `130` interrupted.

## How checking works

Your workspace only ever contains your editable files. To check an exercise, golings copies
your files and the canonical tests (embedded in the binary) into a temporary directory, then
runs `go test` / `go build` there with a clean environment: `GOPROXY=off`, no `go.work`, no
toolchain downloads, and a time limit that kills the whole process tree. Progress lives in
`.golings/state.json` and is written atomically.

Some exercises come with data files (fixtures). They are copied into the exercise folder so
you can read them, but checks always use the original copies.

Two exercises (`execing_processes_02` and `signals_02`) use Unix-only system calls and are
checked on Linux and macOS only. On Windows they show as unsupported and don't count against
your progress. Exercises in the networking module listen on `127.0.0.1`; if a firewall blocks
that, golings says so instead of failing the exercise.

Exercises run with your user's privileges and are not sandboxed. Only run exercises you trust.

## Development

```sh
go test ./...                                 # includes the curriculum contract (use -short to skip slow tests)
go run ./cmd/golings-dev validate-curriculum  # metadata, assets, coverage
go run ./cmd/golings-dev check --all          # starters fail, solutions pass, mutants fail
go run ./cmd/golings-dev coverage --write     # regenerate docs/curriculum-coverage.md
```

More documentation:

- [docs/troubleshooting.md](docs/troubleshooting.md): toolchain, firewall, time limits, progress
- [docs/authoring.md](docs/authoring.md): writing and changing exercises
- [docs/contributing.md](docs/contributing.md): development workflow and releases
- [docs/architecture.md](docs/architecture.md): how the code fits together
- [spec.md](spec.md): the full specification

## Credits

The exercises, tests and solutions are original. Topic selection follows
[Go by Example](https://gobyexample.com/) by Mark McGranaghan and Eli Bendersky
(CC BY 3.0). The learning model comes from [Rustlings](https://github.com/rust-lang/rustlings)
and [Ziglings](https://ziglings.org/).

golings is MIT licensed. Release binaries include the Go runtime and `golang.org/x/sys` and
`golang.org/x/term` (BSD 3-Clause); see [THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md).
