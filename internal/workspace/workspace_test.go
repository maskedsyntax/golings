package workspace_test

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/maskedsyntax/golings/curriculum"
	"github.com/maskedsyntax/golings/internal/catalog"
	"github.com/maskedsyntax/golings/internal/workspace"
)

func loadCatalog(t *testing.T) *catalog.Catalog {
	t.Helper()
	cat, err := catalog.Load(curriculum.FS)
	if err != nil {
		t.Fatal(err)
	}
	return cat
}

func TestInitMaterializesStartersOnly(t *testing.T) {
	cat := loadCatalog(t)
	dest := filepath.Join(t.TempDir(), "ws")
	ws, err := workspace.Init(dest, cat, workspace.InitOptions{CLIVersion: "test"})
	if err != nil {
		t.Fatal(err)
	}
	for _, ex := range cat.Exercises() {
		dir := ws.ExerciseDir(ex)
		for _, f := range ex.StarterFiles {
			if _, err := os.Stat(filepath.Join(dir, f)); err != nil {
				t.Errorf("%s: %v", ex.ID, err)
			}
		}
	}
	// Canonical checks and solutions never reach the workspace.
	filepath.WalkDir(dest, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		name := d.Name()
		if strings.HasSuffix(name, "_test.go") || name == "solution" || name == "checks" || name == "mutants" {
			t.Errorf("workspace contains %s", p)
		}
		return nil
	})
	for _, f := range []string{"go.mod", "README.md", ".golings/workspace.json"} {
		if _, err := os.Stat(filepath.Join(dest, f)); err != nil {
			t.Errorf("missing %s: %v", f, err)
		}
	}
	reopened, err := workspace.Open(dest)
	if err != nil || reopened.Meta.CurriculumVersion != cat.Version() {
		t.Errorf("Open = %+v, %v", reopened, err)
	}
}

func TestInitRefusesUnsafeDestinations(t *testing.T) {
	cat := loadCatalog(t)

	nonEmpty := t.TempDir()
	os.WriteFile(filepath.Join(nonEmpty, "notes.txt"), []byte("mine"), 0o644)
	if _, err := workspace.Init(nonEmpty, cat, workspace.InitOptions{}); !errors.Is(err, workspace.ErrNotEmpty) {
		t.Errorf("non-empty: err = %v, want ErrNotEmpty", err)
	}
	if _, err := workspace.Init(nonEmpty, cat, workspace.InitOptions{AllowNonEmpty: true}); err != nil {
		t.Errorf("non-empty with AllowNonEmpty: %v", err)
	}
	if data, _ := os.ReadFile(filepath.Join(nonEmpty, "notes.txt")); string(data) != "mine" {
		t.Error("init modified an existing file")
	}
	if _, err := workspace.Init(nonEmpty, cat, workspace.InitOptions{AllowNonEmpty: true}); !errors.Is(err, workspace.ErrAlreadyWorkspace) {
		t.Errorf("second init: err = %v, want ErrAlreadyWorkspace", err)
	}

	collide := t.TempDir()
	os.WriteFile(filepath.Join(collide, "go.mod"), []byte("module mine\n"), 0o644)
	if _, err := workspace.Init(collide, cat, workspace.InitOptions{AllowNonEmpty: true}); err == nil {
		t.Error("init over an existing go.mod succeeded")
	}
	if data, _ := os.ReadFile(filepath.Join(collide, "go.mod")); string(data) != "module mine\n" {
		t.Error("init overwrote go.mod")
	}
}

func TestDiscoverWalksUp(t *testing.T) {
	cat := loadCatalog(t)
	dest := filepath.Join(t.TempDir(), "ws")
	ws, err := workspace.Init(dest, cat, workspace.InitOptions{})
	if err != nil {
		t.Fatal(err)
	}
	found, err := workspace.Discover(ws.ExerciseDir(cat.Exercises()[0]))
	if err != nil || found.Root != ws.Root {
		t.Errorf("Discover = %v, %v; want %s", found, err, ws.Root)
	}
	if _, err := workspace.Discover(t.TempDir()); !errors.Is(err, workspace.ErrNotWorkspace) {
		t.Errorf("Discover outside: err = %v", err)
	}
}

