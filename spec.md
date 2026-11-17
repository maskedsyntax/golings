# Golings — Product & Technical Specification

> **Status:** Proposed implementation specification  
> **Version:** 1.0 draft  
> **Date:** 2026-11-17  
> **Implementation language:** Go, end to end  
> **Working project name:** `golings` (renameable before release)  
> **Primary references:** [Go by Example](https://gobyexample.com/), [Rustlings](https://github.com/rust-lang/rustlings), [Ziglings](https://ziglings.org/)

## 1. Executive summary

Golings is a local-first, terminal-native, exercise-driven way to learn and practice Go. Inspired by Rustlings and Ziglings, it presents small, intentionally incomplete or broken Go programs that learners repair until automated checks pass. It follows **every topic listed on Go by Example as of 2026-11-17 (85 topics)**, but the exercises, explanations, tests, and solutions are original. The learning path progresses from syntax and basic collections through interfaces, generics, concurrency, the standard library, networking, testing, and processes.

The application, CLI, watcher, exercise runner, content validators, and tests are implemented in Go. Go is the only runtime required; the user does not need Node, Python, Docker, a database, an account, or a network connection. Exercises must run on macOS, Linux, and Windows. A local folder contains the exercises and the learner's edits.

**Release target:** 85/85 topics represented, **220 core exercises** (planning target, not an arbitrary completion gate), plus optional module capstones. A smaller first release may ship before the entire curriculum, but the final curriculum release must meet the coverage gate.

### 1.1 Product principles

1. **Practice before exposition.** Every concept is learned by editing Go code and observing results.
2. **Fast feedback.** Compile/test feedback is automatic and easy to interpret.
3. **Real Go.** Use `go test`, compiler diagnostics, the race detector, `go vet`, and ordinary Go tooling instead of inventing a language-specific evaluator.
4. **Small and progressively harder.** Most exercises should take 3–15 minutes; advanced challenges may take longer.
5. **Offline and editor-agnostic.** Any text editor works; no cloud dependency.
6. **Reliable grading.** Known-good solutions pass; starters fail intentionally; tests are deterministic and isolated.
7. **Respect learner work.** Never silently overwrite user edits; state persists across CLI restarts and upgrades.
8. **Complete, auditable coverage.** Every Go by Example topic has an explicit mapping, exercises, and tests.
9. **Go-only codebase.** The shipped executable and project tooling are written in Go; shell commands may be documented for users, but no external runtime is required.

## 2. Scope

### 2.1 In scope

- A standalone `golings` executable for macOS, Linux, Windows.
- A distributed curriculum of source exercises, hints, public tests, private/extra validation checks where technically feasible, reference solutions, and metadata.
- Watch mode (default interactive path), one-off verification, listing/filtering, next/resume, hints, progress, reset, and solutions.
- Clear compile/test failure reporting, including expected versus actual results.
- Different exercise evaluation modes: compile/fix, unit test, stdout, subprocess/integration, and conceptual behavior via executable assertions.
- Standard library-centric implementation; third-party dependencies only when justified.
- A manifest-driven curriculum, dependency/prerequisite information, and coverage validation.
- Deterministic CI checks for tests and solutions, including platform compatibility and concurrency-specific race checks.
- Accessibility basics: readable output, terminal without color, non-interactive use, and meaningful exit codes.

### 2.2 Out of scope for v1

- Web app, accounts, sign-in, cloud sync, hosted leaderboard, telemetry.
- AI hints, code generation, automated code review by a remote model.
- A custom editor or language server.
- A perfect anti-cheating system: the learner owns their local machine and files.
- Automatic daily curriculum updates from Go by Example without human review.
- Running learner-written code in a security-grade sandbox. **Local exercises are trusted educational content; arbitrary malicious code is not safely isolated.**
- Supporting every historical Go release.

### 2.3 Target audience

- Go beginners with basic programming familiarity.
- Programmers coming from Java, JavaScript, Python, Rust, or other languages.
- Experienced Go developers refreshing concurrency, tools, and standard-library topics.

### 2.4 Definition of success

- A beginner can install, start, edit, verify, receive a useful error, request a hint, resume later, and finish an exercise without reading the project's internals.
- All 85 source topics are accounted for by the manifest.
- Every released starter deliberately fails its expected initial check; every solution passes all designated checks.
- Watch mode feels responsive and never loses user work.
- All first-party tests pass on macOS, Linux, and Windows.

## 3. User experience

### 3.1 Initial setup

Preferred onboarding:

```sh
# Install the golings binary by release archive or go install (once published).
golings init
cd golings-workspace
golings
```

`golings init [path]` creates a **versioned workspace** from exercise assets embedded in the executable. The source distribution also supports `go run ./cmd/golings ...` from the repository. It must not write into the user's current directory without confirmation if that directory already contains files.

Recommended layout for a learner workspace:

```text
golings-workspace/
  .golings/
    state.json                   # local progress (untrusted, recoverable)
    workspace.json               # format and curriculum version
  exercises/
    01_fundamentals/
      001_hello_world/
        main.go
      002_values_01/
        values.go
      ...
  README.md
```

**Important separation:** shipped assets contain immutable canonical checks and solutions; the workspace contains editable starter files only. The runner reads canonical grading tests from embedded assets and builds a temporary staging directory when evaluating a workspace exercise. Never place writable canonical tests or solutions in the normal learner workspace. Optionally provide public test examples in the workspace as read-only *by convention*, but never rely on them alone for grading.

### 3.2 Primary loop

1. User runs `golings` (equivalent to `golings watch`).
2. The CLI selects the first unfinished exercise unless a previously selected unfinished exercise is recorded.
3. It shows exercise ID, title, concept, path, instructions, and a Go by Example reference URL.
4. User opens the source in any editor.
5. File changes trigger debounced re-evaluation.
6. Compiler errors, failing assertions, or runtime diagnostics are shown succinctly.
7. A passing result is recorded; the CLI offers to advance (configurable auto-advance, default **off** to avoid surprises).
8. The user may ask for hints, skip, navigate, or quit at any time.

Example:

```text
$ golings
GOLINGS  •  learn Go by fixing Go
Exercise 012 / 220   functions_02  •  Functions
Path: exercises/02_functions_types/012_functions_02/functions.go
Source: https://gobyexample.com/functions

Goal: Implement Max(a, b int) int.
Waiting for changes...  [h] hint  [n] next  [l] list  [q] quit

FAIL  TestMaxNegative
  got: 0
 want: -2
PASS  TestMaxEqual

✓ All 5 checks passed
Exercise complete!  Press [n] to continue.
```

The number displayed comes from the actual manifest; the example is illustrative.

### 3.3 CLI contract

| Command | Behavior | Notes |
|---|---|---|
| `golings` / `golings watch` | Watch current/next unfinished exercise | Interactive if terminal; sensible fallback otherwise |
| `golings init [path]` | Materialize new workspace | Refuse nonempty destination unless safe explicit action |
| `golings list` | List exercises with status | `--module`, `--topic`, `--difficulty`, `--status`, `--json` |
| `golings run <id>` | Run one exercise | One-shot, no file watching |
| `golings verify [id]` | Run all required checks for one exercise | Defaults to current exercise |
| `golings verify --all` | Verify every exercise in workspace | Exit nonzero on failures; no surprise modifications |
| `golings next` | Print/select next unfinished exercise | Ignores intentionally skipped entries unless configured |
| `golings hint [id]` | Show next hint | `--level N` optional; persist revealed level |
| `golings progress` | Completion stats | `--json` for scripts |
| `golings reset <id>` | Restore original starter | Requires confirmation; `--force`; optional backup |
| `golings solution <id>` | Display canonical solution | Confirm or require `--show` to discourage accidental spoilers |
| `golings doctor` | Check Go toolchain, workspace, platform | Actionable diagnostics |
| `golings version` | Print CLI/curriculum versions | `--json` optional |
| `golings help` | Usage | Also `--help` |

**Exit codes:** `0` success, `1` exercise verification failure, `2` command/argument error, `3` invalid workspace or configuration, `4` missing/incompatible Go toolchain, `5` infrastructure/time-limit failure. Use these consistently; Ctrl+C exits cleanly (prefer 130 when meaningful in noninteractive workflows).

**Flags:** global `--workspace <path>`, `--no-color`, `--json` where supported, `--verbose`, `--timeout <duration>` for compatible commands. Output in JSON mode must be parseable with no ANSI escapes or interleaved prompts.

**Help behavior:** Unknown IDs show nearby matches; ambiguous prefix matches require disambiguation rather than silently choosing the wrong exercise.

### 3.4 Usability and accessibility

- Distinguish **compiler failure**, **test failure**, **timeout**, **race**, and **runner error**.
- Show paths and useful line numbers; preserve original Go tool diagnostics where relevant.
- Cap displayed output (e.g., last 64 KiB) while retaining bounded diagnostic details; indicate truncation.
- Disable animations and unnecessary terminal redraw in non-TTY mode.
- Honor `NO_COLOR`, `--no-color`, and Windows terminal behavior.
- Do not require vim-like shortcuts: support CLI equivalents for interactive actions.
- Print hints progressively; explain what a particular test expects without spoiling the solution.

## 4. Repository and architecture

```text
golings/
  cmd/golings/main.go
  internal/
    cli/                 # command parsing, TTY interaction
    app/                 # orchestration/use cases
    catalog/             # manifest loading, IDs, ordering, filtering
    workspace/           # initialization, reconciliation, backups
    runner/              # staging, go test/build/run, results
    watcher/             # portable polling/file change detection
    progress/            # state machine and atomic persistence
    render/              # human and JSON output
    diagnostics/         # Go tool output parsing
    toolchain/           # go version/availability
    validate/            # authoring and curriculum CI validation
  curriculum/
    manifest.json        # authoritative exercise metadata
    topics.json          # 85-topic source inventory
    exercises/
      01_fundamentals/
        001_hello_world/
          README.md
          starter/main.go
          checks/main_test.go
          solution/main.go
          exercise.json
      ...
  docs/
    authoring.md
    architecture.md
    contributing.md
    curriculum-coverage.md
  testdata/
  go.mod
  go.sum
  README.md
  LICENSE
  spec.md
```

### 4.1 Key boundaries

- **CLI** parses commands and handles terminal UX; it does not implement evaluation.
- **Catalog** exposes validated, immutable metadata with stable exercise IDs.
- **Workspace** creates and safely restores starter files and maps IDs to user paths.
- **Runner** evaluates exercises with the local Go toolchain and returns typed results.
- **Watcher** identifies relevant file changes and requests runner rechecks.
- **Progress** manages completion status independently of source modification times.
- **Validator** checks curriculum integrity and authoring invariants in CI.

Use standard-library APIs (`os/exec`, `context`, `encoding/json`, `filepath`, `io/fs`, `embed`, `testing`) by default. A polling watcher is sufficient for initial development; an optional filesystem event library can be evaluated later only if profiling demonstrates a problem. An interactive TUI library is optional; start with readable terminal output and basic input handling.

### 4.2 Executable and asset packaging

The built binary embeds a versioned curriculum snapshot (`//go:embed`) containing starters, checks, metadata, and reference solutions. `golings init` copies only editable starters and learner-facing instructions. The executable can read canonical checks/solutions from embedded files and materialize them into temporary staging directories.

**Security caveat:** Embedded tests and solutions can be extracted from a local binary. This is acceptable; the tool is designed for self-directed learning, not high-stakes examinations.

### 4.3 Workspace evaluation strategy

For each exercise:

1. Resolve the exercise by stable ID and validate workspace paths.
2. Read learner-controlled source under that exercise directory.
3. Create a temporary staging folder (outside workspace) with a minimal `go.mod` and any vetted fixture files.
4. Copy learner source files into staging, rejecting traversal and unapproved symlinks.
5. Copy the canonical tests/check harness into staging (not into workspace), avoiding filename collisions.
6. Run the exercise's declared check steps via `exec.CommandContext` with a fresh environment, controlled working directory, and a timeout.
7. Capture bounded stdout/stderr, exit status, duration, and diagnostics.
8. Remove staging resources, including child processes where practical.
9. Commit a progress update only after all required checks pass.

For stdlib-only exercises, the staging module should not rely on network resolution. Avoid calling `go mod tidy` for every verification. Use isolated build cache configuration only where necessary; avoid excessive startup costs. Never run tests against the `golings` production package directly if the package contains tests the learner could overwrite.

Some exercises need cross-file package declarations or embed fixtures. Their `exercise.json` declares the file allowlist and fixture layout. Embedded files may require a temporary directory structure that preserves source-relative assets. Subprocess exercises get a distinct evaluation mode rather than being forced into conventional unit tests.

## 5. Exercise authoring contract

### 5.1 Exercise types

| Type | Learner action | Verification |
|---|---|---|
| `compile_fix` | Fix invalid syntax or type error | `go test`/`go build` compiles, required behavioral checks pass |
| `implement` | Complete TODOs in function or package | Canonical unit tests |
| `debug` | Fix code that compiles but behaves incorrectly | Unit tests and/or behavioral assertions |
| `output` | Produce exact specified observable output | Bounded subprocess/stdout comparison |
| `integration` | Implement file/network/process behavior | Isolated integration test with fixtures and timeout |
| `concurrency` | Implement synchronization/cancellation | Deterministic tests, select race checks, timeout/deadlock checks |

Avoid grading by raw source text or forcing a specific implementation unless the learning objective is explicitly about using a language construct. When construct-specific constraints genuinely matter (e.g., `defer` or a directional channel), use carefully documented static or AST-based validation in Go; do not encourage brittle string matching.

### 5.2 Metadata schema (illustrative, v1)

```json
{
  "schema_version": 1,
  "id": "slices_03",
  "ordinal": 24,
  "module": "01_fundamentals",
  "topic": "slices",
  "title": "Remove duplicate values",
  "type": "implement",
  "difficulty": "beginner",
  "estimated_minutes": 10,
  "go_by_example_url": "https://gobyexample.com/slices",
  "prerequisites": ["slices_01", "maps_01"],
  "starter_files": ["slices.go"],
  "fixture_files": [],
  "checks": ["unit"],
  "timeout_seconds": 10,
  "race_check": false,
  "platforms": ["darwin", "linux", "windows"],
  "hints": [
    "Think about how to track elements already seen.",
    "A map[int]bool or map[int]struct{} can represent membership.",
    "Iterate in input order and append only on first encounter."
  ]
}
```

The full catalog may be built from individual `exercise.json` files during development, but there must be **one deterministic, checked-in source of truth** for published ordering and IDs. Do not manually maintain two conflicting catalogs. Add JSON-schema-like validation in Go so invalid metadata fails CI.

Use stable, semantic IDs such as `slices_03`, not only global numbers. Ordinals may change for presentation; IDs must not change after release except through explicit migrations. Each topic has one canonical slug and source URL.

### 5.3 Starter format

- Top-of-file comment: title, objective, behavioral examples, source link, scope of permitted edits.
- Use TODO markers such as `// TODO:` or intentionally broken code.
- Prefer realistic APIs and Go idioms.
- Scope starter changes so the expected task is unambiguous.
- Tests must not depend on checking comments or exact formatting.
- A user should understand the task without needing to read the solution.
- Keep nontrivial instructions in `README.md` rather than huge source comment blocks.

Sample starter:

```go
package exercise

// Unique returns distinct integers in first-seen order.
// TODO: implement Unique. Empty input should produce an empty, non-nil slice.
func Unique(numbers []int) []int {
    return nil
}
```

Sample canonical check:

```go
package exercise

import (
    "reflect"
    "testing"
)

func TestUnique(t *testing.T) {
    tests := []struct {
        name string
        in, want []int
    }{
        {"repeated", []int{1, 2, 1, 3}, []int{1, 2, 3}},
        {"empty", []int{}, []int{}},
        {"negatives", []int{-1, 0, -1}, []int{-1, 0}},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Unique(tt.in); !reflect.DeepEqual(got, tt.want) {
                t.Fatalf("Unique(%v) = %v, want %v", tt.in, got, tt.want)
            }
        })
    }
}
```

### 5.4 Test case quality rubric

Every released exercise must have:

1. **Happy path** validating the central concept.
2. **Boundary/edge cases** appropriate to the function or program.
3. **Regression cases** for likely wrong implementations.
4. **Understandable failures**, including subtest names and `got/want` where possible.
5. **Reference-solution pass** and **starter fail** validations.
6. **No network dependency** or timing flakiness unless the curriculum specifically teaches networking, in which case local loopback must be used.
7. **Timeout control** for deadlocks, processes, and blocking operations.

For functions, consider empty/nil input, single element, duplicates, Unicode, negative numbers, overflow assumptions, and invalid input where appropriate. Avoid padding tests with irrelevant cases. For error exercises, verify sentinel/custom error behavior using `errors.Is` and `errors.As` where applicable. For resource management, verify close/cancel/cleanup behavior.

### 5.5 Concurrency guidelines

- Prefer barriers, test channels, injected clocks, or synchronization over `time.Sleep` assertions.
- Bound every receive, select, goroutine join, network call, and subprocess.
- Use `go test -race` for specified concurrency exercises on supported platforms where the race detector is available; do not treat absence of race capability as an automatic learner failure on unsupported configurations. Document fallback and enforce race CI on supported targets.
- A race-detector pass does not prove absence of deadlocks or ordering bugs; test both.
- Avoid asserting non-deterministic goroutine ordering unless the exercise's objective explicitly creates a guaranteed order.
- Ensure tests do not leak goroutines or sockets after completion.

### 5.6 Dangerous or special-behavior exercises

- **Panic/recover:** assert behavior inside a subprocess or recover locally; do not crash the runner.
- **`os.Exit` and signals:** spawn a child process and assert exit code/signal behavior. Never call `os.Exit` from the test runner itself.
- **Process spawning/execution:** use local helper executables or the Go test binary; do not require Unix-only commands (`ls`, `grep`, `sleep`, `sh`).
- **TCP/HTTP:** bind `127.0.0.1:0`, release ports properly, and ensure firewall/network failures are diagnosed clearly.
- **File IO:** use `t.TempDir()` and safe fixtures; no writes outside designated temp paths.
- **Environment:** isolate variables using appropriate test helpers; never modify the user's global environment.
- **Randomness:** inject deterministic seeds/sources; assert distributions only when scientifically reasonable and stable.
- **Timing:** control time where possible; do not depend on tight real-clock thresholds.
- **`go:embed`:** ensure embedding patterns and fixture paths are valid in staging.
- **Platform-specific signal behavior:** mark supported platforms or supply equivalent platform-specific tests; never label a Windows-specific skip as a pass for an untested requirement.

## 6. Verification pipeline and result model

### 6.1 Required pipeline

1. Validate metadata and asset integrity.
2. Ensure all expected learner files are present.
3. Stage source/tests/fixtures in isolation.
4. Run the exercise's compile/build/test/integration steps.
5. If enabled, run `go vet` or AST/lint constraints as **explicit declared requirements**.
6. If enabled and supported, run `go test -race`.
7. Parse results, determine a single outcome, display diagnostics.
8. On success, update progress atomically.

Do not impose a global requirement that every exercise must pass `go vet`; some starters intentionally violate its assumptions or teach behavior separately. Formatting is educational feedback, not a default hard gate: `gofmt -l` may be offered as an advisory check.

### 6.2 Typed results

```go
type Status string
const (
    StatusPass Status = "pass"
    StatusFail Status = "fail"
    StatusTimeout Status = "timeout"
    StatusRunnerError Status = "runner_error"
    StatusUnsupported Status = "unsupported"
)

type CheckResult struct {
    Name       string
    Status     Status
    Duration   time.Duration
    ExitCode   int
    Stdout     string
    Stderr     string
    Diagnostic string
}

type ExerciseResult struct {
    ExerciseID string
    Status     Status
    Checks     []CheckResult
    Duration   time.Duration
}
```

Never mark a runner error, skipped check, or unsupported required platform behavior as completed. Tests should emit a machine-readable event or JSON summary without corrupting human output.

### 6.3 Performance targets

- CLI basic commands should launch without a noticeable splash delay.
- Watch changes should be detected promptly (target < 500 ms debounce/recognition under normal local filesystem conditions).
- Simple exercises should generally finish in a few seconds after cache warm-up; Go compilation speed depends on host machine and cannot be guaranteed.
- Only one evaluation per exercise runs at a time. New writes cancel/supersede stale runs; never display an old result as current.
- Cache immutable catalog metadata; do not rewrite progress on every watch poll.

## 7. Progress and workspace lifecycle

### 7.1 Progress model

Per exercise track:

```json
{
  "schema_version": 1,
  "curriculum_version": "1.0.0",
  "current_exercise_id": "slices_03",
  "exercises": {
    "slices_03": {
      "status": "completed",
      "completed_at": "2026-11-17T16:00:00Z",
      "attempts": 4,
      "hints_revealed": 1,
      "source_fingerprint": "sha256:..."
    }
  }
}
```

Statuses: `not_started`, `in_progress`, `completed`, `skipped`. `failed` is the result of an evaluation and not a permanent workflow state. Completion history should remain visible if files change, but the UI must also indicate **modified since verified** when the fingerprint differs; a fresh passing verification is required to restore verified status. A missing exercise should not be counted complete simply because state says so.

**Write safety:** serialize state updates, write temp file in the same directory, fsync where appropriate, atomically rename, and recover from malformed state with a backup and an actionable error. Use a small lock or conflict detection to avoid concurrent writers trampling state. Do not use SQLite in v1.

### 7.2 Reset and solutions

- Reset shows what will be replaced and asks before any destructive operation.
- On reset, create a timestamped backup unless the user explicitly opts out.
- Show a solution to stdout or create a separate clearly named file; never silently replace learner code.
- Do not permanently mark completion just because the learner displayed a solution. Optionally show a `solution_viewed` indicator.

### 7.3 Upgrades

- Store `workspace_format_version` and `curriculum_version` separately.
- Never rewrite edited exercise files during an update.
- Newly introduced exercises can be added safely.
- Removed/renamed IDs require a deliberate migration table with a dry-run preview.
- Changes to canonical tests may cause a previously completed solution to require re-verification; communicate that accurately.
- Keep progress migrations reversible or backed up.

## 8. Curriculum scope and exact coverage map

**Source-of-truth count: 85 topics**, as listed at https://gobyexample.com/ on 2026-11-17. The section below is a **complete explicit checklist**, not a partial sample. Module categories are Golings organization, not necessarily the original website structure. Preserve the original conceptual ordering inside the manifest unless a documented prerequisite improves the learning experience.

**Exercise sizing rule:** at least **two meaningful exercises for each topic** = 170 minimum. Allocate the remaining **50 exercises** among high-impact/complex subjects to reach **220 core exercises**. Counts are planning targets and can be changed through curriculum review without relaxing the 85/85-topic coverage requirement. Optional capstones are tracked separately from the core count.

### Module 01 — Fundamentals (10 topics)

- [ ] 001 Hello World — `https://gobyexample.com/hello-world`
- [ ] 002 Values — `https://gobyexample.com/values`
- [ ] 003 Variables — `https://gobyexample.com/variables`
- [ ] 004 Constants — `https://gobyexample.com/constants`
- [ ] 005 For — `https://gobyexample.com/for`
- [ ] 006 If/Else — `https://gobyexample.com/if-else`
- [ ] 007 Switch — `https://gobyexample.com/switch`
- [ ] 008 Arrays — `https://gobyexample.com/arrays`
- [ ] 009 Slices — `https://gobyexample.com/slices`
- [ ] 010 Maps — `https://gobyexample.com/maps`

**Practice themes:** literals, declaration, zero values, shadowing, loop patterns, switch behavior, arrays versus slices, append/copy, nil versus empty collections, map membership. **Capstone idea:** text statistics from a collection of inputs.

### Module 02 — Functions and types (10 topics)

- [ ] 011 Functions — `https://gobyexample.com/functions`
- [ ] 012 Multiple Return Values — `https://gobyexample.com/multiple-return-values`
- [ ] 013 Variadic Functions — `https://gobyexample.com/variadic-functions`
- [ ] 014 Closures — `https://gobyexample.com/closures`
- [ ] 015 Recursion — `https://gobyexample.com/recursion`
- [ ] 016 Range over Built-in Types — `https://gobyexample.com/range-over-built-in-types`
- [ ] 017 Pointers — `https://gobyexample.com/pointers`
- [ ] 018 Strings and Runes — `https://gobyexample.com/strings-and-runes`
- [ ] 019 Structs — `https://gobyexample.com/structs`
- [ ] 020 Methods — `https://gobyexample.com/methods`

**Practice themes:** function boundaries, multiple-return errors, captures, recursive cases, value versus pointer semantics, UTF-8/runes, constructors, method receivers. **Capstone idea:** mini address book with validation and lookup.

### Module 03 — Interfaces, generics, errors (7 topics)

- [ ] 021 Interfaces — `https://gobyexample.com/interfaces`
- [ ] 022 Enums — `https://gobyexample.com/enums`
- [ ] 023 Struct Embedding — `https://gobyexample.com/struct-embedding`
- [ ] 024 Generics — `https://gobyexample.com/generics`
- [ ] 025 Range over Iterators — `https://gobyexample.com/range-over-iterators`
- [ ] 026 Errors — `https://gobyexample.com/errors`
- [ ] 027 Custom Errors — `https://gobyexample.com/custom-errors`

**Practice themes:** interface implementation, type assertions, idiomatic enum patterns, promoted methods, type parameters, iterator protocols, error wrapping and inspection. **Capstone idea:** composable data transformation pipeline.

### Module 04 — Concurrency (18 topics)

- [ ] 028 Goroutines — `https://gobyexample.com/goroutines`
- [ ] 029 Channels — `https://gobyexample.com/channels`
- [ ] 030 Channel Buffering — `https://gobyexample.com/channel-buffering`
- [ ] 031 Channel Synchronization — `https://gobyexample.com/channel-synchronization`
- [ ] 032 Channel Directions — `https://gobyexample.com/channel-directions`
- [ ] 033 Select — `https://gobyexample.com/select`
- [ ] 034 Timeouts — `https://gobyexample.com/timeouts`
- [ ] 035 Non-Blocking Channel Operations — `https://gobyexample.com/non-blocking-channel-operations`
- [ ] 036 Closing Channels — `https://gobyexample.com/closing-channels`
- [ ] 037 Range over Channels — `https://gobyexample.com/range-over-channels`
- [ ] 038 Timers — `https://gobyexample.com/timers`
- [ ] 039 Tickers — `https://gobyexample.com/tickers`
- [ ] 040 Worker Pools — `https://gobyexample.com/worker-pools`
- [ ] 041 WaitGroups — `https://gobyexample.com/waitgroups`
- [ ] 042 Rate Limiting — `https://gobyexample.com/rate-limiting`
- [ ] 043 Atomic Counters — `https://gobyexample.com/atomic-counters`
- [ ] 044 Mutexes — `https://gobyexample.com/mutexes`
- [ ] 045 Stateful Goroutines — `https://gobyexample.com/stateful-goroutines`

**Practice themes:** liveness, race safety, ownership, synchronization, cancellation, buffered/unbuffered communication, worker fan-out/fan-in, cleanup. Allocate many of the 50 additional exercises here. **Capstone idea:** cancellable concurrent job processor with bounded workers.

### Module 05 — Sorting and control (5 topics)

- [ ] 046 Sorting — `https://gobyexample.com/sorting`
- [ ] 047 Sorting by Functions — `https://gobyexample.com/sorting-by-functions`
- [ ] 048 Panic — `https://gobyexample.com/panic`
- [ ] 049 Defer — `https://gobyexample.com/defer`
- [ ] 050 Recover — `https://gobyexample.com/recover`

**Practice themes:** stable expectations, custom orderings, cleanup, panic boundaries, recover only where appropriate. **Capstone idea:** robust batch sorter and resource-cleanup harness.

### Module 06 — Standard library data formats and utilities (14 topics)

- [ ] 051 String Functions — `https://gobyexample.com/string-functions`
- [ ] 052 String Formatting — `https://gobyexample.com/string-formatting`
- [ ] 053 Text Templates — `https://gobyexample.com/text-templates`
- [ ] 054 Regular Expressions — `https://gobyexample.com/regular-expressions`
- [ ] 055 JSON — `https://gobyexample.com/json`
- [ ] 056 XML — `https://gobyexample.com/xml`
- [ ] 057 Time — `https://gobyexample.com/time`
- [ ] 058 Epoch — `https://gobyexample.com/epoch`
- [ ] 059 Time Formatting / Parsing — `https://gobyexample.com/time-formatting-parsing`
- [ ] 060 Random Numbers — `https://gobyexample.com/random-numbers`
- [ ] 061 Number Parsing — `https://gobyexample.com/number-parsing`
- [ ] 062 URL Parsing — `https://gobyexample.com/url-parsing`
- [ ] 063 SHA256 Hashes — `https://gobyexample.com/sha256-hashes`
- [ ] 064 Base64 Encoding — `https://gobyexample.com/base64-encoding`

**Practice themes:** transformations, serialization, validation, encoding, dates/time zones, templating, deterministic random tests. **Capstone idea:** JSON-to-text report generator.

### Module 07 — Files and directories (7 topics)

- [ ] 065 Reading Files — `https://gobyexample.com/reading-files`
- [ ] 066 Writing Files — `https://gobyexample.com/writing-files`
- [ ] 067 Line Filters — `https://gobyexample.com/line-filters`
- [ ] 068 File Paths — `https://gobyexample.com/file-paths`
- [ ] 069 Directories — `https://gobyexample.com/directories`
- [ ] 070 Temporary Files and Directories — `https://gobyexample.com/temporary-files-and-directories`
- [ ] 071 Embed Directive — `https://gobyexample.com/embed-directive`

**Practice themes:** streaming, buffered IO, paths, portable traversal, temp files, file errors, embedding. **Capstone idea:** local log inspector.

### Module 08 — Testing and CLIs (6 topics)

- [ ] 072 Testing and Benchmarking — `https://gobyexample.com/testing-and-benchmarking`
- [ ] 073 Command-Line Arguments — `https://gobyexample.com/command-line-arguments`
- [ ] 074 Command-Line Flags — `https://gobyexample.com/command-line-flags`
- [ ] 075 Command-Line Subcommands — `https://gobyexample.com/command-line-subcommands`
- [ ] 076 Environment Variables — `https://gobyexample.com/environment-variables`
- [ ] 077 Logging — `https://gobyexample.com/logging`

**Practice themes:** table-driven tests, benchmarks, flags, parser behavior, config precedence, structured logging. **Capstone idea:** a small multi-command utility.

### Module 09 — Networking, context and processes (8 topics)

- [ ] 078 HTTP Client — `https://gobyexample.com/http-client`
- [ ] 079 HTTP Server — `https://gobyexample.com/http-server`
- [ ] 080 TCP Server — `https://gobyexample.com/tcp-server`
- [ ] 081 Context — `https://gobyexample.com/context`
- [ ] 082 Spawning Processes — `https://gobyexample.com/spawning-processes`
- [ ] 083 Exec'ing Processes — `https://gobyexample.com/execing-processes`
- [ ] 084 Signals — `https://gobyexample.com/signals`
- [ ] 085 Exit — `https://gobyexample.com/exit`

**Practice themes:** local HTTP round-trips, handlers, TCP echo, cancellations, process lifecycle, signals, exit behavior. **Capstone idea:** graceful-shutdown HTTP service with CLI and request timeouts.

### 8.1 Curriculum coverage tool

The project must include `go run ./cmd/golings-dev validate-curriculum` (or a Go test with equivalent behavior) that:

1. Parses the checked-in Go by Example topic inventory.
2. Verifies each source topic has a unique slug and canonical URL.
3. Verifies every topic maps to >= 2 published exercise IDs by the complete-curriculum release.
4. Detects orphaned exercises and duplicates.
5. Verifies manifest counts, cross-links, prerequisite references, ordering, asset paths, and each exercise's reference URL.
6. Validates starters, canonical tests, solutions, fixture references, and expected results.
7. Produces an easily reviewable `docs/curriculum-coverage.md` and a machine-readable summary.

**Inventory versioning:** Do not perform an online fetch on every learner run. Maintain a checked-in `topics.json`, with an explicit optional maintainer command to compare it against the live site, report additions/removals/renames, and require human approval before updating. The course remains usable offline if the external site changes.

## 9. Content pedagogy

### 9.1 Difficulty and pacing

| Level | Typical scope | Expected effort |
|---|---|---|
| Beginner | One concept, one small function or simple program | 3–10 min |
| Intermediate | Multiple concepts, edge cases, interfaces, files | 10–20 min |
| Advanced | Concurrent/process/network integration, tricky semantics | 15–40 min |

For each topic, aim for a progression: (A) syntax/fix, (B) implement/behavior, (C) applied/debug if warranted. A third or fourth exercise is not mandatory for low-complexity topics. Make code intentionally broken for a reason, not just to manufacture friction.

### 9.2 Hints

At most three tiers per exercise by default:

1. Point toward the underlying concept or invariant.
2. Name a relevant Go feature/API without supplying the exact answer.
3. Give a near-complete strategy or pseudocode.

A solution reveal is separate and explicit. Hints should explain *why* a pattern works, not just instruct copying syntax.

### 9.3 Prerequisites and sequence

- Default curriculum is sequential.
- Users can jump to any exercise or module, even if prerequisites are unfinished; show a warning rather than enforce a hard lock.
- Each exercise declares only meaningful prerequisites.
- Capstones reinforce prior concepts but should not block all further progress.

### 9.4 Authorship and attribution

Original instructional text, starter code, checks, and solutions must be created for Golings. Link to the applicable Go by Example page as a reference and credit Rustlings/Ziglings for the learning model. Review license requirements for source material before incorporating any copied content. Do not wholesale rehost website examples or descriptions without checking the applicable license and attribution obligations.

## 10. Testing strategy for Golings itself

### 10.1 Unit tests

- Catalog parsing, validation, ID matching, ordering.
- Progress state transitions and serialization, corruption recovery.
- Safe workspace creation, reset/backup, symlink/path traversal checks.
- Diagnostic classification and output truncation.
- CLI validation and machine-readable rendering.

### 10.2 Integration tests

- Execute the CLI in an isolated temp workspace with a controlled Go toolchain.
- Verify first run, compile error, unit test failure, save/recheck, passing transition, restart/resume.
- Verify user changes aren't overwritten by init/update/reset without confirmation.
- Exercise staged source, fixture handling, network loopback, subprocess/timeouts.
- Verify Ctrl+C cancellation and stale result suppression.
- Test JSON output and exit codes.

### 10.3 Curriculum contract tests (mandatory)

For **every** published exercise:

- [ ] Starter files parse/compile or fail in the specifically documented expected way.
- [ ] Running the starter **does not pass all checks** (unless an intentionally instructional exception is explicitly approved and flagged).
- [ ] The canonical solution passes all required tests.
- [ ] A minimal mutation or likely incorrect solution fails at least one check (where applicable).
- [ ] No resource leaks, unbounded waiting, or external network calls.
- [ ] Metadata references valid source, hints, solution, fixtures, and test assets.
- [ ] Test failures yield actionable results.

### 10.4 CI matrix

- `go test ./...` and `go vet ./...` for the **Golings implementation**.
- Curriculum authoring validator and solution/starter contract checks.
- Supported `GOOS/GOARCH` builds: macOS amd64/arm64, Linux amd64/arm64, Windows amd64/arm64 where toolchain/release infrastructure supports them.
- End-to-end smoke tests on Linux, macOS, Windows runners.
- `-race` on supported CI runner platforms for Golings and selected concurrency exercises.
- Formatting check (`gofmt`) on maintained implementation and solution source; intentionally malformed starter code is exempt.

## 11. Security, privacy, and robustness

- No telemetry, accounts, or outbound requests for basic operations.
- Never execute commands based on unvalidated metadata strings; use fixed allowlisted Go tool commands and explicit argument arrays.
- Keep temporary staging paths separate from user workspace and clean them up.
- Avoid following workspace symlinks beyond expected directories; validate traversal and canonical paths.
- Do not automatically download arbitrary Go modules when verifying stdlib exercises. Prefer `GOPROXY=off` for staged stdlib work.
- No security guarantee against intentionally malicious user code. Document that exercises run locally with the user's privileges. Containerization, OS-level sandboxing, or untrusted-code hosting is a separate future project.
- Child process cancellation must prevent endless test hangs; test process tree behavior per OS.
- Avoid collecting or exposing user paths in logs unnecessarily.

## 12. Release strategy and milestones

### Milestone A — Runner foundation (minimum usable vertical slice)

- [ ] CLI and Go toolchain detection.
- [ ] Manifest loader and validation.
- [ ] `init`, `list`, `run`, `verify`, `next`, `progress`.
- [ ] Staging and evaluation for compile/unit checks.
- [ ] Local progress state with atomic writes.
- [ ] 5 high-quality sample exercises across simple types.
- [ ] End-to-end tests for passing/failing exercises.

**Exit criteria:** One new learner can complete a fully working end-to-end exercise path, restart, and resume.

### Milestone B — Rustlings-like ergonomics

- [ ] Default watch mode with debounce and cancel/retry.
- [ ] Hints, reset with backup, solution reveal.
- [ ] Excellent terminal diagnostics, no-color, JSON output.
- [ ] Windows/macOS/Linux smoke tests.
- [ ] Approx. 25–30 polished beginner exercises.

**Exit criteria:** Exercise editing and feedback feel effortless, and bad-input/timeout behavior is understandable.

### Milestone C — Curriculum expansion

- [ ] Modules 01–03 completed.
- [ ] Modules 04–06 completed.
- [ ] Advanced check modes: race, subprocess, loopback network, file fixtures.
- [ ] Modules 07–09 completed.
- [ ] At least 2 exercises per source topic; deliberate allocation to 220 core exercises.
- [ ] Optional module capstones.

**Exit criteria:** Full reconciled topic inventory represented, with passing quality gates.

### Milestone D — Release hardening

- [ ] Update/migration policy and versioning.
- [ ] Automated release archives, checksums, install instructions.
- [ ] Contributor authoring guide and code of conduct if needed.
- [ ] Full CI matrix and curriculum contract test suite.
- [ ] Documentation, troubleshooting, source attribution, licensing audit.

**Exit criteria:** A newcomer can use the packaged program without development tooling beyond Go, and maintainers can accept new exercises without breaking old ones.

## 13. Acceptance criteria and release gates

### Functional

- [ ] User can initialize a safe workspace, run/watch an exercise, edit code, see feedback, complete, and resume.
- [ ] All documented CLI commands work with consistent errors and exit codes.
- [ ] Each exercise can be evaluated independently; no accidental shared state.
- [ ] Every published exercise has hints and a solution.
- [ ] Progress survives restart and tracks modified-since-verified state.
- [ ] Destructive operations prompt and preserve backups by default.

### Curriculum

- [ ] Exact Go by Example source inventory is reconciled and versioned.
- [ ] 100% topic coverage with original exercises, correct URLs, and meaningful tests.
- [ ] At least two meaningful exercises per topic, and 220 core exercises as the intended scope.
- [ ] Every starter fails intentionally; every solution passes.
- [ ] Cross-platform and advanced topics have explicit supported test modes.

### Engineering

- [ ] Entire application implemented in Go.
- [ ] No required network access or account.
- [ ] Reliable cleanup and timeouts on exercises that can hang.
- [ ] CI passes on target OSes.
- [ ] No silent data loss during init, reset, or upgrades.
- [ ] Licenses and attribution reviewed before publication.

## 14. Open design decisions (resolve before implementation or first beta)

1. **Final name:** `golings` is a placeholder; check naming/trademark and package availability before publishing.
2. **Distribution:** GitHub Releases plus `go install` is the simplest path; package managers can follow.
3. **Go version baseline:** choose a specific supported Go version and CI matrix at implementation time. Modern topics (generics and range-over-iterators) require corresponding language versions. Diagnose unsupported toolchains clearly.
4. **Watch implementation:** standard-library polling first; optional native filesystem notifications later.
5. **Solution packaging:** embed with executable for offline use versus distribute as separate repo files. This spec assumes embedded assets.
6. **Exercise numbering:** stable semantic IDs and display ordinals; don't encode strict global ordering into file names unless desired.
7. **Test visibility:** canonical tests are embedded and not automatically copied; some public examples may be shipped for learning.
8. **Capstone progress:** keep capstones optional and separate from 220 core denominator unless explicitly revised.
9. **Live source-site monitoring:** manual maintainer command, not a dependency of learner runs.

## 15. Suggested first implementation tasks

1. Create the Go module and `cmd/golings/main.go` with `help`, `version`, `doctor`.
2. Define catalog and exercise metadata structs with strict validation.
3. Build a single exercise asset with starter, tests, solution, and hints.
4. Implement `init` using embedded assets and safe destination checks.
5. Implement the staging runner using `go test` with cancellation, timeout, and captured diagnostics.
6. Implement `verify`, then progress persistence, then `next`.
7. Add watch mode with a simple periodic content hash and debounce.
8. Add `list`, `hint`, `reset`, `solution`, and JSON output.
9. Build the per-exercise authoring validator and first 25–30 exercises.
10. Expand curriculum module-by-module only after the runner is stable.

---

## Appendix A — Canonical references

- Go by Example topic index: https://gobyexample.com/
- Official Go docs: https://go.dev/doc/
- Go testing package: https://pkg.go.dev/testing
- Go command docs: https://pkg.go.dev/cmd/go
- Rustlings: https://github.com/rust-lang/rustlings
- Ziglings: https://ziglings.org/

**Maintenance note:** The topic checklist is a pinned snapshot. Before claiming full coverage in a release, compare it with the current source index and update the checked-in inventory through a reviewed change.
