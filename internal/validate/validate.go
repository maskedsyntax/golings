// Package validate checks curriculum integrity (assets, coverage, layout)
// and runs the curriculum contract: starters must fail, solutions must pass,
// mutants must fail.
package validate

import (
	"errors"
	"fmt"
	"io/fs"
	"path"
	"slices"
	"sort"
	"strings"

	"github.com/maskedsyntax/golings/internal/catalog"
)

// MinExercisesPerTopic is the coverage requirement (spec §8).
const MinExercisesPerTopic = 2

// Report collects problems. Errors fail CI; warnings are informational.
type Report struct {
	Errors   []string `json:"errors"`
	Warnings []string `json:"warnings"`
}

func (r *Report) errorf(format string, args ...any) {
	r.Errors = append(r.Errors, fmt.Sprintf(format, args...))
}

func (r *Report) warnf(format string, args ...any) {
	r.Warnings = append(r.Warnings, fmt.Sprintf(format, args...))
}

// OK reports whether there are no errors.
func (r *Report) OK() bool { return len(r.Errors) == 0 }

// Curriculum validates asset layout and topic coverage on top of the
// structural checks catalog.Load already performed.
func Curriculum(cat *catalog.Catalog) *Report {
	r := &Report{}
	fsys := cat.FS()

	listed := map[string]bool{}
	for _, ex := range cat.Exercises() {
		listed[ex.AssetDir()] = true
		checkAssets(r, cat, ex)
	}

	// Orphans: exercise directories not listed in the manifest.
	modules, _ := fs.ReadDir(fsys, catalog.ExercisesDir)
	for _, m := range modules {
		if !m.IsDir() {
			continue
		}
		dirs, _ := fs.ReadDir(fsys, path.Join(catalog.ExercisesDir, m.Name()))
		for _, d := range dirs {
			p := path.Join(catalog.ExercisesDir, m.Name(), d.Name())
			if d.IsDir() && !listed[p] {
				r.errorf("%s: orphaned exercise directory (not listed in manifest.json)", p)
			}
		}
	}

	// Coverage.
	counts := TopicCounts(cat)
	for _, t := range cat.Topics.Topics {
		if n := counts[t.ID]; n < MinExercisesPerTopic {
			msg := fmt.Sprintf("topic %s has %d exercise(s), want at least %d", t.ID, n, MinExercisesPerTopic)
			if cat.Manifest.ReleaseComplete {
				r.Errors = append(r.Errors, msg)
			} else {
				r.Warnings = append(r.Warnings, msg)
			}
		}
	}
	sort.Strings(r.Errors)
	return r
}

func checkAssets(r *Report, cat *catalog.Catalog, ex *catalog.Exercise) {
	dir := ex.AssetDir()
	assets, err := cat.Assets(ex)
	if err != nil {
		r.errorf("%s: %v", dir, err)
		return
	}
	if _, err := fs.Stat(assets, "README.md"); err != nil {
		r.errorf("%s: missing README.md", dir)
	}
	starter := fileSet(assets, "starter")
	solution := fileSet(assets, "solution")
	for _, f := range ex.StarterFiles {
		if !starter[f] {
			r.errorf("%s: missing starter/%s", dir, f)
		}
		if !solution[f] {
			r.errorf("%s: missing solution/%s", dir, f)
		}
	}
	for f := range starter {
		if !contains(ex.StarterFiles, f) {
			r.errorf("%s: starter/%s is not listed in starter_files", dir, f)
		}
	}
	for f := range solution {
		if !contains(ex.StarterFiles, f) {
			r.errorf("%s: solution/%s is not listed in starter_files", dir, f)
		}
	}

	checks := fileSet(assets, "checks")
	var tests int
	for f := range checks {
		if strings.HasSuffix(f, "_test.go") {
			tests++
		}
		if contains(ex.StarterFiles, f) {
			r.errorf("%s: checks/%s collides with a starter file", dir, f)
		}
	}
	for _, k := range ex.Checks {
		switch k {
		case catalog.CheckUnit, catalog.CheckRace:
			if tests == 0 {
				r.errorf("%s: %q check declared but checks/ has no _test.go files", dir, k)
			}
		case catalog.CheckOutput:
			if !checks["expected_stdout.txt"] {
				r.errorf("%s: %q check declared but checks/expected_stdout.txt is missing", dir, k)
			}
		}
	}
	for _, f := range ex.FixtureFiles {
		if _, err := fs.Stat(assets, path.Join("fixtures", f)); err != nil {
			r.errorf("%s: missing fixtures/%s", dir, f)
		}
	}

	mutants, err := fs.ReadDir(assets, "mutants")
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		r.errorf("%s: reading mutants: %v", dir, err)
	}
	for _, m := range mutants {
		if !m.IsDir() {
			r.errorf("%s: mutants/%s must be a directory", dir, m.Name())
			continue
		}
		for f := range fileSet(assets, path.Join("mutants", m.Name())) {
			if !contains(ex.StarterFiles, f) {
				r.errorf("%s: mutants/%s/%s is not a starter file", dir, m.Name(), f)
			}
		}
	}
	if len(mutants) == 0 && (ex.Type == catalog.TypeDebug || ex.Type == catalog.TypeConcurrency) {
		r.warnf("%s: %s exercises should include at least one mutant", dir, ex.Type)
	}
}

// fileSet returns the regular files directly inside dir.
func fileSet(fsys fs.FS, dir string) map[string]bool {
	out := map[string]bool{}
	entries, _ := fs.ReadDir(fsys, dir)
	for _, e := range entries {
		if !e.IsDir() {
			out[e.Name()] = true
		}
	}
	return out
}

func contains(list []string, s string) bool { return slices.Contains(list, s) }

// TopicCounts returns the number of exercises per topic id.
func TopicCounts(cat *catalog.Catalog) map[string]int {
	counts := map[string]int{}
	for _, ex := range cat.Exercises() {
		counts[ex.Topic]++
	}
	return counts
}
