// Mutant: never cleans up the temporary file on failure.
package atomic

import (
	"io"
	"os"
	"path/filepath"
)

func WriteFileAtomic(path string, r io.Reader) error {
	f, err := os.CreateTemp(filepath.Dir(path), ".tmp-*")
	if err != nil {
		return err
	}
	if _, err := io.Copy(f, r); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(f.Name(), path)
}
