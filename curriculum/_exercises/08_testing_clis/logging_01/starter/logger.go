// logging_01: Logger prefixes and flags
//
// Fix the logger settings. See README.md.
//
// Reference: https://gobyexample.com/logging
package logger

import (
	"io"
	"log"
)

// NewLogger logs lines like "app: started".
func NewLogger(w io.Writer) *log.Logger {
	return log.New(w, "app", log.LstdFlags)
}

// NewDebugLogger logs lines like "main.go:12: debug: cache miss".
func NewDebugLogger(w io.Writer) *log.Logger {
	return log.New(w, "debug: ", log.Lshortfile)
}
