// Package render formats golings output for humans (with optional color)
// and for scripts (JSON without ANSI escapes).
package render

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/maskedsyntax/golings/internal/catalog"
	"github.com/maskedsyntax/golings/internal/diagnostics"
	"github.com/maskedsyntax/golings/internal/runner"
)

// Style applies ANSI styling when enabled.
type Style struct{ Color bool }

func (s Style) wrap(code, text string) string {
	if !s.Color || text == "" {
		return text
	}
	return "\x1b[" + code + "m" + text + "\x1b[0m"
}

func (s Style) Bold(t string) string   { return s.wrap("1", t) }
func (s Style) Dim(t string) string    { return s.wrap("2", t) }
func (s Style) Red(t string) string    { return s.wrap("31", t) }
func (s Style) Green(t string) string  { return s.wrap("32", t) }
func (s Style) Yellow(t string) string { return s.wrap("33", t) }
func (s Style) Cyan(t string) string   { return s.wrap("36", t) }

// IsTerminal reports whether f is an interactive terminal.
func IsTerminal(f *os.File) bool {
	fi, err := f.Stat()
	return err == nil && fi.Mode()&os.ModeCharDevice != 0
}

// ColorEnabled decides whether to emit ANSI color: only for a terminal, and
// never when NO_COLOR is set or --no-color was passed.
func ColorEnabled(w io.Writer, noColor bool) bool {
	if noColor || os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb" {
		return false
	}
	f, ok := w.(*os.File)
	return ok && IsTerminal(f)
}

// JSON writes v as indented JSON followed by a newline.
func JSON(w io.Writer, v any) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}

// ExerciseHeader prints the exercise identity block.
func ExerciseHeader(w io.Writer, s Style, ex *catalog.Exercise, total int, topicTitle, relDir string) {
	fmt.Fprintf(w, "%s  %s  •  %s\n", s.Bold(fmt.Sprintf("Exercise %d/%d", ex.Ordinal, total)), s.Cyan(ex.ID), topicTitle)
	fmt.Fprintf(w, "Goal:   %s\n", ex.Title)
	files := make([]string, len(ex.StarterFiles))
	for i, f := range ex.StarterFiles {
		files[i] = filepath.Join(relDir, f)
	}
	fmt.Fprintf(w, "Path:   %s\n", strings.Join(files, ", "))
	fmt.Fprintf(w, "Guide:  %s\n", filepath.Join(relDir, "README.md"))
	fmt.Fprintf(w, "Source: %s\n", ex.URL)
}

// failureLabel names the failure category in learner terms.
func failureLabel(c *runner.CheckResult) string {
	switch c.Failure {
	case runner.FailureCompile:
		return "Compiler error"
	case runner.FailureTest:
		return "Test failure"
	case runner.FailureOutput:
		return "Output mismatch"
	case runner.FailureVet:
		return "go vet problem"
	case runner.FailureRace:
		return "Data race"
	case runner.FailureTimeout:
		return "Timeout"
	case runner.FailureRunner:
		return "Runner error"
	}
	if c.Status == runner.StatusUnsupported {
		return "Unsupported"
	}
	return "Problem"
}

// Result prints an evaluation result. With verbose, passing tests and
// captured program output are listed too.
func Result(w io.Writer, s Style, res *runner.ExerciseResult, verbose bool) {
	for i := range res.Checks {
		c := &res.Checks[i]
		if c.Status == runner.StatusPass {
			if verbose || c.Name == string(catalog.CheckOutput) {
				printPassingCheck(w, s, c, verbose)
			}
			continue
		}
		printProblem(w, s, c)
	}
	fmt.Fprintln(w)
	switch res.Status {
	case runner.StatusPass:
		n := 0
		for _, c := range res.Checks {
			if len(c.Tests) > 0 {
				n += len(c.Tests)
			} else {
				n++
			}
		}
		msg := fmt.Sprintf("✓ All %d checks passed", n)
		if n == 1 {
			msg = "✓ Check passed"
		}
		fmt.Fprintf(w, "%s\n", s.Green(msg))
	case runner.StatusUnsupported:
		fmt.Fprintf(w, "%s\n", s.Yellow("! This exercise cannot be verified on this platform"))
	case runner.StatusTimeout:
		fmt.Fprintf(w, "%s\n", s.Red("✗ Timed out"))
	case runner.StatusRunnerError:
		fmt.Fprintf(w, "%s\n", s.Red("✗ golings could not evaluate this exercise"))
	default:
		fmt.Fprintf(w, "%s\n", s.Red("✗ Not passing yet"))
	}
}

