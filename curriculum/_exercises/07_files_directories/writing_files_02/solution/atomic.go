// writing_files_02: Atomic writes (reference solution)
package atomic

import (
	"io"
	"os"
	"path/filepath"
)

// WriteFileAtomic replaces the file at path with the contents of r, all or nothing.
func WriteFileAtomic(path string, r io.Reader) (err error) {
	// Same directory, so the rename never crosses file systems.
	f, err := os.CreateTemp(filepath.Dir(path), "."+filepath.Base(path)+".tmp-*")
	if err != nil {
		return err
	}
	tmp := f.Name()
	defer func() {
		if err != nil {
			f.Close()
			os.Remove(tmp)
		}
	}()

	if _, err = io.Copy(f, r); err != nil {
		return err
	}
	if err = f.Sync(); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
