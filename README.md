# golings

Learn Go by fixing small, broken Go programs until their checks pass, in the spirit of
[Rustlings](https://github.com/rust-lang/rustlings) and [Ziglings](https://ziglings.org/),
following the topics of [Go by Example](https://gobyexample.com/).

Everything runs locally with the Go toolchain you already have: no accounts, no network,
no other runtimes.

> **Status:** early development (Milestone A of [spec.md](spec.md)). The runner works end to end
> with 5 sample exercises. Watch mode, hints, reset and the full 220-exercise curriculum are
> still to come.

## Requirements

- Go 1.24 or newer on your `PATH`.

## Quick start

From a checkout of this repository:

```sh
go install ./cmd/golings   # installs into $(go env GOPATH)/bin, which should be on your PATH
golings init               # creates ./golings-workspace
cd golings-workspace
golings                    # check the current exercise
```

## Commands

| Command | What it does |
|---|---|
| `golings` | Check the current exercise (the next unfinished one) |
| `golings init [path]` | Create a workspace; refuses to touch non-empty directories |
| `golings list` | Exercises and their status (`--module`, `--topic`, `--difficulty`, `--status`, `--json`) |
| `golings verify [id]` / `--all` | Check one exercise, or all of them |
| `golings run <id>` | Check one exercise and show its program output |
| `golings next` | Select the next unfinished exercise |
| `golings progress` | Completion statistics (`--json`) |
| `golings doctor` | Diagnose the toolchain and workspace (`--recover-state` repairs a corrupt progress file) |
| `golings version` | CLI and curriculum versions |

Exercise IDs can be abbreviated to any unambiguous prefix (`golings verify vari`) or given
as their number (`golings verify 3`).

Exit codes: `0` ok, `1` exercise not passing, `2` usage error, `3` workspace problem,
`4` Go toolchain problem, `5` infrastructure error or time limit, `130` interrupted.

## How checking works

Your workspace only ever contains your editable files. To check an exercise, golings copies
your files and the canonical tests (embedded in the binary) into a temporary directory, then
runs `go test` / `go build` there with a clean environment: `GOPROXY=off`, no `go.work`, no
toolchain downloads, and a time limit that kills the whole process tree. Progress lives in
`.golings/state.json` and is written atomically.

Exercises run with your user's privileges and are not sandboxed. Only run exercises you trust.

## Development

```sh
go test ./...                                 # includes the curriculum contract (use -short to skip slow tests)
go run ./cmd/golings-dev validate-curriculum  # metadata, assets, coverage
go run ./cmd/golings-dev check --all          # starters fail, solutions pass, mutants fail
go run ./cmd/golings-dev coverage --write     # regenerate docs/curriculum-coverage.md
```

See [docs/architecture.md](docs/architecture.md) for the code layout and
[spec.md](spec.md) for the full specification.

## Credits

The exercises, tests and solutions are original. Topic selection follows
[Go by Example](https://gobyexample.com/) by Mark McGranaghan and Eli Bendersky. The learning
model comes from Rustlings and Ziglings.
