// Package progress persists learner progress in .golings/state.json with
// serialized, atomic writes and recovery from corruption.
package progress

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

// SchemaVersion is the state.json format version.
const SchemaVersion = 1

// Status is the workflow state of an exercise. "failed" is an evaluation
// result, not a workflow state, so it does not appear here.
type Status string

const (
	NotStarted Status = "not_started"
	InProgress Status = "in_progress"
	Completed  Status = "completed"
	Skipped    Status = "skipped"
)

// ExerciseState is the persisted progress of one exercise.
type ExerciseState struct {
	Status            Status     `json:"status"`
	CompletedAt       *time.Time `json:"completed_at,omitempty"`
	Attempts          int        `json:"attempts"`
	HintsRevealed     int        `json:"hints_revealed"`
	SourceFingerprint string     `json:"source_fingerprint,omitempty"`
	SolutionViewed    bool       `json:"solution_viewed,omitempty"`
	LastResult        string     `json:"last_result,omitempty"`
	LastRunAt         *time.Time `json:"last_run_at,omitempty"`
}

// State is the whole state.json document.
type State struct {
	SchemaVersion     int                       `json:"schema_version"`
	CurriculumVersion string                    `json:"curriculum_version"`
	CurrentExerciseID string                    `json:"current_exercise_id,omitempty"`
	Exercises         map[string]*ExerciseState `json:"exercises"`
}

// New returns an empty state.
func New(curriculumVersion string) *State {
	return &State{SchemaVersion: SchemaVersion, CurriculumVersion: curriculumVersion, Exercises: map[string]*ExerciseState{}}
}

// Get returns the state of id (a not_started zero value if unknown).
func (s *State) Get(id string) ExerciseState {
	if e, ok := s.Exercises[id]; ok && e != nil {
		return *e
	}
	return ExerciseState{Status: NotStarted}
}

// Entry returns a mutable entry for id, creating it if needed.
func (s *State) Entry(id string) *ExerciseState {
	if s.Exercises == nil {
		s.Exercises = map[string]*ExerciseState{}
	}
	e, ok := s.Exercises[id]
	if !ok || e == nil {
		e = &ExerciseState{Status: NotStarted}
		s.Exercises[id] = e
	}
	return e
}

// RecordResult applies one evaluation outcome. Completion is recorded only
// for a pass, and a later failure never erases completion history: the
// fingerprint comparison reports "modified since verified" instead.
func (s *State) RecordResult(id string, passed bool, result string, fingerprint string, now time.Time) {
	e := s.Entry(id)
	e.Attempts++
	e.LastResult = result
	t := now.UTC().Truncate(time.Second)
	e.LastRunAt = &t
	if passed {
		e.Status = Completed
		e.CompletedAt = &t
		e.SourceFingerprint = fingerprint
		return
	}
	if e.Status == NotStarted {
		e.Status = InProgress
	}
}

const stateFile = "state.json"

// Store reads and writes state.json inside a .golings directory.
type Store struct {
	dir               string
	curriculumVersion string
}

// NewStore returns a store for the given .golings directory.
func NewStore(dir, curriculumVersion string) *Store {
	return &Store{dir: dir, curriculumVersion: curriculumVersion}
}

// Path returns the state file path.
func (s *Store) Path() string { return filepath.Join(s.dir, stateFile) }

// CorruptError means state.json could not be parsed. A copy was saved to
// BackupPath before reporting.
type CorruptError struct {
	Path       string
	BackupPath string
	Err        error
}

func (e *CorruptError) Error() string {
	return fmt.Sprintf("progress file %s is corrupt (%v); a copy was saved to %s. Run `golings doctor --recover-state` to start fresh, or fix the file by hand",
		e.Path, e.Err, e.BackupPath)
}

func (e *CorruptError) Unwrap() error { return e.Err }

// Load reads the state. A missing file yields a fresh state.
func (s *Store) Load() (*State, error) {
	data, err := os.ReadFile(s.Path())
	if errors.Is(err, fs.ErrNotExist) {
		return New(s.curriculumVersion), nil
	}
	if err != nil {
		return nil, err
	}
	st := &State{}
	if err := json.Unmarshal(data, st); err != nil || st.SchemaVersion != SchemaVersion {
		if err == nil {
			err = fmt.Errorf("unsupported schema_version %d", st.SchemaVersion)
		}
		backup := s.Path() + ".corrupt-" + time.Now().UTC().Format("20060102T150405Z")
		if werr := os.WriteFile(backup, data, 0o644); werr != nil {
			backup = "(backup failed: " + werr.Error() + ")"
		}
		return nil, &CorruptError{Path: s.Path(), BackupPath: backup, Err: err}
	}
	if st.Exercises == nil {
		st.Exercises = map[string]*ExerciseState{}
	}
	return st, nil
}

// Update loads the state under the lock, applies fn, and atomically writes
// the result. If fn returns an error nothing is written.
func (s *Store) Update(fn func(*State) error) (*State, error) {
	unlock, err := acquireLock(filepath.Join(s.dir, "lock"), lockWait)
	if err != nil {
		return nil, err
	}
	defer unlock()

	st, err := s.Load()
	if err != nil {
		return nil, err
	}
	if err := fn(st); err != nil {
		return nil, err
	}
	st.SchemaVersion = SchemaVersion
	st.CurriculumVersion = s.curriculumVersion
	data, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		return nil, err
	}
	if err := writeAtomic(s.Path(), append(data, '\n')); err != nil {
		return nil, fmt.Errorf("saving progress: %w", err)
	}
	return st, nil
}

// Recover moves a corrupt state file aside and starts fresh. It returns the
// path of the moved file ("" if there was nothing to recover).
func (s *Store) Recover() (string, error) {
	unlock, err := acquireLock(filepath.Join(s.dir, "lock"), lockWait)
	if err != nil {
		return "", err
	}
	defer unlock()
	if _, err := os.Stat(s.Path()); errors.Is(err, fs.ErrNotExist) {
		return "", nil
	}
	moved := s.Path() + ".recovered-" + time.Now().UTC().Format("20060102T150405Z")
	if err := os.Rename(s.Path(), moved); err != nil {
		return "", err
	}
	return moved, nil
}

// writeAtomic writes data to a temp file in the same directory, syncs it, and
// renames it over name, so readers never observe a partial file.
func writeAtomic(name string, data []byte) error {
	f, err := os.CreateTemp(filepath.Dir(name), "."+filepath.Base(name)+".tmp-*")
	if err != nil {
		return err
	}
	tmp := f.Name()
	ok := false
	defer func() {
		if !ok {
			os.Remove(tmp)
		}
	}()
	if _, err := f.Write(data); err != nil {
		f.Close()
		return err
	}
	if err := f.Sync(); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	if err := os.Rename(tmp, name); err != nil {
		return err
	}
	ok = true
	return nil
}
