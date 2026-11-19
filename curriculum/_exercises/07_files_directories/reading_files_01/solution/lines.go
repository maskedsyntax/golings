// reading_files_01: Reading a whole file (reference solution)
package lines

import (
	"os"
	"strings"
)

// CountLines returns the number of lines in the file at path. A last line
// without a trailing newline still counts.
func CountLines(path string) (int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	text := string(data)
	if text == "" {
		return 0, nil
	}
	return strings.Count(strings.TrimSuffix(text, "\n"), "\n") + 1, nil
}
