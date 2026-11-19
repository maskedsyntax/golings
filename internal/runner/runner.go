// Package runner evaluates an exercise: it stages the learner's source with
// the canonical checks in a temporary directory outside the workspace, runs
// the declared checks with the local Go toolchain, and returns typed results.
package runner

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/maskedsyntax/golings/internal/astcheck"
	"github.com/maskedsyntax/golings/internal/catalog"
	"github.com/maskedsyntax/golings/internal/diagnostics"
	"github.com/maskedsyntax/golings/internal/toolchain"
)

const isWindows = runtime.GOOS == "windows"

// compileAllowance is added on top of the exercise timeout for steps that
// compile, so a cold build cache is not mistaken for a hanging exercise.
// The exercise's own code is still bounded by timeout_seconds.
const compileAllowance = 2 * time.Minute

// Request describes one evaluation.
type Request struct {
	Exercise  *catalog.Exercise
	Files     map[string][]byte // learner source files by base name
	Assets    fs.FS             // exercise asset directory (checks/, fixtures/)
	Toolchain *toolchain.Info
	// Timeout overrides the exercise's timeout_seconds when positive.
	Timeout time.Duration
	// DisplayDir is the learner-facing directory of the exercise, used to
	// rewrite diagnostics. Empty leaves bare file names.
	DisplayDir string
}

// Run evaluates the exercise. It stops at the first check that does not
// pass, since later checks are meaningless if, say, the code doesn't compile.
// Errors are reported inside the result (StatusRunnerError), never returned.
func Run(ctx context.Context, req Request) *ExerciseResult {
	start := time.Now()
	ex := req.Exercise
	res := &ExerciseResult{ExerciseID: ex.ID, Status: StatusPass}
	finish := func() *ExerciseResult {
		res.Duration = time.Since(start)
		res.DurationMS = res.Duration.Milliseconds()
		return res
	}

	if !ex.SupportsOS(req.Toolchain.GOOS) {
		res.Status = StatusUnsupported
		res.Checks = append(res.Checks, CheckResult{
			Name:       "platform",
			Status:     StatusUnsupported,
			Diagnostic: fmt.Sprintf("this exercise supports %s and cannot be verified on %s", strings.Join(ex.Platforms, ", "), req.Toolchain.GOOS),
		})
		return finish()
	}

	st, err := newStage(ex, req.Files, req.Assets)
	if err != nil {
		res.Status = StatusRunnerError
		res.Checks = append(res.Checks, CheckResult{Name: "stage", Status: StatusRunnerError, Failure: FailureRunner, Diagnostic: err.Error()})
		return finish()
	}
	defer st.cleanup()

	timeout := req.Timeout
	if timeout <= 0 {
		timeout = time.Duration(ex.TimeoutSeconds) * time.Second
	}
	learner := map[string]bool{}
	for _, f := range ex.StarterFiles {
		learner[f] = true
	}
	c := &checker{
		req:     req,
		stage:   st,
		env:     buildEnv(req.Toolchain),
		timeout: timeout,
		rw:      diagnostics.Rewriter{StagingDir: st.dir, DisplayDir: req.DisplayDir, Learner: learner},
	}

	for _, kind := range ex.Checks {
		if ctx.Err() != nil {
			break
		}
		var cr CheckResult
		switch kind {
		case catalog.CheckBuild:
			cr = c.build(ctx)
		case catalog.CheckUnit:
			cr = c.goTest(ctx, false)
		case catalog.CheckRace:
			cr = c.goTest(ctx, true)
		case catalog.CheckVet:
			cr = c.vet(ctx)
		case catalog.CheckOutput:
			cr = c.output(ctx)
		case catalog.CheckAST:
			cr = c.ast()
		default:
			cr = CheckResult{Status: StatusRunnerError, Failure: FailureRunner, Diagnostic: fmt.Sprintf("unknown check kind %q", kind)}
		}
		cr.Name = string(kind)
		cr.DurationMS = cr.Duration.Milliseconds()
		res.Checks = append(res.Checks, cr)
		res.Status = Worst(res.Status, cr.Status)
		if cr.Status != StatusPass && cr.Status != StatusUnsupported {
			break
		}
	}
	if err := ctx.Err(); err != nil && res.Status == StatusPass {
		// Canceled before every check ran: never report a partial pass.
		res.Status = StatusRunnerError
		res.Checks = append(res.Checks, CheckResult{Name: "canceled", Status: StatusRunnerError, Failure: FailureRunner, Diagnostic: "evaluation canceled"})
	}
	return finish()
}

