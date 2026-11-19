# Contributing

Thanks for helping. golings has two halves:

- **the CLI** (`cmd/golings`, `internal/...`): see [architecture.md](architecture.md);
- **the curriculum** (`curriculum/`): see [authoring.md](authoring.md).

## Setup

You need Go 1.24 or newer and nothing else. Development works offline once the module cache
has `golang.org/x/sys` and `golang.org/x/term`.

```sh
go build ./...
go test ./...          # add -short to skip the slow curriculum contract and CLI journeys
```

## Before you open a pull request

CI runs these on Linux, macOS and Windows with Go 1.24 and the latest stable release. Running
them locally first saves a round trip:

```sh
gofmt -l $(git ls-files '*.go' | grep -v /starter/)   # should print nothing
go vet ./...
go test -race ./...                                    # plain `go test ./...` on Windows
go run ./cmd/golings-dev validate-curriculum
go run ./cmd/golings-dev check --all                   # or just the exercises you touched
```

If you changed exercises, also run `go run ./cmd/golings-dev coverage --write` and commit the
regenerated `docs/curriculum-coverage.md`.

## Guidelines

- **Dependencies:** the standard library plus `golang.org/x/sys` and `golang.org/x/term`.
  Adding another needs a strong reason.
- **No network** in golings itself. The only networked code is the maintainer command
  `golings-dev compare-topics`.
- **Never lose learner work.** Nothing may overwrite a learner's file without asking, except
  `reset --force`, which makes a backup unless `--no-backup` is given. Progress writes are
  atomic.
- **Cross-platform:** use `filepath`, not string concatenation, for paths. Don't assume `/`
  separators or `\n` line endings in tool output. Features that can't work on an OS must
  report "unsupported", never a pass.
- **Errors for humans:** messages say what went wrong and what to do next. Exit codes are
  documented in the README and are part of the interface.
- **Tests:** new behavior comes with tests. Timing-sensitive tests must tolerate slow CI
  machines (see the notes in authoring.md).
- Keep commits focused, with a message that explains why.

## Releasing

1. Make sure CI is green on `master`, and run `go run ./cmd/golings-dev compare-topics`.
2. If exercises changed since the last release, bump `curriculum_version` in
   `curriculum/manifest.json` and add any renames or removals to `curriculum/migrations.json`.
3. Tag and push: `git tag v1.2.3 && git push origin v1.2.3`.
4. The release workflow runs the tests and the curriculum contract, then GoReleaser builds
   archives for macOS, Linux and Windows (amd64 and arm64) with `checksums.txt`, as a
   **draft** release. Review the draft and publish it.

To check the release setup locally: `goreleaser check`, then
`goreleaser release --snapshot --clean` (artifacts land in `dist/`).

## Licensing

golings is MIT licensed (see `LICENSE`). By contributing you agree to license your work under
the same terms. Exercises must be original: Go by Example (CC BY 3.0) is credited as the
inspiration for the topic list, but its code and text must not be copied.
