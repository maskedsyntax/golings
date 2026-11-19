// temporary_files_and_directories_01: Cleaning up a temp file
//
// WithTempFile leaks files. Fix it. See README.md.
//
// Reference: https://gobyexample.com/temporary-files-and-directories
package tempfile

import "os"

// WithTempFile writes data to a new temporary file, calls fn with its path,
// and removes the file afterwards.
func WithTempFile(data []byte, fn func(path string) error) error {
	f, err := os.CreateTemp("", "golings-*.txt")
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		return err
	}
	return fn(f.Name())
}
