package workspace

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/maskedsyntax/golings/internal/catalog"
)

// Move relocates an exercise directory: the exercise was renamed, or moved
// to another module. The learner's files move with it, unchanged.
type Move struct {
	ID    string `json:"id"`     // current exercise ID
	OldID string `json:"old_id"` // ID on disk (equal to ID for a module move)
	From  string `json:"from"`   // workspace-relative directories
	To    string `json:"to"`
}

// UpgradePlan describes what Upgrade would change. It never includes
// rewriting a learner file: only exercise READMEs, fixture copies and the
// workspace README, which golings owns, are ever overwritten.
type UpgradePlan struct {
	FromVersion string `json:"from_version"`
	ToVersion   string `json:"to_version"`
	// Added are new exercises whose starters will be created.
	Added []string `json:"added"`
	// Moved are exercise directories that move with the learner's files.
	Moved []Move `json:"moved"`
	// Retired are directories of exercises that left the curriculum. They
	// are kept on disk, untouched.
	Retired []string `json:"retired"`
	// Refreshed are exercises whose README or fixture copies will be updated.
	Refreshed []string `json:"refreshed"`
	// Conflicts are problems that make upgrading unsafe; nothing is applied
	// while there are any.
	Conflicts []string `json:"conflicts"`
}

// Empty reports whether the plan changes nothing on disk.
func (p UpgradePlan) Empty() bool {
	return len(p.Added) == 0 && len(p.Moved) == 0 && len(p.Refreshed) == 0
}

// ErrWorkspaceNewer means the workspace was made by a newer curriculum than
// this golings ships; upgrading "down" could lose work, so it is refused.
var ErrWorkspaceNewer = errors.New("this workspace was created by a newer golings")

// PlanUpgrade compares the workspace with the catalog.
func (w *Workspace) PlanUpgrade(cat *catalog.Catalog) (UpgradePlan, error) {
	plan := UpgradePlan{
		FromVersion: w.Meta.CurriculumVersion, ToVersion: cat.Version(),
		Added: []string{}, Moved: []Move{}, Retired: []string{}, Refreshed: []string{}, Conflicts: []string{},
	}
	if CompareVersions(plan.FromVersion, plan.ToVersion) > 0 {
		return plan, fmt.Errorf("%w (curriculum %s; this one ships %s): install the newer golings instead",
			ErrWorkspaceNewer, plan.FromVersion, plan.ToVersion)
	}

	onDisk, dupes, err := w.scanExerciseDirs()
	if err != nil {
		return plan, err
	}
	for _, id := range dupes {
		plan.Conflicts = append(plan.Conflicts, fmt.Sprintf("exercise %s exists in more than one module directory", id))
	}
	renames := cat.Renames()
	// Reverse index: current ID -> old IDs that were renamed to it.
	oldIDs := map[string][]string{}
	for old, cur := range renames {
		if cur != "" {
			oldIDs[cur] = append(oldIDs[cur], old)
		}
	}
	claimed := map[string]bool{}

	for _, ex := range cat.Exercises() {
		want := w.RelExerciseDir(ex)
		if rel, ok := onDisk[ex.ID]; ok && rel == want {
			claimed[ex.ID] = true
		} else {
			// Moved module (same ID), or renamed (old ID on disk).
			var found []Move
			if ok {
				found = append(found, Move{ID: ex.ID, OldID: ex.ID, From: rel, To: want})
			}
			olds := oldIDs[ex.ID]
			sort.Strings(olds)
			for _, old := range olds {
				if rel, ok := onDisk[old]; ok {
					found = append(found, Move{ID: ex.ID, OldID: old, From: rel, To: want})
				}
			}
			switch {
			case len(found) == 0:
				plan.Added = append(plan.Added, ex.ID)
				continue
			case len(found) > 1:
				plan.Conflicts = append(plan.Conflicts, fmt.Sprintf("exercise %s has several candidate directories: %s", ex.ID, joinFrom(found)))
				continue
			}
			mv := found[0]
			claimed[mv.OldID] = true
			if _, err := os.Lstat(filepath.Join(w.Root, mv.To)); err == nil {
				plan.Conflicts = append(plan.Conflicts, fmt.Sprintf("cannot move %s to %s: the destination already exists", mv.From, mv.To))
				continue
			}
			plan.Moved = append(plan.Moved, mv)
		}
		stale, err := w.ownedFilesDiffer(cat, ex)
		if err != nil {
			return plan, err
		}
		if stale {
			plan.Refreshed = append(plan.Refreshed, ex.ID)
		}
	}
	for id, rel := range onDisk {
		if !claimed[id] {
			plan.Retired = append(plan.Retired, rel)
		}
	}
	sort.Strings(plan.Retired)
	return plan, nil
}

func joinFrom(ms []Move) string {
	var s []string
	for _, m := range ms {
		s = append(s, m.From)
	}
	return strings.Join(s, ", ")
}

