// temporary_files_and_directories_01: Cleaning up a temp file (reference solution)
package tempfile

import "os"

// WithTempFile writes data to a new temporary file, calls fn with its path,
// and removes the file afterwards.
func WithTempFile(data []byte, fn func(path string) error) error {
	f, err := os.CreateTemp("", "golings-*.txt")
	if err != nil {
		return err
	}
	name := f.Name()
	defer os.Remove(name) // runs however we return

	if _, err := f.Write(data); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return fn(name)
}
