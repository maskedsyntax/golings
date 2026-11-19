// Mutant: the debug prefix comes before file:line.
package logger

import (
	"io"
	"log"
)

func NewLogger(w io.Writer) *log.Logger {
	return log.New(w, "app: ", 0)
}

func NewDebugLogger(w io.Writer) *log.Logger {
	return log.New(w, "debug: ", log.Lshortfile)
}
