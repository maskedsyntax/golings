package cli

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/maskedsyntax/golings/internal/exitcode"
	"github.com/maskedsyntax/golings/internal/toolchain"
)

// syncBuffer is a goroutine-safe bytes.Buffer.
type syncBuffer struct {
	mu sync.Mutex
	b  bytes.Buffer
}

func (s *syncBuffer) Write(p []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.b.Write(p)
}

func (s *syncBuffer) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.b.String()
}

// waitFor waits until out contains want after offset, and returns the new
// offset just past the match.
func waitFor(t *testing.T, out *syncBuffer, offset int, want string) int {
	t.Helper()
	deadline := time.Now().Add(90 * time.Second)
	for time.Now().Before(deadline) {
		s := out.String()
		if i := strings.Index(s[offset:], want); i >= 0 {
			return offset + i + len(want)
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Fatalf("timed out waiting for %q; output so far:\n%s", want, out.String()[offset:])
	return 0
}

// TestWatchLineMode drives watch mode through a pipe (the non-TTY fallback):
// initial check, re-check on save, hint, next exercise, quit.
func TestWatchLineMode(t *testing.T) {
	if testing.Short() {
		t.Skip("watch test runs the Go toolchain")
	}
	if _, err := toolchain.Detect(context.Background()); err != nil {
		t.Skipf("Go toolchain unavailable: %v", err)
	}
	ws := filepath.Join(t.TempDir(), "ws")
	golings(t, "init", ws).expect(t, exitcode.OK)

	inR, inW := io.Pipe()
	out := &syncBuffer{}
	done := make(chan int, 1)
	go func() {
		done <- run(context.Background(), []string{"--workspace", ws, "--no-color"}, inR, out, out)
	}()
	send := func(s string) { io.WriteString(inW, s+"\n") }

	off := waitFor(t, out, 0, "hello_world_01")
	off = waitFor(t, out, off, "Not passing yet")

	copySolution(t, ws, "01_fundamentals", "hello_world_01", "main.go")
	off = waitFor(t, out, off, "change detected")
	off = waitFor(t, out, off, "Exercise complete! Press n")

	send("h")
	off = waitFor(t, out, off, "Hint 1/3 for hello_world_01")

	send("n")
	off = waitFor(t, out, off, "hello_world_02")
	off = waitFor(t, out, off, "Compiler error")

	// An edit that doesn't change content does not trigger a re-check.
	f := filepath.Join(ws, "exercises", "01_fundamentals", "hello_world_02", "main.go")
	data, _ := os.ReadFile(f)
	os.WriteFile(f, data, 0o644)
	time.Sleep(600 * time.Millisecond)
	if strings.Contains(out.String()[off:], "change detected") {
		t.Error("identical save triggered a re-check")
	}

	send("q")
	select {
	case code := <-done:
		if code != exitcode.OK {
			t.Errorf("exit code %d", code)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("watch did not quit")
	}
	if !strings.Contains(golings(t, "--workspace", ws, "progress").stdout, fmt.Sprintf("1/%d", curriculumSize(t))) {
		t.Error("progress not recorded by watch mode")
	}
}

func TestWatchRejectsJSON(t *testing.T) {
	ws := filepath.Join(t.TempDir(), "ws")
	golings(t, "init", ws).expect(t, exitcode.OK)
	golings(t, "--workspace", ws, "--json").expect(t, exitcode.Usage, "verify --json")
}

func TestHintResetSolution(t *testing.T) {
	ws := filepath.Join(t.TempDir(), "ws")
	golings(t, "init", ws).expect(t, exitcode.OK)
	w := func(args ...string) result {
		return golings(t, append([]string{"--workspace", ws, "--no-color"}, args...)...)
	}

	w("hint").expect(t, exitcode.OK, "Hint 1/3 for hello_world_01")
	w("hint").expect(t, exitcode.OK, "Hint 2/3")
	w("hint").expect(t, exitcode.OK, "Hint 3/3", "last hint")
	w("hint").expect(t, exitcode.OK, "Hint 3/3")
	w("hint", "values_01", "--level", "2").expect(t, exitcode.OK, "Hint 2/3 for values_01")
	w("hint", "values_01", "--level", "9").expect(t, exitcode.Usage, "between 1 and 3")

	// Solutions require explicit consent when not interactive.
	w("solution", "values_01").expect(t, exitcode.Usage, "--show")
	r := w("solution", "values_01", "--show")
	r.expect(t, exitcode.OK, "float64(sum) / float64(count)", "not changed")
	outDir := filepath.Join(t.TempDir(), "sol")
	w("solution", "values_01", "--show", "--out", outDir).expect(t, exitcode.OK, "Wrote the solution")
	if _, err := os.Stat(filepath.Join(outDir, "values.go")); err != nil {
		t.Error(err)
	}
	w("solution", "values_01", "--show", "--out", outDir).expect(t, exitcode.Usage, "not empty")

	// Reset needs --force when not interactive, and keeps a backup.
	f := filepath.Join(ws, "exercises", "01_fundamentals", "values_01", "values.go")
	original, _ := os.ReadFile(f)
	w("reset", "values_01").expect(t, exitcode.OK, "nothing to reset")
	os.WriteFile(f, []byte("package values // my attempt\n"), 0o644)
	w("reset", "values_01").expect(t, exitcode.Usage, "--force")
	if data, _ := os.ReadFile(f); string(data) != "package values // my attempt\n" {
		t.Fatal("reset without consent changed the file")
	}
	w("reset", "values_01", "--force").expect(t, exitcode.OK, "Restored", "backups")
	if data, _ := os.ReadFile(f); !bytes.Equal(data, original) {
		t.Error("reset did not restore the starter")
	}
	backups, _ := filepath.Glob(filepath.Join(ws, ".golings", "backups", "values_01", "*", "values.go"))
	if len(backups) != 1 {
		t.Errorf("backups = %v", backups)
	}
}
