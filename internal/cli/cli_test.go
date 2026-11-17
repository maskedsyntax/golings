package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/maskedsyntax/golings/curriculum"
	"github.com/maskedsyntax/golings/internal/exitcode"
	"github.com/maskedsyntax/golings/internal/toolchain"
)

type result struct {
	code           int
	stdout, stderr string
}

func golings(t *testing.T, args ...string) result {
	t.Helper()
	var out, errOut bytes.Buffer
	code := run(context.Background(), args, strings.NewReader(""), &out, &errOut)
	return result{code, out.String(), errOut.String()}
}

func (r result) expect(t *testing.T, code int, contains ...string) {
	t.Helper()
	if r.code != code {
		t.Fatalf("exit code %d, want %d\nstdout:\n%s\nstderr:\n%s", r.code, code, r.stdout, r.stderr)
	}
	for _, c := range contains {
		if !strings.Contains(r.stdout+r.stderr, c) {
			t.Errorf("output does not contain %q\nstdout:\n%s\nstderr:\n%s", c, r.stdout, r.stderr)
		}
	}
}

func copySolution(t *testing.T, ws, module, id, file string) {
	t.Helper()
	data, err := fs.ReadFile(curriculum.FS, "_exercises/"+module+"/"+id+"/solution/"+file)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(ws, "exercises", module, id, file), data, 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestCommandsWithoutToolchain(t *testing.T) {
	golings(t, "help").expect(t, exitcode.OK, "Commands:", "verify")
	golings(t, "--help").expect(t, exitcode.OK, "Commands:")
	golings(t, "help", "verify").expect(t, exitcode.OK, "--all")
	golings(t, "frobnicate").expect(t, exitcode.Usage, "unknown command")
	golings(t, "list", "--bogus-flag").expect(t, exitcode.Usage)

	r := golings(t, "version", "--json")
	r.expect(t, exitcode.OK)
	var v map[string]any
	if err := json.Unmarshal([]byte(r.stdout), &v); err != nil || v["curriculum"] == nil {
		t.Errorf("version --json = %q (%v)", r.stdout, err)
	}

	golings(t, "--workspace", t.TempDir(), "list").expect(t, exitcode.Workspace, "not inside a golings workspace")
}

func TestInitRefusesNonEmpty(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "keep.txt"), []byte("x"), 0o644)
	golings(t, "init", dir).expect(t, exitcode.Workspace, "not empty", "--force-into-nonempty")
	golings(t, "init", "--force-into-nonempty", dir).expect(t, exitcode.OK)
	golings(t, "init", dir).expect(t, exitcode.Workspace, "already a golings workspace")
}

// TestLearnerJourney is the Milestone A end-to-end path: init, fail, fix,
// pass, move on, "restart" (a fresh invocation) and resume.
func TestLearnerJourney(t *testing.T) {
	if testing.Short() {
		t.Skip("end-to-end test runs the Go toolchain")
	}
	if _, err := toolchain.Detect(context.Background()); err != nil {
		t.Skipf("Go toolchain unavailable: %v", err)
	}
	ws := filepath.Join(t.TempDir(), "my workspace")
	golings(t, "init", ws).expect(t, exitcode.OK, "Created a workspace")
	w := []string{"--workspace", ws, "--no-color"}
	cmd := func(args ...string) result { return golings(t, append(append([]string{}, w...), args...)...) }

	// The first exercise fails until it is fixed.
	cmd().expect(t, exitcode.VerifyFailed, "hello_world_01", "Output mismatch", "Not passing yet")

	// Unknown and misspelled ids are usage errors with suggestions.
	cmd("verify", "hello_wrld_01").expect(t, exitcode.Usage, "did you mean", "hello_world_01")

	// A compile error points at the learner's file.
	cmd("verify", "values").expect(t, exitcode.VerifyFailed, "Compiler error",
		filepath.Join("exercises", "01_fundamentals", "values_01", "values.go")+":")

	copySolution(t, ws, "01_fundamentals", "hello_world_01", "main.go")
	cmd("verify").expect(t, exitcode.OK, "Hello, Golings!", "Exercise complete", "values_01")

	cmd("next").expect(t, exitcode.OK, "values_01")

	// A fresh invocation resumes at the selected exercise.
	cmd().expect(t, exitcode.VerifyFailed, "Exercise 2/5", "values_01")

	// JSON output is parseable and free of ANSI escapes.
	r := cmd("progress", "--json")
	r.expect(t, exitcode.OK)
	var sum struct{ Total, Completed int }
	if err := json.Unmarshal([]byte(r.stdout), &sum); err != nil || sum.Completed != 1 || sum.Total != 5 {
		t.Errorf("progress --json = %s (%v)", r.stdout, err)
	}
	r = golings(t, "--workspace", ws, "list", "--json")
	r.expect(t, exitcode.OK)
	if strings.Contains(r.stdout, "\x1b[") {
		t.Error("JSON output contains ANSI escapes")
	}
	var rows []struct{ ID, Status string }
	if err := json.Unmarshal([]byte(r.stdout), &rows); err != nil || len(rows) != 5 || rows[0].Status != "completed" || rows[1].Status != "in_progress" {
		t.Errorf("list --json = %s (%v)", r.stdout, err)
	}

	r = cmd("verify", "variables_01", "--json")
	r.expect(t, exitcode.VerifyFailed)
	var vr struct {
		Result struct {
			Status string
			Checks []struct{ Failure string }
		}
	}
	if err := json.Unmarshal([]byte(r.stdout), &vr); err != nil || vr.Result.Status != "fail" || vr.Result.Checks[0].Failure != "compile" {
		t.Errorf("verify --json = %s (%v)", r.stdout, err)
	}

	// Editing a completed exercise marks it modified, not lost.
	f := filepath.Join(ws, "exercises", "01_fundamentals", "hello_world_01", "main.go")
	data, _ := os.ReadFile(f)
	os.WriteFile(f, append(data, []byte("\n// tweak\n")...), 0o644)
	cmd("list", "--status", "modified").expect(t, exitcode.OK, "hello_world_01")

	// Deleting learner files is reported, never counted as complete.
	os.Remove(f)
	cmd("list").expect(t, exitcode.OK, "missing")
	cmd("verify", "hello_world_01").expect(t, exitcode.Workspace, "missing")

	// --all reports every exercise and fails overall.
	cmd("verify", "--all").expect(t, exitcode.VerifyFailed, "of 5 exercises pass")
}
