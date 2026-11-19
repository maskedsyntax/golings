// reading_files_02: Reading part of a file (reference solution)
package peek

import (
	"errors"
	"io"
	"os"
)

// Peek returns up to n bytes of the file at path, starting at offset.
func Peek(path string, offset int64, n int) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	buf := make([]byte, n)
	got, err := f.ReadAt(buf, offset)
	if err != nil && !errors.Is(err, io.EOF) { // EOF just means a short read
		return "", err
	}
	return string(buf[:got]), nil
}
