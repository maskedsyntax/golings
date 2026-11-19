// logging_02: Structured logs (reference solution)
package jsonlog

import (
	"io"
	"log/slog"
)

// NewJSONLogger returns a JSON logger without timestamps that tags every record with service=api.
func NewJSONLogger(w io.Writer, level slog.Level) *slog.Logger {
	h := slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if len(groups) == 0 && a.Key == slog.TimeKey {
				return slog.Attr{} // an empty Attr is dropped
			}
			return a
		},
	})
	return slog.New(h).With("service", "api")
}
