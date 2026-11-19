// temporary_files_and_directories_02: A scratch directory
//
// Implement Scratch. See README.md.
//
// Reference: https://gobyexample.com/temporary-files-and-directories
package scratch

import "errors"

// Scratch runs fn in a fresh temporary directory and deletes it afterwards.
func Scratch(fn func(dir string) error) error {
	// TODO
	return errors.New("not implemented")
}
