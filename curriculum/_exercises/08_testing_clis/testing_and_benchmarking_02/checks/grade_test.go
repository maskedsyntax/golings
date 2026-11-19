package reverse

import (
	"errors"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// golings grader: your tests must pass against the real Reverse and fail
// against each broken version below.

type bug struct {
	name, desc string
	fn         func(string) string
}

var bugs = []bug{
	{"bytes", "Reverse reverses bytes instead of runes", func(s string) string {
		b := []byte(s)
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
		return string(b)
	}},
	{"keep_first", "Reverse leaves the first character in place", func(s string) string {
		r := []rune(s)
		if len(r) < 2 {
			return s
		}
		rest := r[1:]
		for i, j := 0, len(rest)-1; i < j; i, j = i+1, j-1 {
			rest[i], rest[j] = rest[j], rest[i]
		}
		return string(r)
	}},
	{"empty_panics", "Reverse panics on the empty string", func(s string) string {
		r := []rune(s)
		_ = s[0]
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}},
}

func TestMain(m *testing.M) {
	if name := os.Getenv("GOLINGS_BUG"); name != "" {
		for _, b := range bugs {
			if b.name == name {
				Reverse = b.fn
			}
		}
	}
	os.Exit(m.Run())
}

func runYourTests(t *testing.T, env string, args ...string) ([]byte, bool) {
	t.Helper()
	args = append([]string{"-test.run=^Test", "-test.skip=^TestYour", "-test.count=1"}, args...)
	cmd := exec.Command(os.Args[0], args...)
	cmd.Env = append(os.Environ(), env)
	out, err := cmd.CombinedOutput()
	var ee *exec.ExitError
	switch {
	case err == nil:
		return out, true
	case errors.As(err, &ee):
		return out, false
	}
	t.Fatalf("could not run your tests: %v\n%s", err, out)
	return nil, false
}

func TestYourTestsCatchBugs(t *testing.T) {
	if os.Getenv("GOLINGS_BUG") != "" {
		t.Skip("running inside a bug check")
	}
	for _, b := range bugs {
		if _, passed := runYourTests(t, "GOLINGS_BUG="+b.name); passed {
			t.Errorf("your tests still pass when %s; add a case that catches it", b.desc)
		}
	}
}

func TestYourTestsUseSubtests(t *testing.T) {
	if os.Getenv("GOLINGS_BUG") != "" {
		t.Skip("running inside a bug check")
	}
	out, passed := runYourTests(t, "GOLINGS_SUBTESTS=1", "-test.v")
	if !passed {
		t.Fatalf("your tests fail against the correct Reverse:\n%s", out)
	}
	if !strings.Contains(string(out), "=== RUN   TestReverse/") {
		t.Errorf("TestReverse should run each case as a subtest with t.Run")
	}
}
