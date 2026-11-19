// Mutant: removes the file only when fn succeeds.
package tempfile

import "os"

func WithTempFile(data []byte, fn func(path string) error) error {
	f, err := os.CreateTemp("", "golings-*.txt")
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		f.Close()
		return err
	}
	f.Close()
	if err := fn(f.Name()); err != nil {
		return err
	}
	return os.Remove(f.Name())
}
