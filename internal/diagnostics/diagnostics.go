// Package diagnostics parses and rewrites Go tool output so that errors point
// at the learner's files instead of the temporary staging directory.
package diagnostics

import (
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Diag is one file:line:col diagnostic from the Go toolchain.
type Diag struct {
	File    string `json:"file"`
	Line    int    `json:"line"`
	Col     int    `json:"col,omitempty"`
	Message string `json:"message"`
}

var diagRE = regexp.MustCompile(`^\s*(\S+?\.go):(\d+)(?::(\d+))?: (.+)$`)

// Parse extracts file:line[:col]: message diagnostics from compiler or vet
// output. Lines that do not look like diagnostics are ignored.
func Parse(text string) []Diag {
	var out []Diag
	for line := range strings.SplitSeq(text, "\n") {
		m := diagRE.FindStringSubmatch(strings.TrimRight(line, "\r"))
		if m == nil {
			continue
		}
		d := Diag{File: m[1], Message: m[4]}
		d.Line, _ = strconv.Atoi(m[2])
		if m[3] != "" {
			d.Col, _ = strconv.Atoi(m[3])
		}
		out = append(out, d)
	}
	return out
}

// Rewriter maps staging-directory paths to learner-facing paths.
type Rewriter struct {
	StagingDir string          // absolute staging directory
	DisplayDir string          // directory shown to the learner, e.g. exercises/01_x/id
	Learner    map[string]bool // base names of learner-owned files
}

// Rewrite rewrites tool output. References to learner files ("./main.go:3:1"
// or "/tmp/golings-123/main.go:3:1") become DisplayDir paths; the synthetic
// package header and summary lines ("# golings.local/exercise",
// "FAIL\tgolings.local/exercise [build failed]") are dropped. References to
// canonical check files are left as bare names.
func (r Rewriter) Rewrite(text string) string {
	if text == "" {
		return text
	}
	var b strings.Builder
	for _, line := range strings.SplitAfter(text, "\n") {
		trimmed := strings.TrimRight(line, "\r\n")
		if strings.HasPrefix(trimmed, "# golings.local/") || strings.HasPrefix(trimmed, "FAIL\tgolings.local/") ||
			strings.HasPrefix(trimmed, "ok  \tgolings.local/") {
			continue
		}
		b.WriteString(r.rewriteLine(line))
	}
	return b.String()
}

func (r Rewriter) rewriteLine(line string) string {
	if r.StagingDir != "" {
		for _, prefix := range []string{r.StagingDir + string(filepath.Separator), filepath.ToSlash(r.StagingDir) + "/"} {
			line = r.replacePrefixed(line, prefix)
		}
	}
	// The go command prints ".\name.go" on Windows and "./name.go" elsewhere.
	return r.replacePrefixed(r.replacePrefixed(line, "./"), ".\\")
}

// replacePrefixed rewrites every occurrence of prefix+name.go where name.go is
// a learner file. Other files (canonical checks) lose the prefix only.
func (r Rewriter) replacePrefixed(line, prefix string) string {
	var b strings.Builder
	for {
		i := strings.Index(line, prefix)
		if i < 0 {
			b.WriteString(line)
			return b.String()
		}
		// "./" and ".\\" must start a token, so we don't rewrite "a./b".
		if (prefix == "./" || prefix == ".\\") && i > 0 && !isBoundary(line[i-1]) {
			b.WriteString(line[:i+len(prefix)])
			line = line[i+len(prefix):]
			continue
		}
		b.WriteString(line[:i])
		rest := line[i+len(prefix):]
		end := strings.IndexAny(rest, ": \t\n\r)")
		if end < 0 {
			end = len(rest)
		}
		name := rest[:end]
		switch {
		case r.Learner[name] && r.DisplayDir != "":
			b.WriteString(filepath.Join(r.DisplayDir, name))
		default:
			b.WriteString(name)
		}
		line = rest[end:]
	}
}

func isBoundary(c byte) bool {
	return c == ' ' || c == '\t' || c == '(' || c == '"' || c == '\''
}

// Truncate returns the last limit bytes of s (cut at a line boundary when
// possible) and whether truncation happened.
func Truncate(s string, limit int) (string, bool) {
	if len(s) <= limit {
		return s, false
	}
	s = s[len(s)-limit:]
	if i := strings.IndexByte(s, '\n'); i >= 0 && i < len(s)-1 {
		s = s[i+1:]
	}
	return s, true
}
