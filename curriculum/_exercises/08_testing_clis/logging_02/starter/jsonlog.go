// logging_02: Structured logs
//
// Implement NewJSONLogger. See README.md.
//
// Reference: https://gobyexample.com/logging
package jsonlog

import (
	"io"
	"log/slog"
)

// NewJSONLogger returns a JSON logger without timestamps that tags every record with service=api.
func NewJSONLogger(w io.Writer, level slog.Level) *slog.Logger {
	// TODO
	return slog.New(slog.NewJSONHandler(w, nil))
}
