// Mutant: os.Remove can't delete a non-empty directory.
package scratch

import "os"

func Scratch(fn func(dir string) error) error {
	dir, err := os.MkdirTemp("", "scratch-*")
	if err != nil {
		return err
	}
	defer os.Remove(dir)
	return fn(dir)
}
