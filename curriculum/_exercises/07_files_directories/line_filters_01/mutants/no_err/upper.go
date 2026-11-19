// Mutant: swallows read errors.
package upper

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Upper(r io.Reader, w io.Writer) error {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		fmt.Fprintln(w, strings.ToUpper(sc.Text()))
	}
	return nil
}