type checker struct {
	req     Request
	stage   *stage
	env     []string
	timeout time.Duration
	rw      diagnostics.Rewriter
}

func (c *checker) goCmd(ctx context.Context, limit time.Duration, stdoutLimit int, args ...string) execResult {
	ctx, cancel := context.WithTimeout(ctx, limit)
	defer cancel()
	return execCommand(ctx, c.stage.dir, c.env, stdoutLimit, c.req.Toolchain.GoBin, args...)
}

// infraResult converts start failures, cancellation and outer timeouts into a
// CheckResult. ok is false when r is not a normal completed run.
func (c *checker) infraResult(r execResult, what string) (CheckResult, bool) {
	cr := CheckResult{Duration: r.Duration, ExitCode: r.ExitCode}
	switch {
	case r.StartErr != nil:
		cr.Status, cr.Failure = StatusRunnerError, FailureRunner
		cr.Diagnostic = fmt.Sprintf("could not run %s: %v", what, r.StartErr)
	case r.Canceled:
		cr.Status, cr.Failure = StatusRunnerError, FailureRunner
		cr.Diagnostic = "evaluation canceled"
	case r.TimedOut:
		cr.Status, cr.Failure = StatusTimeout, FailureTimeout
		cr.Diagnostic = fmt.Sprintf("%s did not finish in time and was stopped; look for an infinite loop or a blocked channel operation", what)
		cr.Stderr = c.rw.Rewrite(string(r.Stderr))
	default:
		return cr, true
	}
	return cr, false
}

func (c *checker) build(ctx context.Context) CheckResult {
	r := c.goCmd(ctx, compileAllowance, outputLimit, "build", "-o", c.binPath(), ".")
	return c.compileResult(r, "go build")
}

func (c *checker) compileResult(r execResult, what string) CheckResult {
	cr, ok := c.infraResult(r, what)
	if !ok {
		return cr
	}
	if r.ExitCode == 0 {
		cr.Status = StatusPass
		return cr
	}
	out := c.rw.Rewrite(string(r.Stderr) + string(r.Stdout))
	cr.Status, cr.Failure = StatusFail, FailureCompile
	cr.Stderr, cr.Truncated = out, r.Truncated
	cr.Diagnostics = diagnostics.Parse(out)
	cr.Diagnostic = "your code does not compile"
	return cr
}

func (c *checker) vet(ctx context.Context) CheckResult {
	r := c.goCmd(ctx, compileAllowance, outputLimit, "vet", ".")
	cr := c.compileResult(r, "go vet")
	if cr.Failure == FailureCompile {
		cr.Failure = FailureVet
		cr.Diagnostic = "go vet reported problems"
	}
	return cr
}

