package runner

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/maskedsyntax/golings/internal/catalog"
)

// StagingModulePath is the module path of every staged exercise.
const StagingModulePath = "golings.local/exercise"

// stagingGoMod pins the language version so staged exercises behave the same
// under every supported toolchain.
const stagingGoMod = "module " + StagingModulePath + "\n\ngo 1.24\n"

// stage is a temporary directory, outside the learner workspace, that holds
// the learner's source, fixtures and canonical checks.
type stage struct {
	dir string
}

// newStage builds the staging directory. files are the learner's sources by
// base name; assets is the exercise asset FS (checks/, fixtures/).
func newStage(ex *catalog.Exercise, files map[string][]byte, assets fs.FS) (st *stage, err error) {
	dir, err := os.MkdirTemp("", "golings-*")
	if err != nil {
		return nil, fmt.Errorf("creating staging directory: %w", err)
	}
	// Resolve symlinks (macOS's /var -> /private/var, Windows 8.3 short
	// names) so every path the go tools print uses the same form the
	// diagnostics rewriter strips.
	if real, err := filepath.EvalSymlinks(dir); err == nil {
		dir = real
	}
	st = &stage{dir: dir}
	defer func() {
		if err != nil {
			st.cleanup()
			st = nil
		}
	}()

	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte(stagingGoMod), 0o644); err != nil {
		return nil, err
	}
	written := map[string]bool{"go.mod": true}

	for _, name := range ex.StarterFiles {
		data, ok := files[name]
		if !ok {
			return nil, fmt.Errorf("missing learner file %s", name)
		}
		if err := st.write(name, data, written); err != nil {
			return nil, err
		}
	}

	for _, rel := range ex.FixtureFiles {
		data, err := fs.ReadFile(assets, path.Join("fixtures", rel))
		if err != nil {
			return nil, fmt.Errorf("reading fixture %s: %w", rel, err)
		}
		if err := st.write(rel, data, written); err != nil {
			return nil, err
		}
	}

	// Canonical checks: every .go file in checks/. Other files there (such
	// as expected_stdout.txt) are read directly from assets, never staged.
	entries, err := fs.ReadDir(assets, "checks")
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, fmt.Errorf("reading checks: %w", err)
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
			continue
		}
		data, err := fs.ReadFile(assets, path.Join("checks", e.Name()))
		if err != nil {
			return nil, err
		}
		if err := st.write(e.Name(), data, written); err != nil {
			return nil, err
		}
	}
	return st, nil
}

// write creates rel (a slash-separated path validated by the catalog) inside
// the stage. A name collision is a curriculum bug, so it is an error rather
// than a silent overwrite.
func (s *stage) write(rel string, data []byte, written map[string]bool) error {
	if !fs.ValidPath(rel) {
		return fmt.Errorf("invalid staged path %q", rel)
	}
	if written[rel] {
		return fmt.Errorf("staging collision: %s is provided twice", rel)
	}
	written[rel] = true
	full := filepath.Join(s.dir, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		return err
	}
	return os.WriteFile(full, data, 0o644)
}

func (s *stage) cleanup() {
	if s != nil && s.dir != "" {
		_ = os.RemoveAll(s.dir)
	}
}
