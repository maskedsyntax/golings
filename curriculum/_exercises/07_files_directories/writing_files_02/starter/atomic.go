// writing_files_02: Atomic writes
//
// Implement WriteFileAtomic. See README.md.
//
// Reference: https://gobyexample.com/writing-files
package atomic

import (
	"errors"
	"io"
)

// WriteFileAtomic replaces the file at path with the contents of r, all or nothing.
func WriteFileAtomic(path string, r io.Reader) error {
	// TODO
	return errors.New("not implemented")
}
