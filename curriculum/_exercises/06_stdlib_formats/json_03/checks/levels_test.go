package levels

import (
	"maps"
	"strings"
	"testing"
)

func TestCountByLevel(t *testing.T) {
	in := `{"level":"info","msg":"started"}
{"level":"warn","msg":"disk 90% full"}

{"msg":"no level"}
{"level":"info","msg":"done"}
`
	got, err := CountByLevel(strings.NewReader(in))
	if err != nil {
		t.Fatalf("CountByLevel returned an error: %v", err)
	}
	if want := map[string]int{"info": 2, "warn": 1, "unknown": 1}; !maps.Equal(got, want) {
		t.Fatalf("CountByLevel = %v, want %v", got, want)
	}
}

func TestCountByLevelEmpty(t *testing.T) {
	got, err := CountByLevel(strings.NewReader(""))
	if err != nil {
		t.Fatalf("CountByLevel(\"\") returned an error: %v", err)
	}
	if got == nil || len(got) != 0 {
		t.Fatalf("CountByLevel(\"\") = %#v, want an empty map", got)
	}
}

func TestCountByLevelMalformed(t *testing.T) {
	in := "{\"level\":\"info\"}\n{\"level\": oops}\n{\"level\":\"info\"}\n"
	if _, err := CountByLevel(strings.NewReader(in)); err == nil {
		t.Fatal("CountByLevel returned no error for malformed input")
	}
}
