package runner

import (
	"bufio"
	"bytes"
	"encoding/json"
	"strings"
)

// testEvent mirrors cmd/test2json's output (`go test -json`). Since Go 1.24,
// build output and failures are reported as JSON too (build-output,
// build-fail).
type testEvent struct {
	Action      string
	Package     string
	ImportPath  string
	Test        string
	Output      string
	FailedBuild string
}

// testReport is the parsed form of a `go test -json` run.
type testReport struct {
	BuildFailed   bool
	BuildOutput   string
	PackageOutput string
	PackageFailed bool
	Tests         []TestResult
	TimedOut      bool // the test binary hit -timeout
	Race          bool // the race detector reported a data race
	Panicked      bool
}

const perTestOutputLimit = 8 << 10

func parseTestJSON(stdout []byte) testReport {
	var r testReport
	index := map[string]int{}
	var build, pkg strings.Builder
	outputs := map[string]*strings.Builder{}

	sc := bufio.NewScanner(bytes.NewReader(stdout))
	sc.Buffer(make([]byte, 64<<10), 4<<20)
	for sc.Scan() {
		line := sc.Bytes()
		var ev testEvent
		if len(line) == 0 || line[0] != '{' || json.Unmarshal(line, &ev) != nil {
			// Non-JSON noise (for example, from a truncated stream).
			pkg.Write(line)
			pkg.WriteByte('\n')
			continue
		}
		switch {
		case ev.Action == "build-output":
			build.WriteString(ev.Output)
		case ev.Action == "build-fail":
			r.BuildFailed = true
		case ev.Test != "":
			i, ok := index[ev.Test]
			if !ok {
				i = len(r.Tests)
				index[ev.Test] = i
				r.Tests = append(r.Tests, TestResult{Name: ev.Test})
				outputs[ev.Test] = &strings.Builder{}
			}
			switch ev.Action {
			case "output":
				if b := outputs[ev.Test]; b.Len() < perTestOutputLimit {
					b.WriteString(ev.Output)
				}
				scanSignals(&r, ev.Output)
			case "pass":
				r.Tests[i].Status = StatusPass
			case "fail":
				r.Tests[i].Status = StatusFail
			case "skip":
				r.Tests[i].Status = StatusUnsupported
			}
		default: // package-level event
			switch ev.Action {
			case "output":
				pkg.WriteString(ev.Output)
				scanSignals(&r, ev.Output)
			case "fail":
				r.PackageFailed = true
				if ev.FailedBuild != "" {
					r.BuildFailed = true
				}
			}
		}
	}
	for i := range r.Tests {
		r.Tests[i].Output = cleanTestOutput(outputs[r.Tests[i].Name].String())
	}
	r.BuildOutput = build.String()
	r.PackageOutput = pkg.String()
	if strings.Contains(r.PackageOutput, "[build failed]") || strings.Contains(r.PackageOutput, "[setup failed]") {
		r.BuildFailed = true
	}
	return r
}

func scanSignals(r *testReport, out string) {
	switch {
	case strings.Contains(out, "panic: test timed out after"):
		r.TimedOut = true
	case strings.Contains(out, "WARNING: DATA RACE"), strings.Contains(out, "race detected during execution of test"):
		r.Race = true
	case strings.HasPrefix(out, "panic: "):
		r.Panicked = true
	}
}

// cleanTestOutput drops go test's framing lines (=== RUN, --- FAIL, ...) and
// keeps what the test itself logged.
func cleanTestOutput(s string) string {
	var b strings.Builder
	for _, line := range strings.SplitAfter(s, "\n") {
		t := strings.TrimSpace(line)
		if t == "" || strings.HasPrefix(t, "=== ") || strings.HasPrefix(t, "--- PASS") ||
			strings.HasPrefix(t, "--- FAIL") || strings.HasPrefix(t, "--- SKIP") {
			continue
		}
		b.WriteString(line)
	}
	return strings.TrimRight(b.String(), "\n")
}

// leafTests returns tests that have no subtests: the ones a learner acts on.
func leafTests(tests []TestResult) []TestResult {
	var out []TestResult
	for i, t := range tests {
		parent := false
		for j, u := range tests {
			if i != j && strings.HasPrefix(u.Name, t.Name+"/") {
				parent = true
				break
			}
		}
		if !parent {
			out = append(out, t)
		}
	}
	return out
}
