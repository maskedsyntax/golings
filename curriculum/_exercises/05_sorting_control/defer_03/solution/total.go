// defer_03: Defer in a loop (reference solution)
package total

import "io"

// Opener opens a named file.
type Opener func(name string) (io.ReadCloser, error)

// TotalSize returns the combined size in bytes of the named files.
func TotalSize(open Opener, names []string) (int, error) {
	total := 0
	for _, name := range names {
		n, err := size(open, name)
		if err != nil {
			return 0, err
		}
		total += n
	}
	return total, nil
}

// size reads one file. Its defer runs when size returns, once per file.
func size(open Opener, name string) (int, error) {
	f, err := open(name)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	return len(b), err
}
