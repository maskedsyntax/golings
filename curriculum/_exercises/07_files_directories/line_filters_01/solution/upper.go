// line_filters_01: An uppercase filter (reference solution)
package upper

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Upper copies r to w line by line, uppercasing each line.
func Upper(r io.Reader, w io.Writer) error {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		fmt.Fprintln(w, strings.ToUpper(sc.Text()))
	}
	return sc.Err()
}
