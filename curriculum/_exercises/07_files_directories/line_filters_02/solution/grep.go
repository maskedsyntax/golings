// line_filters_02: A tiny grep (reference solution)
package grep

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Grep writes "n:line" for every line of r containing pattern and returns the match count.
func Grep(r io.Reader, w io.Writer, pattern string) (int, error) {
	sc := bufio.NewScanner(r) // ScanLines also drops a trailing \r
	matches := 0
	for lineNo := 1; sc.Scan(); lineNo++ {
		if line := sc.Text(); strings.Contains(line, pattern) {
			fmt.Fprintf(w, "%d:%s\n", lineNo, line)
			matches++
		}
	}
	return matches, sc.Err()
}
