// temporary_files_and_directories_02: A scratch directory (reference solution)
package scratch

import "os"

// Scratch runs fn in a fresh temporary directory and deletes it afterwards.
func Scratch(fn func(dir string) error) error {
	dir, err := os.MkdirTemp("", "scratch-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir) // also runs if fn panics
	return fn(dir)
}
