package join

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"sync/atomic"
	"testing"
)

// golings grader: runs your benchmark briefly and counts the calls it makes.

func TestMain(m *testing.M) {
	if os.Getenv("GOLINGS_COUNT") == "" {
		os.Exit(m.Run())
	}
	var plus, builder atomic.Int64
	origPlus, origBuilder := JoinPlus, JoinBuilder
	JoinPlus = func(p []string) string { plus.Add(1); return origPlus(p) }
	JoinBuilder = func(p []string) string { builder.Add(1); return origBuilder(p) }
	code := m.Run()
	fmt.Printf("golings-calls plus=%d builder=%d\n", plus.Load(), builder.Load())
	os.Exit(code)
}

var callsRE = regexp.MustCompile(`golings-calls plus=(\d+) builder=(\d+)`)

func TestYourBenchmark(t *testing.T) {
	if os.Getenv("GOLINGS_COUNT") != "" {
		t.Skip("running inside the benchmark check")
	}
	cmd := exec.Command(os.Args[0], "-test.run=^$", "-test.bench=^BenchmarkJoin$", "-test.benchtime=20x")
	cmd.Env = append(os.Environ(), "GOLINGS_COUNT=1")
	out, err := cmd.CombinedOutput()
	var ee *exec.ExitError
	if errors.As(err, &ee) {
		t.Fatalf("your benchmark failed:\n%s", out)
	} else if err != nil {
		t.Fatalf("could not run your benchmark: %v", err)
	}
	for _, sub := range []string{"BenchmarkJoin/plus", "BenchmarkJoin/builder"} {
		if !regexp.MustCompile(sub + `\b`).Match(out) {
			t.Errorf("no %s sub-benchmark ran; use b.Run(%q, ...)\noutput:\n%s", sub, sub[len("BenchmarkJoin/"):], out)
		}
	}
	m := callsRE.FindSubmatch(out)
	if m == nil {
		t.Fatalf("could not count calls; output:\n%s", out)
	}
	plus, _ := strconv.Atoi(string(m[1]))
	builder, _ := strconv.Atoi(string(m[2]))
	if plus < 20 || builder < 20 {
		t.Errorf("with -benchtime=20x each sub-benchmark should call its function at least 20 times; got JoinPlus %d, JoinBuilder %d", plus, builder)
	}
}
