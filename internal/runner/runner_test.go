package runner

import (
	"context"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"
	"time"

	"github.com/maskedsyntax/golings/internal/catalog"
	"github.com/maskedsyntax/golings/internal/toolchain"
)

func TestParseTestJSON(t *testing.T) {
	stream := strings.Join([]string{
		`{"Action":"start","Package":"p"}`,
		`{"Action":"run","Package":"p","Test":"TestA"}`,
		`{"Action":"output","Package":"p","Test":"TestA","Output":"=== RUN   TestA\n"}`,
		`{"Action":"run","Package":"p","Test":"TestA/sub"}`,
		`{"Action":"output","Package":"p","Test":"TestA/sub","Output":"    a_test.go:9: got 1, want 2\n"}`,
		`{"Action":"output","Package":"p","Test":"TestA/sub","Output":"--- FAIL: TestA/sub (0.00s)\n"}`,
		`{"Action":"fail","Package":"p","Test":"TestA/sub"}`,
		`{"Action":"fail","Package":"p","Test":"TestA"}`,
		`{"Action":"run","Package":"p","Test":"TestB"}`,
		`{"Action":"pass","Package":"p","Test":"TestB"}`,
		`not json at all`,
		`{"Action":"fail","Package":"p"}`,
	}, "\n")
	r := parseTestJSON([]byte(stream))
	if r.BuildFailed || !r.PackageFailed {
		t.Errorf("BuildFailed=%v PackageFailed=%v", r.BuildFailed, r.PackageFailed)
	}
	leaves := leafTests(r.Tests)
	if len(leaves) != 2 || leaves[0].Name != "TestA/sub" || leaves[1].Name != "TestB" {
		t.Fatalf("leaves = %+v", leaves)
	}
	if leaves[0].Status != StatusFail || leaves[0].Output != "    a_test.go:9: got 1, want 2" {
		t.Errorf("TestA/sub = %+v", leaves[0])
	}
	if leaves[1].Status != StatusPass {
		t.Errorf("TestB = %+v", leaves[1])
	}
}

func TestParseTestJSONBuildFailure(t *testing.T) {
	stream := `{"ImportPath":"p","Action":"build-output","Output":"./a.go:3:2: undefined: x\n"}
{"ImportPath":"p","Action":"build-fail"}
{"Action":"fail","Package":"p","FailedBuild":"p"}`
	r := parseTestJSON([]byte(stream))
	if !r.BuildFailed || !strings.Contains(r.BuildOutput, "undefined: x") {
		t.Errorf("report = %+v", r)
	}
}

func TestTailBuffer(t *testing.T) {
	b := newTailBuffer(10)
	for range 100 {
		b.Write([]byte("0123456789"))
	}
	b.Write([]byte("abc"))
	if got := string(b.Bytes()); got != "3456789abc" || !b.Truncated() {
		t.Errorf("tail = %q, truncated=%v", got, b.Truncated())
	}
	small := newTailBuffer(10)
	small.Write([]byte("hi"))
	if string(small.Bytes()) != "hi" || small.Truncated() {
		t.Error("small buffer truncated")
	}
}

func TestDescribeMismatch(t *testing.T) {
	for _, tt := range []struct{ got, want, contains string }{
		{"", "Hello", "missing line 1"},
		{"Hello\nextra", "Hello", "extra line 2"},
		{"hello", "Hello", "line 1 differs"},
	} {
		if msg := describeMismatch(tt.got, tt.want); !strings.Contains(msg, tt.contains) {
			t.Errorf("describeMismatch(%q, %q) = %q", tt.got, tt.want, msg)
		}
	}
}

func TestWorst(t *testing.T) {
	if Worst(StatusPass, StatusUnsupported) != StatusUnsupported ||
		Worst(StatusFail, StatusTimeout) != StatusTimeout ||
		Worst(StatusRunnerError, StatusFail) != StatusRunnerError {
		t.Error("severity order wrong")
	}
}

// ---- integration tests with the real Go toolchain ---------------------------

func detect(t *testing.T) *toolchain.Info {
	t.Helper()
	if testing.Short() {
		t.Skip("skipping toolchain integration test in -short mode")
	}
	tc, err := toolchain.Detect(context.Background())
	if err != nil {
		t.Skipf("Go toolchain unavailable: %v", err)
	}
	return tc
}

func fakeExercise(checks ...catalog.CheckKind) *catalog.Exercise {
	return &catalog.Exercise{ID: "fake_01", StarterFiles: []string{"fake.go"}, Checks: checks, TimeoutSeconds: 10}
}

