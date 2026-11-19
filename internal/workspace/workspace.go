// Package workspace creates and reads learner workspaces: the directory that
// holds editable starter files and local progress. Canonical checks and
// solutions never live here.
package workspace

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/maskedsyntax/golings/internal/catalog"
)

// FormatVersion is the on-disk workspace layout version.
const FormatVersion = 1

const (
	// MetaDir holds golings' own files inside a workspace.
	MetaDir      = ".golings"
	metaFile     = "workspace.json"
	exercisesDir = "exercises"
	// maxSourceSize bounds how much of a learner file is read.
	maxSourceSize = 1 << 20
)

// Meta is .golings/workspace.json.
type Meta struct {
	WorkspaceFormatVersion int       `json:"workspace_format_version"`
	CurriculumVersion      string    `json:"curriculum_version"`
	CreatedAt              time.Time `json:"created_at"`
	CreatedBy              string    `json:"created_by"`
	// UpgradedAt is when `golings upgrade` last brought the workspace to
	// CurriculumVersion.
	UpgradedAt *time.Time `json:"upgraded_at,omitempty"`
}

// Workspace is an opened learner workspace.
type Workspace struct {
	Root string // absolute path
	Meta Meta
}

// ErrNotWorkspace means no workspace was found.
var ErrNotWorkspace = errors.New("not inside a golings workspace (run `golings init` to create one, or pass --workspace)")

// Open opens the workspace rooted at root.
func Open(root string) (*Workspace, error) {
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(filepath.Join(abs, MetaDir, metaFile))
	if errors.Is(err, fs.ErrNotExist) {
		return nil, fmt.Errorf("%s: %w", abs, ErrNotWorkspace)
	}
	if err != nil {
		return nil, err
	}
	w := &Workspace{Root: abs}
	if err := json.Unmarshal(data, &w.Meta); err != nil {
		return nil, fmt.Errorf("%s is corrupt: %w", filepath.Join(MetaDir, metaFile), err)
	}
	if w.Meta.WorkspaceFormatVersion != FormatVersion {
		return nil, fmt.Errorf("workspace format version %d is not supported by this golings (want %d)", w.Meta.WorkspaceFormatVersion, FormatVersion)
	}
	return w, nil
}

// Discover walks up from start to find the nearest workspace.
func Discover(start string) (*Workspace, error) {
	dir, err := filepath.Abs(start)
	if err != nil {
		return nil, err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, MetaDir, metaFile)); err == nil {
			return Open(dir)
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return nil, ErrNotWorkspace
		}
		dir = parent
	}
}

// MetaPath returns the absolute path of the .golings directory.
func (w *Workspace) MetaPath() string { return filepath.Join(w.Root, MetaDir) }

// RelExerciseDir is the exercise directory relative to the workspace root,
// for display.
func (w *Workspace) RelExerciseDir(ex *catalog.Exercise) string {
	return filepath.Join(exercisesDir, ex.Module, ex.ID)
}

// ExerciseDir is the absolute exercise directory.
func (w *Workspace) ExerciseDir(ex *catalog.Exercise) string {
	return filepath.Join(w.Root, w.RelExerciseDir(ex))
}

// MissingFilesError lists learner files that are absent.
type MissingFilesError struct {
	Exercise string
	Files    []string
}

func (e *MissingFilesError) Error() string {
	return fmt.Sprintf("exercise %s is missing %v (restore it with `golings reset %s`)", e.Exercise, e.Files, e.Exercise)
}

// ReadLearnerFiles reads the exercise's starter files from the workspace.
// Reads go through os.Root, so symlinks or paths escaping the exercise
// directory are rejected, and symlinked files are refused outright.
func (w *Workspace) ReadLearnerFiles(ex *catalog.Exercise) (map[string][]byte, error) {
	root, err := os.OpenRoot(w.ExerciseDir(ex))
	if errors.Is(err, fs.ErrNotExist) {
		return nil, &MissingFilesError{Exercise: ex.ID, Files: ex.StarterFiles}
	}
	if err != nil {
		return nil, err
	}
	defer root.Close()

	files := map[string][]byte{}
	var missing []string
	for _, name := range ex.StarterFiles {
		fi, err := root.Lstat(name)
		if errors.Is(err, fs.ErrNotExist) {
			missing = append(missing, name)
			continue
		}
		if err != nil {
			return nil, err
		}
		if !fi.Mode().IsRegular() {
			return nil, fmt.Errorf("%s must be a regular file (symlinks are not allowed)", filepath.Join(w.RelExerciseDir(ex), name))
		}
		if fi.Size() > maxSourceSize {
			return nil, fmt.Errorf("%s is larger than %d bytes", filepath.Join(w.RelExerciseDir(ex), name), maxSourceSize)
		}
		f, err := root.Open(name)
		if err != nil {
			return nil, err
		}
		data, err := io.ReadAll(io.LimitReader(f, maxSourceSize+1))
		f.Close()
		if err != nil {
			return nil, err
		}
		files[name] = data
	}
	if len(missing) > 0 {
		return nil, &MissingFilesError{Exercise: ex.ID, Files: missing}
	}
	return files, nil
}

// Fingerprint hashes a set of files independent of map order and
// modification times.
func Fingerprint(files map[string][]byte) string {
	names := make([]string, 0, len(files))
	for n := range files {
		names = append(names, n)
	}
	sort.Strings(names)
	h := sha256.New()
	for _, n := range names {
		fmt.Fprintf(h, "%d:%s\n%d:", len(n), n, len(files[n]))
		h.Write(files[n])
	}
	return "sha256:" + hex.EncodeToString(h.Sum(nil))
}
