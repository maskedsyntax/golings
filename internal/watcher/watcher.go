// Package watcher detects content changes to a fixed set of files by
// polling. Polling is portable and plenty fast for a handful of small files;
// a cheap size/mtime check gates the content hash.
package watcher

import (
	"context"
	"crypto/sha256"
	"io"
	"os"
	"time"
)

// Defaults tuned for "save in editor -> feedback" latency well under 500 ms.
const (
	DefaultInterval = 150 * time.Millisecond
	DefaultDebounce = 250 * time.Millisecond
)

// Options configures Watch.
type Options struct {
	Interval time.Duration // poll period (default DefaultInterval)
	Debounce time.Duration // quiet period before reporting (default DefaultDebounce)
}

type fileStat struct {
	exists  bool
	size    int64
	modTime time.Time
	hash    [32]byte
}

// Watch reports on the returned channel whenever the content of any of
// paths changes (including creation and deletion), after the files have been
// quiet for the debounce period. Saves that do not change content (an editor
// touching the file) are ignored. The channel is closed when ctx ends.
func Watch(ctx context.Context, paths []string, opts Options) <-chan struct{} {
	if opts.Interval <= 0 {
		opts.Interval = DefaultInterval
	}
	if opts.Debounce <= 0 {
		opts.Debounce = DefaultDebounce
	}
	out := make(chan struct{}, 1)
	go func() {
		defer close(out)
		last := snapshot(paths, nil)
		reported := last
		var quietSince time.Time
		pending := false

		ticker := time.NewTicker(opts.Interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case now := <-ticker.C:
				cur := snapshot(paths, last)
				changed := !equal(cur, last)
				last = cur // keep fresh stats so unchanged files are not re-hashed
				if changed {
					quietSince = now
					pending = !equal(cur, reported)
					continue
				}
				if pending && now.Sub(quietSince) >= opts.Debounce {
					pending = false
					reported = cur
					select {
					case out <- struct{}{}:
					default: // a notification is already queued
					}
				}
			}
		}
	}()
	return out
}

// snapshot stats every path and hashes only files whose size or mtime
// changed since prev.
func snapshot(paths []string, prev []fileStat) []fileStat {
	out := make([]fileStat, len(paths))
	for i, p := range paths {
		fi, err := os.Stat(p)
		if err != nil {
			continue
		}
		s := fileStat{exists: true, size: fi.Size(), modTime: fi.ModTime()}
		if prev != nil && prev[i].exists && prev[i].size == s.size && prev[i].modTime.Equal(s.modTime) {
			s.hash = prev[i].hash
		} else {
			s.hash = hashFile(p)
		}
		out[i] = s
	}
	return out
}

func hashFile(path string) [32]byte {
	var sum [32]byte
	f, err := os.Open(path)
	if err != nil {
		return sum
	}
	defer f.Close()
	h := sha256.New()
	io.Copy(h, io.LimitReader(f, 4<<20))
	copy(sum[:], h.Sum(nil))
	return sum
}

// equal compares content identity (existence and hash) only.
func equal(a, b []fileStat) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].exists != b[i].exists || a[i].hash != b[i].hash {
			return false
		}
	}
	return true
}