func run(t *testing.T, tc *toolchain.Info, ex *catalog.Exercise, src string, assets fstest.MapFS, timeout time.Duration) *ExerciseResult {
	t.Helper()
	return Run(context.Background(), Request{
		Exercise: ex, Files: map[string][]byte{"fake.go": []byte(src)},
		Assets: assets, Toolchain: tc, Timeout: timeout, DisplayDir: "exercises/fake_01",
	})
}

const okTest = `package fake
import "testing"
func TestF(t *testing.T) { if F() != 2 { t.Errorf("F() = %d, want 2", F()) } }
`

func TestRunPassFailCompile(t *testing.T) {
	tc := detect(t)
	assets := fstest.MapFS{"checks/fake_test.go": {Data: []byte(okTest)}}
	ex := fakeExercise(catalog.CheckUnit)

	if res := run(t, tc, ex, "package fake\nfunc F() int { return 2 }\n", assets, 0); res.Status != StatusPass {
		t.Errorf("passing code: %+v", res)
	}

	res := run(t, tc, ex, "package fake\nfunc F() int { return 1 }\n", assets, 0)
	if res.Status != StatusFail || res.Checks[0].Failure != FailureTest || len(res.Checks[0].Tests) != 1 ||
		!strings.Contains(res.Checks[0].Tests[0].Output, "F() = 1, want 2") {
		t.Errorf("failing code: %+v", res)
	}

	res = run(t, tc, ex, "package fake\nfunc F() int { return \"x\" }\n", assets, 0)
	c := res.Checks[0]
	if res.Status != StatusFail || c.Failure != FailureCompile || len(c.Diagnostics) == 0 ||
		!strings.HasPrefix(filepath.ToSlash(c.Diagnostics[0].File), "exercises/fake_01/") {
		t.Errorf("non-compiling code: %+v", res)
	}
}

func TestRunTimeoutOnInfiniteLoop(t *testing.T) {
	tc := detect(t)
	assets := fstest.MapFS{"checks/fake_test.go": {Data: []byte(okTest)}}
	start := time.Now()
	res := run(t, tc, fakeExercise(catalog.CheckUnit), "package fake\nfunc F() int { for {} }\n", assets, 2*time.Second)
	if res.Status != StatusTimeout || res.Checks[0].Failure != FailureTimeout {
		t.Errorf("result = %+v", res)
	}
	if d := time.Since(start); d > time.Minute {
		t.Errorf("timeout took %v", d)
	}
}

func TestRunPanickingCodeFailsCleanly(t *testing.T) {
	tc := detect(t)
	assets := fstest.MapFS{"checks/fake_test.go": {Data: []byte(okTest)}}
	res := run(t, tc, fakeExercise(catalog.CheckUnit), "package fake\nfunc F() int { panic(\"boom\") }\n", assets, 0)
	if res.Status != StatusFail || res.Checks[0].Failure != FailureTest {
		t.Errorf("result = %+v", res)
	}
}

func TestRunOutputCheck(t *testing.T) {
	tc := detect(t)
	assets := fstest.MapFS{"checks/expected_stdout.txt": {Data: []byte("hi\r\n")}}
	ex := fakeExercise(catalog.CheckOutput)
	ok := "package main\nimport \"fmt\"\nfunc main() { fmt.Println(\"hi\") }\n"
	if res := run(t, tc, ex, ok, assets, 0); res.Status != StatusPass {
		t.Errorf("matching output: %+v", res)
	}
	bad := "package main\nimport \"fmt\"\nfunc main() { fmt.Println(\"bye\") }\n"
	if res := run(t, tc, ex, bad, assets, 0); res.Status != StatusFail || res.Checks[0].Failure != FailureOutput || res.Checks[0].Actual != "bye" {
		t.Errorf("wrong output: %+v", res)
	}
	hang := "package main\nfunc main() { select {} }\n"
	if res := run(t, tc, ex, hang, assets, time.Second); res.Status != StatusTimeout && res.Status != StatusFail {
		// select{} with no goroutines is detected as a deadlock (exit 2);
		// either outcome must not hang or pass.
		t.Errorf("blocking program: %+v", res)
	}
	spin := "package main\nfunc main() { for {} }\n"
	if res := run(t, tc, ex, spin, assets, time.Second); res.Status != StatusTimeout {
		t.Errorf("spinning program: %+v", res)
	}
}

