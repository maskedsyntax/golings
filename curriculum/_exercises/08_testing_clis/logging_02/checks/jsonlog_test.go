package jsonlog

import (
	"log/slog"
	"strings"
	"testing"
)

func TestNewJSONLogger(t *testing.T) {
	var b strings.Builder
	l := NewJSONLogger(&b, slog.LevelInfo)
	l.Info("started", "port", 8080)
	l.Debug("noisy detail")
	l.Warn("slow request", "ms", 950, "path", "/x")
	want := `{"level":"INFO","msg":"started","service":"api","port":8080}
{"level":"WARN","msg":"slow request","service":"api","ms":950,"path":"/x"}
`
	if b.String() != want {
		t.Fatalf("logged:\n%s\nwant:\n%s", b.String(), want)
	}
}

func TestDebugLevel(t *testing.T) {
	var b strings.Builder
	NewJSONLogger(&b, slog.LevelDebug).Debug("detail", "k", "v")
	if want := `{"level":"DEBUG","msg":"detail","service":"api","k":"v"}` + "\n"; b.String() != want {
		t.Fatalf("logged %s, want %s", b.String(), want)
	}
}

func TestGroupsKeepTheirTime(t *testing.T) {
	var b strings.Builder
	NewJSONLogger(&b, slog.LevelInfo).Info("job", slog.Group("job", "time", "5m"))
	if !strings.Contains(b.String(), `"job":{"time":"5m"}`) {
		t.Fatalf("logged %s; only the top-level time should be dropped", b.String())
	}
}
