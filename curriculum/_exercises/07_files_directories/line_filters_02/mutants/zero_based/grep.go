// Mutant: line numbers start at 0.
package grep

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Grep(r io.Reader, w io.Writer, pattern string) (int, error) {
	sc := bufio.NewScanner(r)
	matches := 0
	for lineNo := 0; sc.Scan(); lineNo++ {
		if line := sc.Text(); strings.Contains(line, pattern) {
			fmt.Fprintf(w, "%d:%s\n", lineNo, line)
			matches++
		}
	}
	return matches, sc.Err()
}
