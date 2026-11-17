package progress

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

func TestLoadMissingIsFresh(t *testing.T) {
	st, err := NewStore(t.TempDir(), "1.0.0").Load()
	if err != nil {
		t.Fatal(err)
	}
	if st.SchemaVersion != SchemaVersion || len(st.Exercises) != 0 || st.Get("x").Status != NotStarted {
		t.Errorf("fresh state = %+v", st)
	}
}

func TestUpdatePersistsAtomically(t *testing.T) {
	dir := t.TempDir()
	s := NewStore(dir, "1.0.0")
	now := time.Date(2026, 11, 17, 16, 0, 0, 0, time.UTC)
	if _, err := s.Update(func(st *State) error {
		st.CurrentExerciseID = "maps_01"
		st.RecordResult("maps_01", false, "fail", "sha256:a", now)
		st.RecordResult("maps_01", true, "pass", "sha256:b", now)
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	st, err := s.Load()
	if err != nil {
		t.Fatal(err)
	}
	e := st.Get("maps_01")
	if e.Status != Completed || e.Attempts != 2 || e.SourceFingerprint != "sha256:b" || e.CompletedAt == nil || !e.CompletedAt.Equal(now) {
		t.Errorf("state = %+v", e)
	}
	if st.CurrentExerciseID != "maps_01" {
		t.Errorf("current = %q", st.CurrentExerciseID)
	}
	// No temp files or lock are left behind.
	entries, _ := os.ReadDir(dir)
	if len(entries) != 1 || entries[0].Name() != stateFile {
		var names []string
		for _, e := range entries {
			names = append(names, e.Name())
		}
		t.Errorf("directory contains %v, want only %s", names, stateFile)
	}
}

func TestFailureKeepsCompletion(t *testing.T) {
	st := New("1")
	st.RecordResult("a", true, "pass", "fp1", time.Now())
	st.RecordResult("a", false, "fail", "fp2", time.Now())
	if e := st.Get("a"); e.Status != Completed || e.SourceFingerprint != "fp1" || e.LastResult != "fail" {
		t.Errorf("after pass then fail: %+v", e)
	}
	st.RecordResult("b", false, "fail", "fp", time.Now())
	if e := st.Get("b"); e.Status != InProgress {
		t.Errorf("after first failure: %+v", e)
	}
}

func TestUpdateErrorWritesNothing(t *testing.T) {
	s := NewStore(t.TempDir(), "1")
	boom := errors.New("boom")
	if _, err := s.Update(func(*State) error { return boom }); !errors.Is(err, boom) {
		t.Fatalf("err = %v", err)
	}
	if _, err := os.Stat(s.Path()); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("state file written despite error: %v", err)
	}
}

func TestCorruptStateIsBackedUpAndRecoverable(t *testing.T) {
	dir := t.TempDir()
	s := NewStore(dir, "1")
	if err := os.WriteFile(s.Path(), []byte("{not json"), 0o644); err != nil {
		t.Fatal(err)
	}
	_, err := s.Load()
	var ce *CorruptError
	if !errors.As(err, &ce) {
		t.Fatalf("Load error = %v, want *CorruptError", err)
	}
	if data, err := os.ReadFile(ce.BackupPath); err != nil || string(data) != "{not json" {
		t.Errorf("backup %s = %q, %v", ce.BackupPath, data, err)
	}
	// Update refuses to overwrite the corrupt file.
	if _, err := s.Update(func(*State) error { return nil }); !errors.As(err, &ce) {
		t.Errorf("Update on corrupt state: err = %v", err)
	}
	moved, err := s.Recover()
	if err != nil || moved == "" {
		t.Fatalf("Recover = %q, %v", moved, err)
	}
	if _, err := s.Load(); err != nil {
		t.Errorf("Load after Recover: %v", err)
	}
}

func TestConcurrentUpdatesAreSerialized(t *testing.T) {
	s := NewStore(t.TempDir(), "1")
	const n = 20
	var wg sync.WaitGroup
	errs := make(chan error, n)
	for i := range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := s.Update(func(st *State) error {
				st.Entry(fmt.Sprintf("ex_%02d", i)).Attempts++
				st.Entry("shared").Attempts++
				return nil
			})
			errs <- err
		}()
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}
	st, err := s.Load()
	if err != nil {
		t.Fatal(err)
	}
	if got := st.Get("shared").Attempts; got != n {
		t.Errorf("shared attempts = %d, want %d (lost updates)", got, n)
	}
}

func TestLockHeldByLiveProcessTimesOut(t *testing.T) {
	old := lockWait
	lockWait = 200 * time.Millisecond
	t.Cleanup(func() { lockWait = old })

	dir := t.TempDir()
	// Our parent process is alive and is not us, so the lock is not stale.
	lock := fmt.Sprintf("%d\n%s\n", os.Getppid(), time.Now().UTC().Format(time.RFC3339))
	if err := os.WriteFile(filepath.Join(dir, "lock"), []byte(lock), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := NewStore(dir, "1").Update(func(*State) error { return nil }); !errors.Is(err, ErrLocked) {
		t.Errorf("err = %v, want ErrLocked", err)
	}
}

func TestStaleLockIsBroken(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "lock")
	// A PID that is essentially guaranteed not to exist.
	if err := os.WriteFile(path, []byte("2147483646\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := NewStore(dir, "1").Update(func(*State) error { return nil }); err != nil {
		t.Errorf("Update with a dead owner's lock: %v", err)
	}

	// An old lock is stale even if its owner looks alive.
	if err := os.WriteFile(path, []byte(fmt.Sprintf("%d\n", os.Getppid())), 0o644); err != nil {
		t.Fatal(err)
	}
	old := time.Now().Add(-2 * staleLockAge)
	if err := os.Chtimes(path, old, old); err != nil {
		t.Fatal(err)
	}
	if _, err := NewStore(dir, "1").Update(func(*State) error { return nil }); err != nil {
		t.Errorf("Update with an old lock: %v", err)
	}
}
