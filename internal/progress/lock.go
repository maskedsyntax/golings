package progress

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

// staleLockAge is how old a lock may be before it is considered abandoned
// even if its owner appears alive (for example after PID reuse).
const staleLockAge = 10 * time.Minute

// lockWait is how long Update waits for another process to release the lock.
var lockWait = 5 * time.Second

// ErrLocked means another golings process holds the state lock.
var ErrLocked = errors.New("another golings process is updating progress; try again in a moment")

// acquireFileLock creates path exclusively, retrying until wait elapses. Locks
// whose owner process is gone, or that are older than staleLockAge, are
// broken.
// inProcess serializes goroutines of this process per lock path, so the lock
// file only arbitrates between processes and goroutines never starve each
// other by polling it.
var inProcess sync.Map // lock path -> *sync.Mutex

func acquireLock(path string, wait time.Duration) (func(), error) {
	mu, _ := inProcess.LoadOrStore(filepath.Clean(path), &sync.Mutex{})
	m := mu.(*sync.Mutex)
	m.Lock()
	fileRelease, err := acquireFileLock(path, wait)
	if err != nil {
		m.Unlock()
		return nil, err
	}
	return func() { fileRelease(); m.Unlock() }, nil
}

func acquireFileLock(path string, wait time.Duration) (release func(), err error) {
	deadline := time.Now().Add(wait)
	for {
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
		if err == nil {
			fmt.Fprintf(f, "%d\n%s\n", os.Getpid(), time.Now().UTC().Format(time.RFC3339))
			f.Close()
			return func() { os.Remove(path) }, nil
		}
		if !errors.Is(err, fs.ErrExist) {
			return nil, fmt.Errorf("creating lock file: %w", err)
		}
		if lockIsStale(path) {
			os.Remove(path)
			continue
		}
		if time.Now().After(deadline) {
			return nil, ErrLocked
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func lockIsStale(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false // vanished: the next attempt will succeed
	}
	if time.Since(fi.ModTime()) > staleLockAge {
		return true
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	first, _, _ := strings.Cut(string(data), "\n")
	pid, err := strconv.Atoi(strings.TrimSpace(first))
	if err != nil {
		// Unreadable owner: only stale once it has had time to be written.
		return time.Since(fi.ModTime()) > 5*time.Second
	}
	return pid != os.Getpid() && !pidAlive(pid)
}
