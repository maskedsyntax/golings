// Mutant: reports the error but keeps the 64 KiB limit.
package longest

import (
	"bufio"
	"os"
)

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
	return longest, sc.Err()
}
