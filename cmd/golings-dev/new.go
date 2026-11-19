package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/maskedsyntax/golings/internal/catalog"
)

// newCmd scaffolds the next exercise for a topic: metadata, README, a starter
// that fails, a test, a solution that passes and one mutant, so the new
// exercise satisfies the contract from the start. Authors then replace the
// placeholder content. The ID is added to manifest.json after the topic's
// existing exercises.
func newCmd(cat *catalog.Catalog, args []string, stdout, stderr io.Writer) int {
	if len(args) != 1 {
		fmt.Fprintln(stderr, "usage: golings-dev new <topic>")
		return 2
	}
	topic := cat.Topic(args[0])
	if topic == nil {
		fmt.Fprintf(stderr, "unknown topic %q (see curriculum/topics.json)\n", args[0])
		return 2
	}
	id, last := nextExerciseID(cat, topic.ID)
	dir := filepath.Join("curriculum", catalog.ExercisesDir, topic.Module, id)
	if _, err := os.Stat(dir); err == nil {
		fmt.Fprintf(stderr, "%s already exists\n", dir)
		return 1
	}
	manifest, err := os.ReadFile(filepath.Join("curriculum", "manifest.json"))
	if err != nil {
		fmt.Fprintf(stderr, "run golings-dev from the repository root: %v\n", err)
		return 1
	}
	updated, err := insertIntoManifest(string(manifest), topic.Module, last, id)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}

	pkg := strings.ReplaceAll(topic.ID, "_", "")
	file := topic.ID + ".go"
	prereqs := "[]"
	if last != "" {
		prereqs = `["` + last + `"]`
	}
	files := map[string]string{
		"exercise.json": fmt.Sprintf(`{
  "schema_version": 1,
  "id": %q,
  "topic": %q,
  "title": "TODO: a short title",
  "type": "implement",
  "difficulty": "beginner",
  "estimated_minutes": 10,
  "prerequisites": %s,
  "starter_files": [%q],
  "fixture_files": [],
  "checks": ["unit"],
  "timeout_seconds": 10,
  "platforms": [],
  "starter_expect": "check_failure",
  "hints": [
    "TODO: a nudge in the right direction.",
    "TODO: name the API or idea to use.",
    "TODO: nearly the answer."
  ]
}
`, id, topic.ID, prereqs, file),
		"README.md": fmt.Sprintf("# %s: TODO title\n\nTODO: explain the idea with a short example.\n\n## Your task\n\nTODO\n\nReference: %s\n", id, topic.URL),
		filepath.Join("starter", file): fmt.Sprintf(`// %s: TODO title
//
// TODO: one line saying what to do. See README.md.
//
// Reference: %s
package %s

// Answer returns the answer.
func Answer() int {
	// TODO
	return 0
}
`, id, topic.URL, pkg),
		filepath.Join("solution", file): fmt.Sprintf("// %s (reference solution)\npackage %s\n\n// Answer returns the answer.\nfunc Answer() int {\n\treturn 42\n}\n", id, pkg),
		filepath.Join("checks", topic.ID+"_test.go"): fmt.Sprintf(`package %s

import "testing"

func TestAnswer(t *testing.T) {
	if got := Answer(); got != 42 {
		t.Fatalf("Answer() = %%d, want 42", got)
	}
}
`, pkg),
		filepath.Join("mutants", "off_by_one", file): fmt.Sprintf("// Mutant: a plausible wrong answer.\npackage %s\n\nfunc Answer() int {\n\treturn 41\n}\n", pkg),
	}
	for rel, content := range files {
		p := filepath.Join(dir, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
	}
	if err := os.WriteFile(filepath.Join("curriculum", "manifest.json"), []byte(updated), 0o644); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	fmt.Fprintf(stdout, "created %s and added %s to manifest.json\n\nnext:\n", dir, id)
	fmt.Fprintf(stdout, "  edit the files (see docs/authoring.md)\n  go run ./cmd/golings-dev check %s\n  go run ./cmd/golings-dev coverage --write\n", id)
	return 0
}

var idNumberRE = regexp.MustCompile(`_(\d{2})$`)

// nextExerciseID returns the next free topic_NN and the topic's last
// existing exercise ("" if none).
func nextExerciseID(cat *catalog.Catalog, topic string) (id, last string) {
	n := 0
	for _, ex := range cat.Exercises() {
		if ex.Topic != topic {
			continue
		}
		last = ex.ID
		if m := idNumberRE.FindStringSubmatch(ex.ID); m != nil {
			v, _ := strconv.Atoi(m[1])
			n = max(n, v)
		}
	}
	return fmt.Sprintf("%s_%02d", topic, n+1), last
}

// insertIntoManifest adds id after `after` (or as the first exercise of
// module when after is ""), editing the text so the hand-kept layout of
// manifest.json survives.
func insertIntoManifest(manifest, module, after, id string) (string, error) {
	if after != "" {
		quoted := `"` + after + `"`
		if strings.Count(manifest, quoted) != 1 {
			return "", fmt.Errorf("manifest.json: cannot find %s exactly once", quoted)
		}
		return strings.Replace(manifest, quoted, quoted+`, "`+id+`"`, 1), nil
	}
	re := regexp.MustCompile(`("id": "` + regexp.QuoteMeta(module) + `",[^\[]*"exercises": \[)(\s*)(\]?)`)
	loc := re.FindStringSubmatchIndex(manifest)
	if loc == nil {
		return "", fmt.Errorf("manifest.json: module %s not found", module)
	}
	head, closing := manifest[:loc[3]], manifest[loc[6]:loc[7]]
	rest := manifest[loc[7]:]
	if closing == "]" {
		return head + `"` + id + `"]` + rest, nil
	}
	return head + "\n        \"" + id + "\"," + manifest[loc[4]:loc[5]] + rest, nil
}