func TestRunRaceCheck(t *testing.T) {
	tc := detect(t)
	ex := fakeExercise(catalog.CheckRace)
	test := `package fake
import "testing"
func TestRacy(t *testing.T) { if Count(100) <= 0 { t.Fatal("no count") } }
`
	racy := `package fake
import "sync"
func Count(n int) int {
	var c int
	var wg sync.WaitGroup
	for range n { wg.Add(1); go func() { defer wg.Done(); c++ }() }
	wg.Wait()
	return c
}
`
	res := run(t, tc, ex, racy, fstest.MapFS{"checks/fake_test.go": {Data: []byte(test)}}, 30*time.Second)
	if !tc.RaceSupported {
		if res.Status != StatusUnsupported {
			t.Errorf("race unsupported: got %+v", res)
		}
		return
	}
	if res.Status != StatusFail || res.Checks[0].Failure != FailureRace {
		t.Errorf("racy code: %+v", res)
	}
}

func TestRunUnsupportedPlatformAndCollisions(t *testing.T) {
	tc := &toolchain.Info{GOOS: "plan9"}
	ex := fakeExercise(catalog.CheckUnit)
	ex.Platforms = []string{"linux"}
	if res := run(t, tc, ex, "", nil, 0); res.Status != StatusUnsupported || res.Passed() {
		t.Errorf("unsupported platform: %+v", res)
	}

	tc = &toolchain.Info{GOOS: "linux"}
	ex = fakeExercise(catalog.CheckUnit)
	assets := fstest.MapFS{"checks/fake.go": {Data: []byte("package fake")}}
	if res := run(t, tc, ex, "package fake", assets, 0); res.Status != StatusRunnerError {
		t.Errorf("collision: %+v", res)
	}
}

func TestRunMissingLearnerFile(t *testing.T) {
	tc := &toolchain.Info{GOOS: "linux"}
	res := Run(context.Background(), Request{Exercise: fakeExercise(catalog.CheckUnit), Files: map[string][]byte{}, Assets: fstest.MapFS{}, Toolchain: tc})
	if res.Status != StatusRunnerError {
		t.Errorf("missing file: %+v", res)
	}
}

func TestRunASTConstraint(t *testing.T) {
	tc := &toolchain.Info{GOOS: "linux"} // the ast check runs in-process
	ex := fakeExercise(catalog.CheckAST)
	ex.ASTRules = []string{"uses_defer:F"}

	res := run(t, tc, ex, "package fake\nfunc F() int { return 2 }\n", fstest.MapFS{}, 0)
	if res.Status != StatusFail || res.Checks[0].Failure != FailureConstraint || len(res.Checks[0].Tests) != 1 {
		t.Errorf("missing defer: %+v", res)
	}
	res = run(t, tc, ex, "package fake\nfunc F() int { defer func() {}(); return 2 }\n", fstest.MapFS{}, 0)
	if res.Status != StatusPass {
		t.Errorf("with defer: %+v", res)
	}
	res = run(t, tc, ex, "package fake\nfunc F( {\n", fstest.MapFS{}, 0)
	if res.Status != StatusFail || res.Checks[0].Failure != FailureCompile {
		t.Errorf("syntax error: %+v", res)
	}
}

func TestLoopbackBlocked(t *testing.T) {
	for _, tt := range []struct {
		out  string
		want bool
	}{
		{"panic: httptest: failed to listen on a port: listen tcp6 [::1]:0: socket: operation not permitted", true},
		{"server_test.go:12: listen tcp 127.0.0.1:0: bind: permission denied", true},
		{"listen tcp4 127.0.0.1:0: socket: An attempt was made to access a socket in a way forbidden by its access permissions.", true},
		{"listen tcp 127.0.0.1:8080: bind: address already in use", false},
		{"--- FAIL: TestEcho: got \"a\", want \"A\"", false},
	} {
		got := loopbackBlocked(testReport{}, []TestResult{{Output: tt.out}})
		if got != tt.want {
			t.Errorf("loopbackBlocked(%q) = %v, want %v", tt.out, got, tt.want)
		}
		if got := loopbackBlocked(testReport{PackageOutput: tt.out}, nil); got != tt.want {
			t.Errorf("loopbackBlocked(package output %q) = %v, want %v", tt.out, got, tt.want)
		}
	}
}

// Tools sometimes print absolute paths with symlinks resolved (macOS's
// /var -> /private/var); the staging directory must already be in that form
// or the diagnostics rewriter leaves a mangled prefix behind.
func TestStageDirIsCanonical(t *testing.T) {
	ex := &catalog.Exercise{ID: "x_01", StarterFiles: []string{"a.go"}}
	st, err := newStage(ex, map[string][]byte{"a.go": []byte("package x\n")}, fstest.MapFS{})
	if err != nil {
		t.Fatal(err)
	}
	defer st.cleanup()
	real, err := filepath.EvalSymlinks(st.dir)
	if err != nil {
		t.Fatal(err)
	}
	if real != st.dir {
		t.Fatalf("staging dir %q is not canonical (%q)", st.dir, real)
	}
}
