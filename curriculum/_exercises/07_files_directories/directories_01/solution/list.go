// directories_01: Listing a directory (reference solution)
package list

import (
	"os"
	"strings"
)

// ListGoFiles returns the sorted names of the non-test .go files directly in dir.
func ListGoFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir) // sorted by name
	if err != nil {
		return nil, err
	}
	var names []string
	for _, e := range entries {
		name := e.Name()
		if !e.IsDir() && strings.HasSuffix(name, ".go") && !strings.HasSuffix(name, "_test.go") {
			names = append(names, name)
		}
	}
	return names, nil
}
