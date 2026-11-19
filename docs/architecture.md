# Architecture

```
cmd/golings        CLI entry point (thin wrapper around internal/cli)
cmd/golings-dev    maintainer tool: validate-curriculum, check, coverage
curriculum/        embedded curriculum (manifest.json, topics.json, _exercises/)
internal/
  cli/             argument parsing, terminal output, exit codes
  app/             use cases: evaluate + record, current/next selection, summaries
  catalog/         strict metadata loading and validation, ID matching
  workspace/       init, discovery, safe reads (os.Root), fingerprints
  runner/          staging, check execution, process-tree kill, go test -json parsing
  progress/        state.json with lockfile + atomic rename, corruption recovery
  render/          human (color/no-color) and JSON output
  diagnostics/     file:line parsing, staging path -> workspace path rewriting
  toolchain/       go detection, version gate (>= 1.24), race support
  validate/        curriculum integrity, coverage report, contract checks
  exitcode/        exit code constants
```

## Curriculum layout

`curriculum/manifest.json` is the single source of truth for which exercises exist and in
what order; an exercise's ordinal is its position in the manifest. Each exercise lives at
`curriculum/_exercises/<module>/<id>/`:

```
exercise.json   metadata (no ordinal, no URL: both are derived)
README.md       learner-facing instructions, copied into the workspace
starter/        files the learner edits, copied into the workspace (may include _test.go
                files when the learner writes the tests, as in the testing topic)
checks/         canonical *_test.go files and expected_stdout.txt, never copied to the workspace
solution/       reference solution (same file names as starter/)
mutants/<name>/ plausible wrong answers that must fail the checks
fixtures/       optional data files staged next to the code; reference copies are also
                written to the workspace, but checks always use the embedded originals
```

The directory is named `_exercises` so the go tool ignores it: intentionally broken
starters never affect `go build ./...` or `go vet ./...` for golings itself. It is embedded
with `//go:embed all:_exercises`.

## Evaluation

1. Read the learner's files through `os.Root`, which rejects symlinks and paths escaping the exercise directory.
2. Create a temporary staging directory holding a `go.mod` (`go 1.24`), the learner files, any fixtures, and `checks/*.go`.
3. Run each declared check kind from a fixed allowlist (`build`, `unit`, `race`, `vet`, `output`). Each kind maps to a hard-coded `go` invocation, so metadata never becomes a command.
4. The environment is built from scratch: `GOPROXY=off GOWORK=off GOTOOLCHAIN=local GOENV=off GOFLAGS=`, with the user's `GOCACHE` reused for speed.
5. The whole process tree is killed on timeout or cancel: a process group on Unix, a Job Object on Windows.
6. Results are typed (`pass`, `fail`, `timeout`, `runner_error`, `unsupported`) with a failure category: compile, test, output, vet, race, timeout, or runner.
7. Progress is recorded only after evaluation completes. Runner errors and unsupported platforms are never counted as attempts or completions.
