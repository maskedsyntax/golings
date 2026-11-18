package runner

import (
	"time"

	"github.com/maskedsyntax/golings/internal/diagnostics"
)

// Status is the outcome of a check or an exercise (spec §6.2).
type Status string

const (
	StatusPass        Status = "pass"
	StatusFail        Status = "fail"
	StatusTimeout     Status = "timeout"
	StatusRunnerError Status = "runner_error"
	StatusUnsupported Status = "unsupported"
)

// severity orders statuses for aggregation:
// runner_error > timeout > fail > unsupported > pass.
func (s Status) severity() int {
	switch s {
	case StatusRunnerError:
		return 4
	case StatusTimeout:
		return 3
	case StatusFail:
		return 2
	case StatusUnsupported:
		return 1
	default:
		return 0
	}
}

// Worst returns the more severe of a and b.
func Worst(a, b Status) Status {
	if b.severity() > a.severity() {
		return b
	}
	return a
}

// Failure classifies why a check did not pass, so the UI can tell a compiler
// error from a failing test, a timeout, a data race, or a runner problem.
type Failure string

const (
	FailureNone       Failure = ""
	FailureCompile    Failure = "compile"
	FailureTest       Failure = "test"
	FailureOutput     Failure = "output"
	FailureVet        Failure = "vet"
	FailureRace       Failure = "race"
	FailureConstraint Failure = "constraint" // an ast rule is not satisfied
	FailureTimeout    Failure = "timeout"
	FailureRunner     Failure = "runner"
)

// TestResult is one test (or subtest) reported by `go test -json`.
type TestResult struct {
	Name   string `json:"name"`
	Status Status `json:"status"` // pass, fail, or unsupported (skipped)
	Output string `json:"output,omitempty"`
}

// CheckResult is the outcome of one declared check.
type CheckResult struct {
	Name        string             `json:"name"`
	Status      Status             `json:"status"`
	Failure     Failure            `json:"failure,omitempty"`
	Duration    time.Duration      `json:"-"`
	DurationMS  int64              `json:"duration_ms"`
	ExitCode    int                `json:"exit_code"`
	Stdout      string             `json:"stdout,omitempty"`
	Stderr      string             `json:"stderr,omitempty"`
	Truncated   bool               `json:"truncated,omitempty"`
	Diagnostic  string             `json:"diagnostic,omitempty"`
	Diagnostics []diagnostics.Diag `json:"diagnostics,omitempty"`
	Tests       []TestResult       `json:"tests,omitempty"`
	Expected    string             `json:"expected,omitempty"` // output checks
	Actual      string             `json:"actual,omitempty"`   // output checks
}

// ExerciseResult aggregates every check run for an exercise.
type ExerciseResult struct {
	ExerciseID string        `json:"exercise_id"`
	Status     Status        `json:"status"`
	Checks     []CheckResult `json:"checks"`
	Duration   time.Duration `json:"-"`
	DurationMS int64         `json:"duration_ms"`
}

// Passed reports whether every required check passed. Unsupported checks
// never count as passing.
func (r *ExerciseResult) Passed() bool { return r.Status == StatusPass }

// FirstProblem returns the first check that did not pass, or nil.
func (r *ExerciseResult) FirstProblem() *CheckResult {
	for i := range r.Checks {
		if r.Checks[i].Status != StatusPass {
			return &r.Checks[i]
		}
	}
	return nil
}