func printPassingCheck(w io.Writer, s Style, c *runner.CheckResult, verbose bool) {
	if c.Name == string(catalog.CheckOutput) && c.Stdout != "" {
		fmt.Fprintf(w, "%s\n", s.Dim("Program output:"))
		fmt.Fprint(w, indent(strings.TrimRight(c.Stdout, "\n"), "  "), "\n")
		return
	}
	if verbose {
		for _, t := range c.Tests {
			fmt.Fprintf(w, "%s  %s\n", s.Green("PASS"), t.Name)
		}
		if len(c.Tests) == 0 {
			fmt.Fprintf(w, "%s  %s\n", s.Green("PASS"), c.Name)
		}
	}
}

func printProblem(w io.Writer, s Style, c *runner.CheckResult) {
	label := failureLabel(c)
	color := s.Red
	if c.Status == runner.StatusUnsupported {
		color = s.Yellow
	}
	fmt.Fprintf(w, "%s %s\n", color(label+":"), c.Diagnostic)

	switch c.Failure {
	case runner.FailureCompile, runner.FailureVet:
		if out := strings.TrimSpace(c.Stderr); out != "" {
			fmt.Fprintln(w, indent(out, "  "))
		}
	case runner.FailureTest:
		passed := 0
		for _, t := range c.Tests {
			switch t.Status {
			case runner.StatusFail:
				fmt.Fprintf(w, "%s  %s\n", s.Red("FAIL"), t.Name)
				if t.Output != "" {
					fmt.Fprintln(w, indent(dedent(t.Output), "      "))
				}
			case runner.StatusPass:
				passed++
			}
		}
		if len(c.Tests) > 0 {
			fmt.Fprintf(w, "%s\n", s.Dim(fmt.Sprintf("%d of %d checks passing", passed, len(c.Tests))))
		}
		if out := strings.TrimSpace(c.Stderr); out != "" {
			fmt.Fprintln(w, indent(out, "  "))
		}
	case runner.FailureOutput:
		if c.Stdout != "" {
			fmt.Fprintf(w, "%s\n%s\n", s.Dim("Your program printed:"), indent(strings.TrimRight(c.Stdout, "\n"), "  "))
		} else {
			fmt.Fprintf(w, "%s\n", s.Dim("Your program printed nothing."))
		}
		if c.Expected != "" {
			fmt.Fprintf(w, "%s\n%s\n", s.Dim("Expected:"), indent(c.Expected, "  "))
		}
		if out := strings.TrimSpace(c.Stderr); out != "" {
			fmt.Fprintf(w, "%s\n%s\n", s.Dim("Standard error:"), indent(out, "  "))
		}
	default:
		if out := strings.TrimSpace(c.Stderr); out != "" {
			out, cut := diagnostics.Truncate(out, 8<<10)
			if cut {
				out = "…(earlier output truncated)\n" + out
			}
			fmt.Fprintln(w, indent(out, "  "))
		}
	}
	if c.Truncated {
		fmt.Fprintf(w, "%s\n", s.Dim("(output was truncated)"))
	}
}

func indent(text, prefix string) string {
	lines := strings.Split(text, "\n")
	for i, l := range lines {
		if l != "" {
			lines[i] = prefix + l
		}
	}
	return strings.Join(lines, "\n")
}

// dedent removes the common leading whitespace go test adds to log lines.
func dedent(text string) string {
	lines := strings.Split(text, "\n")
	common := -1
	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}
		n := len(l) - len(strings.TrimLeft(l, " \t"))
		if common < 0 || n < common {
			common = n
		}
	}
	if common <= 0 {
		return text
	}
	for i, l := range lines {
		if len(l) >= common {
			lines[i] = l[common:]
		}
	}
	return strings.Join(lines, "\n")
}
