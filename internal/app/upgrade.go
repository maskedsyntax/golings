package app

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/maskedsyntax/golings/internal/progress"
	"github.com/maskedsyntax/golings/internal/workspace"
)

// UpgradeResult reports what Upgrade did (or, for a dry run, would do).
type UpgradeResult struct {
	workspace.UpgradePlan
	// ProgressRenamed lists progress entries carried over to a new ID.
	ProgressRenamed map[string]string `json:"progress_renamed"`
	// StateBackup is where state.json was copied before it was changed.
	StateBackup string `json:"state_backup,omitempty"`
	DryRun      bool   `json:"dry_run"`
}

// Upgrade brings the workspace to the catalog's curriculum version. With
// dryRun it only reports the plan.
func (a *App) Upgrade(dryRun bool) (*UpgradeResult, error) {
	plan, err := a.Workspace.PlanUpgrade(a.Catalog)
	if err != nil {
		return nil, err
	}
	res := &UpgradeResult{UpgradePlan: plan, DryRun: dryRun, ProgressRenamed: map[string]string{}}

	st, err := a.Store.Load()
	if err != nil {
		return nil, err
	}
	renames := a.Catalog.Renames()
	for old, cur := range renames {
		if _, ok := st.Exercises[old]; ok && cur != "" {
			if _, taken := st.Exercises[cur]; !taken {
				res.ProgressRenamed[old] = cur
			}
		}
	}
	if dryRun || len(plan.Conflicts) > 0 {
		return res, nil
	}

	if err := a.Workspace.ApplyUpgrade(a.Catalog, plan); err != nil {
		return res, err
	}
	if len(res.ProgressRenamed) == 0 && !renamesCurrent(st, renames) {
		return res, nil
	}
	backup, err := backupState(a.Store.Path(), a.Workspace.MetaPath())
	if err != nil {
		return res, fmt.Errorf("backing up progress: %w", err)
	}
	res.StateBackup = backup
	_, err = a.Store.Update(func(s *progress.State) error {
		for old, cur := range res.ProgressRenamed {
			if _, taken := s.Exercises[cur]; !taken {
				s.Exercises[cur] = s.Exercises[old]
				delete(s.Exercises, old)
			}
		}
		if cur, ok := renames[s.CurrentExerciseID]; ok {
			s.CurrentExerciseID = cur // "" (removed) falls back to the next unfinished
		}
		return nil
	})
	return res, err
}

func renamesCurrent(st *progress.State, renames map[string]string) bool {
	_, ok := renames[st.CurrentExerciseID]
	return ok
}

// backupState copies state.json into .golings/backups/ and returns the copy's
// path ("" if there was no state yet).
func backupState(statePath, metaDir string) (string, error) {
	data, err := os.ReadFile(statePath)
	if errors.Is(err, fs.ErrNotExist) {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	dir := filepath.Join(metaDir, "backups")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	dest := filepath.Join(dir, "state-"+time.Now().UTC().Format("20060102T150405.000Z")+".json")
	return dest, os.WriteFile(dest, data, 0o644)
}