func (c *checker) goTest(ctx context.Context, race bool) CheckResult {
	if race && !c.req.Toolchain.RaceSupported {
		return CheckResult{
			Status:     StatusUnsupported,
			Diagnostic: fmt.Sprintf("the race detector is not available for %s/%s (cgo enabled: %v); this check cannot be verified here", c.req.Toolchain.GOOS, c.req.Toolchain.GOARCH, c.req.Toolchain.CGOEnabled),
		}
	}
	args := []string{"test", "-json", "-count=1", "-timeout=" + c.timeout.String()}
	if race {
		args = append(args, "-race")
	}
	args = append(args, ".")
	// The test binary enforces c.timeout itself (and prints goroutine dumps);
	// the outer limit also covers compilation.
	r := c.goCmd(ctx, c.timeout+compileAllowance, testJSONLimit, args...)
	cr, ok := c.infraResult(r, "go test")
	if !ok {
		return cr
	}
	rep := parseTestJSON(r.Stdout)
	cr.Truncated = r.Truncated
	for _, t := range leafTests(rep.Tests) {
		t.Output = c.rw.Rewrite(t.Output)
		cr.Tests = append(cr.Tests, t)
	}
	stderr := c.rw.Rewrite(string(r.Stderr))

	switch {
	case rep.BuildFailed || (r.ExitCode != 0 && len(rep.Tests) == 0 && !rep.TimedOut && !rep.Panicked && !rep.PackageFailed):
		out := c.rw.Rewrite(rep.BuildOutput + rep.PackageOutput)
		if strings.TrimSpace(out) == "" {
			out = stderr
		} else if stderr != "" {
			out += stderr
		}
		cr.Status, cr.Failure = StatusFail, FailureCompile
		cr.Stderr = strings.TrimSpace(out)
		cr.Diagnostics = diagnostics.Parse(out)
		cr.Diagnostic = "your code does not compile"
		cr.Tests = nil
	case rep.TimedOut:
		cr.Status, cr.Failure = StatusTimeout, FailureTimeout
		cr.Diagnostic = fmt.Sprintf("tests did not finish within %s; look for an infinite loop, a deadlock or a blocked channel operation", c.timeout)
		cr.Stderr, _ = diagnostics.Truncate(c.rw.Rewrite(rep.PackageOutput), outputLimit)
	case rep.Race:
		cr.Status, cr.Failure = StatusFail, FailureRace
		cr.Diagnostic = "the race detector found a data race"
		cr.Stderr, _ = diagnostics.Truncate(c.rw.Rewrite(rep.PackageOutput), outputLimit)
	case r.ExitCode == 0:
		cr.Status = StatusPass
	case rep.PackageFailed || anyFailed(rep.Tests):
		cr.Status, cr.Failure = StatusFail, FailureTest
		cr.Diagnostic = "some checks failed"
		if rep.Panicked {
			cr.Diagnostic = "a check panicked"
			cr.Stderr, _ = diagnostics.Truncate(c.rw.Rewrite(rep.PackageOutput), outputLimit)
		}
	default:
		cr.Status, cr.Failure = StatusRunnerError, FailureRunner
		cr.Diagnostic = "go test failed unexpectedly (exit code " + strconv.Itoa(r.ExitCode) + ")"
		cr.Stderr = stderr
	}
	if cr.Status == StatusFail && cr.Failure != FailureCompile && loopbackBlocked(rep, cr.Tests) {
		// Not the learner's fault: the checks could not open a local socket.
		cr.Status, cr.Failure = StatusRunnerError, FailureRunner
		cr.Diagnostic = "the checks could not listen on the loopback network (127.0.0.1). " +
			"A firewall, sandbox or security tool is probably blocking local connections; allow them for the Go test binary and try again"
	}
	return cr
}

// loopbackListenFailure matches the errors net.Listen and httptest report
// when the OS refuses a local socket.
var loopbackListenFailure = regexp.MustCompile(
	`(?i)(listen tcp[46]? [^\s]*:\d+: (bind|socket): (operation not permitted|permission denied|address family not supported|can't assign requested address|an attempt was made to access a socket))` +
		`|httptest: failed to listen on a port`)

func loopbackBlocked(rep testReport, tests []TestResult) bool {
	if loopbackListenFailure.MatchString(rep.PackageOutput) {
		return true
	}
	for _, t := range tests {
		if loopbackListenFailure.MatchString(t.Output) {
			return true
		}
	}
	return false
}

