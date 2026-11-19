// Mutant: splits on "\n" by hand and keeps the "\r".
package grep

import (
	"fmt"
	"io"
	"strings"
)

func Grep(r io.Reader, w io.Writer, pattern string) (int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	matches := 0
	for i, line := range strings.Split(strings.TrimSuffix(string(data), "\n"), "\n") {
		if strings.Contains(line, pattern) {
			fmt.Fprintf(w, "%d:%s\n", i+1, line)
			matches++
		}
	}
	return matches, nil
}
