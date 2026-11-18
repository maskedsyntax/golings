package service

import (
	"slices"
	"testing"
)

func TestServiceLogPrefixes(t *testing.T) {
	l := &Logger{}
	api := &Service{Logger: l, Name: "api"}
	db := &Service{Logger: l, Name: "db"}
	api.Log("started")
	db.Log("connected")
	api.Log("ready")
	want := []string{"[api] started", "[db] connected", "[api] ready"}
	if got := l.Lines(); !slices.Equal(got, want) {
		t.Errorf("logger lines = %q, want %q", got, want)
	}
	if got := api.Lines(); !slices.Equal(got, want) {
		t.Errorf("api.Lines() (promoted) = %q, want %q", got, want)
	}
}

func TestLoggerItselfUnchanged(t *testing.T) {
	l := &Logger{}
	l.Log("plain")
	if got := l.Lines(); !slices.Equal(got, []string{"plain"}) {
		t.Errorf("Logger.Log added a prefix: %q", got)
	}
}
