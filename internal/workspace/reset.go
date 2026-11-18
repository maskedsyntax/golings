package workspace

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/maskedsyntax/golings/internal/catalog"
)

// ResetPlan describes what Reset would change.
type ResetPlan struct {
	Changed []string // learner files that differ from the starter
	Missing []string // learner files that are absent
}

// Empty reports whether the exercise already matches its starter.
func (p ResetPlan) Empty() bool { return len(p.Changed) == 0 && len(p.Missing) == 0 }

// PlanReset compares the workspace files with the original starters.
func (w *Workspace) PlanReset(cat *catalog.Catalog, ex *catalog.Exercise) (ResetPlan, error) {
	var plan ResetPlan
	starters, err := starterFiles(cat, ex)
	if err != nil {
		return plan, err
	}
	dir := w.ExerciseDir(ex)
	for _, name := range ex.StarterFiles {
		fi, err := os.Lstat(filepath.Join(dir, name))
		switch {
		case errors.Is(err, fs.ErrNotExist):
			plan.Missing = append(plan.Missing, name)
			continue
		case err != nil:
			return plan, err
		case !fi.Mode().IsRegular():
			plan.Changed = append(plan.Changed, name)
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			return plan, err
		}
		if !bytes.Equal(data, starters[name]) {
			plan.Changed = append(plan.Changed, name)
		}
	}
	return plan, nil
}

// Reset restores the exercise's starter files. Unless backup is false, the
// current versions of changed files are first copied to
// .golings/backups/<id>/<timestamp>/. It returns the backup directory ("" if
// none was made).
func (w *Workspace) Reset(cat *catalog.Catalog, ex *catalog.Exercise, backup bool) (string, error) {
	plan, err := w.PlanReset(cat, ex)
	if err != nil || plan.Empty() {
		return "", err
	}
	starters, err := starterFiles(cat, ex)
	if err != nil {
		return "", err
	}
	dir := w.ExerciseDir(ex)

	backupDir := ""
	if backup && len(plan.Changed) > 0 {
		backupDir = filepath.Join(w.MetaPath(), "backups", ex.ID, time.Now().UTC().Format("20060102T150405.000Z"))
		if err := os.MkdirAll(backupDir, 0o755); err != nil {
			return "", err
		}
		for _, name := range plan.Changed {
			src := filepath.Join(dir, name)
			fi, err := os.Lstat(src)
			if err != nil || !fi.Mode().IsRegular() {
				continue // symlinks and other oddities are not followed
			}
			data, err := os.ReadFile(src)
			if err != nil {
				return "", fmt.Errorf("backing up %s: %w", name, err)
			}
			if err := os.WriteFile(filepath.Join(backupDir, name), data, 0o644); err != nil {
				return "", fmt.Errorf("backing up %s: %w", name, err)
			}
		}
	}

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return backupDir, err
	}
	for _, name := range append(append([]string{}, plan.Changed...), plan.Missing...) {
		target := filepath.Join(dir, name)
		// Replace a symlink instead of writing through it.
		if fi, err := os.Lstat(target); err == nil && !fi.Mode().IsRegular() {
			if err := os.Remove(target); err != nil {
				return backupDir, err
			}
		}
		if err := replaceFile(target, starters[name]); err != nil {
			return backupDir, fmt.Errorf("restoring %s: %w", name, err)
		}
	}
	return backupDir, nil
}

func starterFiles(cat *catalog.Catalog, ex *catalog.Exercise) (map[string][]byte, error) {
	return assetFiles(cat, ex, "starter")
}

// SolutionFiles returns the reference solution.
func SolutionFiles(cat *catalog.Catalog, ex *catalog.Exercise) (map[string][]byte, error) {
	return assetFiles(cat, ex, "solution")
}

func assetFiles(cat *catalog.Catalog, ex *catalog.Exercise, dir string) (map[string][]byte, error) {
	assets, err := cat.Assets(ex)
	if err != nil {
		return nil, err
	}
	out := map[string][]byte{}
	for _, name := range ex.StarterFiles {
		data, err := fs.ReadFile(assets, path.Join(dir, name))
		if err != nil {
			return nil, err
		}
		out[name] = data
	}
	return out, nil
}

// replaceFile atomically replaces name with data.
func replaceFile(name string, data []byte) error {
	f, err := os.CreateTemp(filepath.Dir(name), "."+filepath.Base(name)+".tmp-*")
	if err != nil {
		return err
	}
	tmp := f.Name()
	if _, err := f.Write(data); err != nil {
		f.Close()
		os.Remove(tmp)
		return err
	}
	if err := f.Close(); err != nil {
		os.Remove(tmp)
		return err
	}
	if err := os.Chmod(tmp, 0o644); err != nil {
		os.Remove(tmp)
		return err
	}
	if err := os.Rename(tmp, name); err != nil {
		os.Remove(tmp)
		return err
	}
	return nil
}
