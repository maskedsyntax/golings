// defer_03: Defer in a loop
//
// TotalSize keeps every file open until it returns. Fix it. See README.md.
//
// Reference: https://gobyexample.com/defer
package total

import "io"

// Opener opens a named file.
type Opener func(name string) (io.ReadCloser, error)

// TotalSize returns the combined size in bytes of the named files.
func TotalSize(open Opener, names []string) (int, error) {
	total := 0
	for _, name := range names {
		f, err := open(name)
		if err != nil {
			return 0, err
		}
		defer f.Close()
		b, err := io.ReadAll(f)
		if err != nil {
			return 0, err
		}
		total += len(b)
	}
	return total, nil
}
