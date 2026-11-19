// reading_files_03: Very long lines (reference solution)
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
	sc.Buffer(make([]byte, 0, 64*1024), 1<<20) // allow lines up to 1 MiB
	for sc.Scan() {
		longest = max(longest, len(sc.Bytes()))
	}
	if err := sc.Err(); err != nil { // bufio.ErrTooLong, or a read error
		return 0, err
	}
	return longest, nil
}
