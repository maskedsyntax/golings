// exit_01: os.Exit skips defers (reference solution)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Exit(run()) // nothing deferred in main, so nothing is skipped
}

func run() int {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush() // runs when run returns, before os.Exit

	files := []string{"a.go", "b.go", "c.go"}
	fmt.Fprintf(w, "checking %d files...\n", len(files))
	problems := 0
	for _, f := range files {
		if !strings.HasPrefix(f, "a") {
			problems++
		}
	}
	fmt.Fprintf(w, "found %d problems\n", problems)
	return problems
}
