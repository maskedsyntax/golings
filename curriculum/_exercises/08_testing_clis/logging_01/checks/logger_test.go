package logger

import (
	"regexp"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	var b strings.Builder
	l := NewLogger(&b)
	l.Print("started")
	l.Printf("listening on %d", 8080)
	if want := "app: started\napp: listening on 8080\n"; b.String() != want {
		t.Fatalf("logged %q, want %q", b.String(), want)
	}
}

func TestNewDebugLogger(t *testing.T) {
	var b strings.Builder
	NewDebugLogger(&b).Print("cache miss")
	re := regexp.MustCompile(`^logger_test\.go:\d+: debug: cache miss\n$`)
	if !re.MatchString(b.String()) {
		t.Fatalf("logged %q, want something like %q", b.String(), "logger_test.go:31: debug: cache miss\n")
	}
}