func anyFailed(tests []TestResult) bool {
	for _, t := range tests {
		if t.Status == StatusFail {
			return true
		}
	}
	return false
}

func (c *checker) binPath() string {
	name := "exercise"
	if isWindows {
		name += ".exe"
	}
	return filepath.Join(c.stage.dir, ".golings-bin", name)
}

// output builds the program, runs it with no input, and compares its
// standard output with checks/expected_stdout.txt.
func (c *checker) output(ctx context.Context) CheckResult {
	expected, err := fs.ReadFile(c.req.Assets, "checks/expected_stdout.txt")
	if err != nil {
		return CheckResult{Status: StatusRunnerError, Failure: FailureRunner, Diagnostic: "curriculum error: missing checks/expected_stdout.txt"}
	}
	b := c.build(ctx)
	if b.Status != StatusPass {
		return b
	}
	runCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	r := execCommand(runCtx, c.stage.dir, c.env, outputLimit, c.binPath())
	cr, ok := c.infraResult(r, "your program")
	cr.Duration += b.Duration
	if !ok {
		cr.Stdout = string(r.Stdout)
		return cr
	}
	got := normalizeOutput(r.Stdout)
	want := normalizeOutput(expected)
	cr.Stdout, cr.Stderr, cr.Truncated = string(r.Stdout), c.rw.Rewrite(string(r.Stderr)), r.Truncated
	cr.Expected, cr.Actual = want, got
	switch {
	case r.ExitCode != 0:
		cr.Status, cr.Failure = StatusFail, FailureOutput
		cr.Diagnostic = fmt.Sprintf("your program exited with status %d", r.ExitCode)
	case got != want:
		cr.Status, cr.Failure = StatusFail, FailureOutput
		cr.Diagnostic = describeMismatch(got, want)
	default:
		cr.Status = StatusPass
	}
	return cr
}

// normalizeOutput makes comparisons line-ending agnostic and ignores
// trailing newlines.
func normalizeOutput(b []byte) string {
	b = bytes.ReplaceAll(b, []byte("\r\n"), []byte("\n"))
	return strings.TrimRight(string(b), "\n")
}

func describeMismatch(got, want string) string {
	gl, wl := strings.Split(got, "\n"), strings.Split(want, "\n")
	if got == "" {
		gl = nil
	}
	for i := 0; i < max(len(gl), len(wl)); i++ {
		var g, w string
		if i < len(gl) {
			g = gl[i]
		}
		if i < len(wl) {
			w = wl[i]
		}
		if i >= len(gl) {
			return fmt.Sprintf("output is missing line %d: want %q", i+1, w)
		}
		if i >= len(wl) {
			return fmt.Sprintf("output has an extra line %d: %q", i+1, g)
		}
		if g != w {
			return fmt.Sprintf("line %d differs:\n   got: %q\n  want: %q", i+1, g, w)
		}
	}
	return "output differs"
}

// ast evaluates the exercise's syntax constraints in-process.
func (c *checker) ast() CheckResult {
	start := time.Now()
	violations, err := astcheck.Check(c.req.Files, c.req.Exercise.ASTRules)
	cr := CheckResult{Duration: time.Since(start)}
	switch {
	case err != nil:
		cr.Status, cr.Failure = StatusFail, FailureCompile
		cr.Diagnostic = "your code does not compile"
		cr.Stderr = c.rw.Rewrite(err.Error())
		cr.Diagnostics = diagnostics.Parse(cr.Stderr)
	case len(violations) > 0:
		cr.Status, cr.Failure = StatusFail, FailureConstraint
		cr.Diagnostic = "the code works, but the exercise asks for a specific construct"
		for _, v := range violations {
			cr.Tests = append(cr.Tests, TestResult{Name: v.Rule, Status: StatusFail, Output: v.Message})
		}
	default:
		cr.Status = StatusPass
	}
	return cr
}
