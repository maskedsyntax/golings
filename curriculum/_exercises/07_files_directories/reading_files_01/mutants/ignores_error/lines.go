// Mutant: a missing file counts as empty.
package lines

import (
	"os"
	"strings"
)

func CountLines(path string) (int, error) {
	data, _ := os.ReadFile(path)
	text := string(data)
	if text == "" {
		return 0, nil
	}
	return strings.Count(strings.TrimSuffix(text, "\n"), "\n") + 1, nil
}
