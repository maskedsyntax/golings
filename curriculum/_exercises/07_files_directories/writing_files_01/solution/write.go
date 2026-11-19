// writing_files_01: Flush before you close (reference solution)
package write

import (
	"bufio"
	"fmt"
	"os"
)

// WriteLines creates (or truncates) the file at path and writes each line followed by "\n".
func WriteLines(path string, lines []string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	for _, l := range lines {
		fmt.Fprintln(w, l) // errors are sticky; Flush reports the first one
	}
	if err := w.Flush(); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}
