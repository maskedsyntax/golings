// Mutant: always reads from the start.
package peek

import (
	"errors"
	"io"
	"os"
)

func Peek(path string, offset int64, n int) (string, error) {
	if offset < 0 {
		return "", errors.New("negative offset")
	}
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	buf := make([]byte, n)
	got, err := io.ReadFull(f, buf)
	if err != nil && !errors.Is(err, io.ErrUnexpectedEOF) && !errors.Is(err, io.EOF) {
		return "", err
	}
	return string(buf[:got]), nil
}