// scanExerciseDirs maps the exercise IDs present on disk (exercises/<module>/<id>)
// to their workspace-relative directories.
func (w *Workspace) scanExerciseDirs() (map[string]string, []string, error) {
	out := map[string]string{}
	var dupes []string
	mods, err := os.ReadDir(filepath.Join(w.Root, exercisesDir))
	if errors.Is(err, fs.ErrNotExist) {
		return out, nil, nil
	}
	if err != nil {
		return nil, nil, err
	}
	for _, m := range mods {
		if !m.IsDir() {
			continue
		}
		entries, err := os.ReadDir(filepath.Join(w.Root, exercisesDir, m.Name()))
		if err != nil {
			return nil, nil, err
		}
		for _, e := range entries {
			if !e.IsDir() {
				continue
			}
			rel := filepath.Join(exercisesDir, m.Name(), e.Name())
			if _, ok := out[e.Name()]; ok {
				dupes = append(dupes, e.Name())
				continue
			}
			out[e.Name()] = rel
		}
	}
	sort.Strings(dupes)
	return out, dupes, nil
}

// ownedFiles are the files golings owns inside an exercise directory: the
// README and the fixture copies, keyed by slash-separated relative path.
func ownedFiles(cat *catalog.Catalog, ex *catalog.Exercise) (map[string][]byte, error) {
	assets, err := cat.Assets(ex)
	if err != nil {
		return nil, err
	}
	out := map[string][]byte{}
	if data, err := fs.ReadFile(assets, "README.md"); err == nil {
		out["README.md"] = data
	}
	for _, rel := range ex.FixtureFiles {
		data, err := fs.ReadFile(assets, path.Join("fixtures", rel))
		if err != nil {
			return nil, err
		}
		out[rel] = data
	}
	return out, nil
}

// ownedFilesDiffer reports whether the README or a fixture copy of an
// existing exercise differs from the catalog's.
func (w *Workspace) ownedFilesDiffer(cat *catalog.Catalog, ex *catalog.Exercise) (bool, error) {
	owned, err := ownedFiles(cat, ex)
	if err != nil {
		return false, err
	}
	for rel, data := range owned {
		cur, err := os.ReadFile(filepath.Join(w.ExerciseDir(ex), filepath.FromSlash(rel)))
		if err != nil || !bytes.Equal(cur, data) {
			// A move is planned before the files are at ExerciseDir; the
			// refresh then simply rewrites them after the move.
			return true, nil
		}
	}
	return false, nil
}

// ApplyUpgrade carries out plan: it moves renamed exercises, creates new
// ones, refreshes READMEs and fixture copies, and records the new curriculum
// version. Learner files are moved but never rewritten. Progress (which lives
// in the progress store) is migrated separately by the caller.
func (w *Workspace) ApplyUpgrade(cat *catalog.Catalog, plan UpgradePlan) error {
	if len(plan.Conflicts) > 0 {
		return fmt.Errorf("upgrade blocked: %s", strings.Join(plan.Conflicts, "; "))
	}
	for _, mv := range plan.Moved {
		to := filepath.Join(w.Root, mv.To)
		if err := os.MkdirAll(filepath.Dir(to), 0o755); err != nil {
			return err
		}
		if err := os.Rename(filepath.Join(w.Root, mv.From), to); err != nil {
			return fmt.Errorf("moving %s to %s: %w", mv.From, mv.To, err)
		}
		// Drop the old module directory if the move emptied it.
		os.Remove(filepath.Dir(filepath.Join(w.Root, mv.From)))
	}
	for _, id := range plan.Added {
		ex, _ := cat.Get(id)
		if err := w.materialize(cat, ex); err != nil {
			return fmt.Errorf("creating %s: %w", id, err)
		}
	}
	for _, id := range plan.Refreshed {
		ex, _ := cat.Get(id)
		owned, err := ownedFiles(cat, ex)
		if err != nil {
			return err
		}
		for rel, data := range owned {
			target := filepath.Join(w.ExerciseDir(ex), filepath.FromSlash(rel))
			if cur, err := os.ReadFile(target); err == nil && bytes.Equal(cur, data) {
				continue
			}
			if fi, err := os.Lstat(target); err == nil && !fi.Mode().IsRegular() {
				if err := os.Remove(target); err != nil {
					return err
				}
			}
			if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
				return err
			}
			if err := replaceFile(target, data); err != nil {
				return fmt.Errorf("refreshing %s/%s: %w", id, rel, err)
			}
		}
	}
	if err := replaceFile(filepath.Join(w.Root, "README.md"), []byte(workspaceReadme(cat))); err != nil {
		return err
	}
	w.Meta.CurriculumVersion = cat.Version()
	now := time.Now().UTC().Truncate(time.Second)
	w.Meta.UpgradedAt = &now
	data, err := json.MarshalIndent(w.Meta, "", "  ")
	if err != nil {
		return err
	}
	return replaceFile(filepath.Join(w.MetaPath(), metaFile), append(data, '\n'))
}

// CompareVersions compares dotted numeric versions ("0.10.2" > "0.9.0").
// Non-numeric parts compare as strings; a missing part counts as 0.
func CompareVersions(a, b string) int {
	as, bs := strings.Split(strings.TrimPrefix(a, "v"), "."), strings.Split(strings.TrimPrefix(b, "v"), ".")
	for i := range max(len(as), len(bs)) {
		x, y := "0", "0"
		if i < len(as) {
			x = as[i]
		}
		if i < len(bs) {
			y = bs[i]
		}
		xn, xerr := strconv.Atoi(x)
		yn, yerr := strconv.Atoi(y)
		switch {
		case xerr == nil && yerr == nil:
			if xn != yn {
				if xn < yn {
					return -1
				}
				return 1
			}
		case x != y:
			return strings.Compare(x, y)
		}
	}
	return 0
}
