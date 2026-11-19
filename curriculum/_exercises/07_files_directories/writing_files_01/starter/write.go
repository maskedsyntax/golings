// writing_files_01: Flush before you close
//
// WriteLines loses data. Fix it. See README.md.
//
// Reference: https://gobyexample.com/writing-files
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
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, l := range lines {
		fmt.Fprintln(w, l)
	}
	return nil
}
