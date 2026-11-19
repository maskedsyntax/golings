// reading_files_03: Very long lines
//
// LongestLine gives up quietly on long lines. Fix it. See README.md.
//
// Reference: https://gobyexample.com/reading-files
package longest

import (
	"bufio"
	"os"
)

// LongestLine returns the length in bytes of the longest line in the file.
func LongestLine(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	longest := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		longest = max(longest, len(sc.Bytes()))
	}
	return longest, nil
}
