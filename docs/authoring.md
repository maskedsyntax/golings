# Authoring exercises

This guide is for adding or changing exercises. Read [architecture.md](architecture.md)
first if you want to know how checking works under the hood.

## Quick start

```sh
go run ./cmd/golings-dev new maps          # scaffolds maps_05 and adds it to the manifest
$EDITOR curriculum/_exercises/01_fundamentals/maps_05/
go run ./cmd/golings-dev check maps_05     # starter fails, solution passes, mutants fail
go run ./cmd/golings-dev coverage --write  # refresh docs/curriculum-coverage.md
go run ./cmd/golings-dev validate-curriculum
```

`new` creates a placeholder that already satisfies the contract, so you can change one piece
at a time and re-run `check` after each change.

## Layout

Each exercise lives in `curriculum/_exercises/<module>/<id>/`:

| path | purpose |
|---|---|
| `exercise.json` | metadata (below) |
| `README.md` | the learner's instructions; copied into the workspace |
| `starter/` | the files the learner edits, exactly as listed in `starter_files` |
| `solution/` | a reference solution with the same file names |
| `checks/` | canonical `_test.go` files (and `expected_stdout.txt` for output checks); never shown to learners |
| `fixtures/` | optional data files, staged next to the code at their relative paths |
| `mutants/<name>/` | plausible wrong answers (same file names as the starter) that must fail the checks |

The ID is `<topic>_NN`, and the order comes only from `curriculum/manifest.json`. An exercise's
number is its position there, so don't put ordinals anywhere else.

## exercise.json

| field | values |
|---|---|
| `type` | `compile_fix`, `implement`, `debug`, `output`, `integration`, `concurrency` |
| `difficulty` | `beginner`, `intermediate`, `advanced` |
| `starter_files` | plain file names, `.go` only; `_test.go` is allowed when the learner writes the tests |
| `fixture_files` | relative paths under `fixtures/`; must not collide with starter files |
| `checks` | any of `build`, `unit`, `race`, `vet`, `output`, `ast`, run in order; the first non-pass stops |
| `ast_rules` | required with the `ast` check: `uses_defer:F`, `uses_range:F`, `uses_switch:F`, `recursive:F`, `declares_const:Name`, `calls:pkg.Func`, `no_call:pkg.Func`, `chan_params_directional:F` (`F` may be `Type.Method`) |
| `platforms` | empty for all, or a subset of `darwin`, `linux`, `windows` |
| `starter_expect` | `check_failure` or `compile_error`: how the untouched starter must fail |
| `hints` | 1 to 3, from a nudge to nearly the answer |
| `timeout_seconds` | 1 to 600; the test binary gets this, and compilation gets extra time on top |

Metadata strings are never turned into commands: each check kind maps to a fixed invocation.

## The contract

`golings-dev check` (and `go test ./internal/validate`) enforces, for every exercise:

1. The **starter** fails, in the way `starter_expect` says. A `go vet` failure reported by
   `go test` counts as `compile_error`.
2. The **solution** passes every check.
3. Every **mutant** fails. A mutant should be a mistake a learner might really make, not
   random breakage. `debug` and `concurrency` exercises should have at least one.

A check that the platform can't run (the race detector without cgo, or an exercise whose
`platforms` exclude this OS) is skipped, never counted as a pass.

## Writing checks

- **Fail fast, never hang.** Contract runs reject timeouts for starters and mutants. Guard
  anything that can block with a deadline helper that fails the test, for example:

  ```go
  func within(t *testing.T, what string, d time.Duration, f func()) {
  	t.Helper()
  	done := make(chan struct{})
  	go func() { defer close(done); f() }()
  	select {
  	case <-done:
  	case <-time.After(d):
  		t.Fatalf("%s did not finish within %v", what, d)
  	}
  }
  ```

  A value from `time.After` is delivered once. To share one deadline between several waits,
  close a channel with `time.AfterFunc` instead.
- **Prefer determinism.** Use barriers and channels rather than sleeps, seeded
  `math/rand/v2` generators, and injected clocks where you can. Where real time is
  unavoidable, use generous bounds, and make sure the mutants fail by a wide margin, not by a
  few milliseconds.
- **Leaks.** Poll `runtime.NumGoroutine` for a while before declaring a leak; goroutines take
  time to exit.
- **Files** go in `t.TempDir()`. **Networking** uses `127.0.0.1:0` or `httptest`. If the OS
  refuses a loopback socket, golings reports an environment problem instead of a failure.
- **Processes, exits and signals:** re-run the test binary as the child program. Check an
  environment variable in `TestMain` and act as the child when it's set (see `exit_01` and
  `spawning_processes_01`). This works on every OS. Real Unix signals and `syscall.Exec` need
  `"platforms": ["darwin", "linux"]`.
- **Messages** should tell the learner what's wrong in their terms:
  `Slugify("Go World") = "Go-World", want "go-world"`, not `assertion failed`.

### When the learner writes the tests

In the testing topic the learner writes `_test.go` starters. The code under test is a fixture
that declares the function as a variable (`var Reverse = func(...)`), and a canonical
`grade_test.go`:

- swaps in broken versions under `TestMain` when an environment variable names one, and
- re-runs the learner's tests in a child process once per broken version, requiring each run
  to fail.

See `testing_and_benchmarking_01` to `_03`.

## Fixtures

Fixtures are staged next to the code for checks and copied into the workspace for the
learner to read. Checks always use the embedded originals, so a learner's edits to the copies
never change a result. `//go:embed` exercises use fixtures for the embedded files.

## Changing an existing exercise

Learners keep their work across curriculum versions with `golings upgrade`, so:

- **Bump `curriculum_version`** in `manifest.json` for any release that changes exercises.
- **Changing checks or fixtures** changes the exercise's checks hash. Learners who passed it
  before see it as `recheck` and are asked to verify it again. Changing hints, titles or the
  README doesn't trigger that.
- **Renaming or removing** an exercise needs an entry in `curriculum/migrations.json`:

  ```json
  {"to_version": "0.2.0", "renamed": {"maps_04": "maps_05"}, "removed": ["values_02"]}
  ```

  `upgrade` moves the learner's directory to the new ID (with their files), carries their
  progress over, and leaves removed exercises' files untouched. Renames are followed across
  versions, and the validator rejects renames to unknown IDs or of IDs still in use.
- Starters are never rewritten in a learner's workspace. A fix to a starter reaches only new
  workspaces and `golings reset`.

## Topics

`curriculum/topics.json` mirrors the Go by Example index. Before a release, run

```sh
go run ./cmd/golings-dev compare-topics
```

It is the only command that uses the network. It fetches the index read-only and prints the
differences without writing anything. Once `release_complete` is true in the manifest, every
topic must keep at least two exercises.

## Style

- Exercises, tests and solutions are original. Follow Go by Example's choice of topics, not its
  code.
- Keep the starter's header comment short: what to do, a pointer to `README.md`, and the
  reference URL. The README teaches the idea with a small example and states the task
  precisely.
- Reference solutions are idiomatic and may carry a short comment on the key line.
- Canonical files are gofmt'd. Starters are exempt, since some can't parse on purpose.
