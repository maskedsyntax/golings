// Mutant: writes straight into the target.
package atomic

import (
	"io"
	"os"
)

func WriteFileAtomic(path string, r io.Reader) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if _, err := io.Copy(f, r); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}
