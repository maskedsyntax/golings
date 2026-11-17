// Command golings teaches Go through small exercises that you fix until
// their checks pass.
package main

import (
	"os"

	"github.com/maskedsyntax/golings/internal/cli"
)

func main() {
	os.Exit(cli.Main(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
