// logging_01: Logger prefixes and flags (reference solution)
package logger

import (
	"io"
	"log"
)

// NewLogger logs lines like "app: started".
func NewLogger(w io.Writer) *log.Logger {
	return log.New(w, "app: ", 0) // 0: no date or time
}

// NewDebugLogger logs lines like "main.go:12: debug: cache miss".
func NewDebugLogger(w io.Writer) *log.Logger {
	return log.New(w, "debug: ", log.Lshortfile|log.Lmsgprefix)
}
