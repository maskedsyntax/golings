// exit_01: os.Exit skips defers
//
// The program prints nothing. Fix it. See README.md.
//
// Reference: https://gobyexample.com/exit
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	files := []string{"a.go", "b.go", "c.go"}
	fmt.Fprintf(w, "checking %d files...\n", len(files))
	problems := 0
	for _, f := range files {
		if !strings.HasPrefix(f, "a") {
			problems++
		}
	}
	fmt.Fprintf(w, "found %d problems\n", problems)
	os.Exit(problems)
}
