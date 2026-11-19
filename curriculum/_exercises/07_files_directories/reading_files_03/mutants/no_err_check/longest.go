// Mutant: raises the limit but never checks Err.
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
	sc.Buffer(make([]byte, 0, 64*1024), 1<<20)
	for sc.Scan() {
		longest = max(longest, len(sc.Bytes()))
	}
	return longest, nil
}