func TestReadLearnerFiles(t *testing.T) {
	cat := loadCatalog(t)
	ws, err := workspace.Init(filepath.Join(t.TempDir(), "ws"), cat, workspace.InitOptions{})
	if err != nil {
		t.Fatal(err)
	}
	ex := cat.Exercises()[0]
	files, err := ws.ReadLearnerFiles(ex)
	if err != nil || len(files) != len(ex.StarterFiles) {
		t.Fatalf("ReadLearnerFiles = %v, %v", files, err)
	}
	fp := workspace.Fingerprint(files)
	if fp != workspace.Fingerprint(files) || !strings.HasPrefix(fp, "sha256:") {
		t.Errorf("fingerprint unstable: %s", fp)
	}

	target := filepath.Join(ws.ExerciseDir(ex), ex.StarterFiles[0])
	os.Remove(target)
	var missing *workspace.MissingFilesError
	if _, err := ws.ReadLearnerFiles(ex); !errors.As(err, &missing) {
		t.Errorf("missing file: err = %v", err)
	}

	if runtime.GOOS == "windows" {
		t.Skip("symlinks need privileges on Windows")
	}
	outside := filepath.Join(t.TempDir(), "secret.go")
	os.WriteFile(outside, []byte("package main"), 0o644)
	if err := os.Symlink(outside, target); err != nil {
		t.Fatal(err)
	}
	if _, err := ws.ReadLearnerFiles(ex); err == nil || !strings.Contains(err.Error(), "symlink") {
		t.Errorf("symlinked starter: err = %v, want a symlink rejection", err)
	}
}

func TestFingerprintDependsOnNamesAndContent(t *testing.T) {
	a := workspace.Fingerprint(map[string][]byte{"a.go": []byte("x"), "b.go": []byte("y")})
	b := workspace.Fingerprint(map[string][]byte{"a.go": []byte("xb"), "b.go": nil})
	c := workspace.Fingerprint(map[string][]byte{"a.go": []byte("x"), "b.go": []byte("z")})
	if a == b || a == c {
		t.Error("fingerprints collide")
	}
}

func TestResetRestoresWithBackup(t *testing.T) {
	cat := loadCatalog(t)
	ws, err := workspace.Init(filepath.Join(t.TempDir(), "ws"), cat, workspace.InitOptions{})
	if err != nil {
		t.Fatal(err)
	}
	ex := cat.Exercises()[0]
	file := filepath.Join(ws.ExerciseDir(ex), ex.StarterFiles[0])
	original, _ := os.ReadFile(file)

	if plan, err := ws.PlanReset(cat, ex); err != nil || !plan.Empty() {
		t.Fatalf("fresh exercise plan = %+v, %v", plan, err)
	}
	if dir, err := ws.Reset(cat, ex, true); err != nil || dir != "" {
		t.Fatalf("no-op reset = %q, %v", dir, err)
	}

	os.WriteFile(file, []byte("my work"), 0o644)
	plan, err := ws.PlanReset(cat, ex)
	if err != nil || len(plan.Changed) != 1 {
		t.Fatalf("plan = %+v, %v", plan, err)
	}
	backup, err := ws.Reset(cat, ex, true)
	if err != nil || backup == "" {
		t.Fatalf("reset = %q, %v", backup, err)
	}
	if data, _ := os.ReadFile(file); !bytes.Equal(data, original) {
		t.Error("file not restored")
	}
	if data, _ := os.ReadFile(filepath.Join(backup, ex.StarterFiles[0])); string(data) != "my work" {
		t.Errorf("backup = %q", data)
	}

	os.Remove(file)
	if _, err := ws.Reset(cat, ex, false); err != nil {
		t.Fatal(err)
	}
	if data, _ := os.ReadFile(file); !bytes.Equal(data, original) {
		t.Error("missing file not restored")
	}
}
