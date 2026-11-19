package config

import (
	"flag"
	"slices"
	"testing"
	"time"
)

var _ flag.Value = (*TagList)(nil)

func TestParse(t *testing.T) {
	cfg, err := Parse([]string{"-tag", "a", "-v", "-tag", "b", "-timeout", "2m30s", "x.txt", "y.txt"})
	if err != nil {
		t.Fatalf("Parse returned an error: %v", err)
	}
	if !slices.Equal(cfg.Tags, TagList{"a", "b"}) {
		t.Errorf("Tags = %q, want [a b]", cfg.Tags)
	}
	if cfg.Timeout != 150*time.Second {
		t.Errorf("Timeout = %v, want 2m30s", cfg.Timeout)
	}
	if !cfg.Verbose {
		t.Error("Verbose = false, want true")
	}
	if !slices.Equal(cfg.Files, []string{"x.txt", "y.txt"}) {
		t.Errorf("Files = %q, want [x.txt y.txt]", cfg.Files)
	}
}

func TestParseDefaults(t *testing.T) {
	cfg, err := Parse(nil)
	if err != nil {
		t.Fatalf("Parse(nil) returned an error: %v", err)
	}
	if len(cfg.Tags) != 0 || cfg.Timeout != 5*time.Second || cfg.Verbose || len(cfg.Files) != 0 {
		t.Errorf("Parse(nil) = %+v; want no tags, 5s, not verbose, no files", cfg)
	}
}

func TestParseErrors(t *testing.T) {
	for _, args := range [][]string{
		{"-timeout", "soon"},
		{"-tag", ""},
		{"-nope"},
		{"-tag"},
	} {
		if _, err := Parse(args); err == nil {
			t.Errorf("Parse(%q) returned no error", args)
		}
	}
}

func TestTagListString(t *testing.T) {
	tags := TagList{"a", "b"}
	if got := tags.String(); got != "a,b" {
		t.Errorf("TagList{a b}.String() = %q, want \"a,b\"", got)
	}
}
