package validate

import (
	"context"
	"fmt"
	"io/fs"
	"maps"
	"path"
	"sort"

	"github.com/maskedsyntax/golings/internal/catalog"
	"github.com/maskedsyntax/golings/internal/runner"
	"github.com/maskedsyntax/golings/internal/toolchain"
)

// Variant is one evaluated version of an exercise in a contract check.
type Variant struct {
	Name    string                 `json:"name"` // "starter", "solution", or "mutant:<name>"
	OK      bool                   `json:"ok"`
	Problem string                 `json:"problem,omitempty"`
	Result  *runner.ExerciseResult `json:"result,omitempty"`
}

// ContractResult is the contract outcome for one exercise.
type ContractResult struct {
	ExerciseID string    `json:"exercise_id"`
	OK         bool      `json:"ok"`
	Variants   []Variant `json:"variants"`
}

// Contract verifies that the starter does not pass (and fails in the
// documented way), the reference solution passes, and every mutant fails.
func Contract(ctx context.Context, cat *catalog.Catalog, ex *catalog.Exercise, tc *toolchain.Info) *ContractResult {
	cr := &ContractResult{ExerciseID: ex.ID, OK: true}
	add := func(v Variant) {
		if !v.OK {
			cr.OK = false
		}
		cr.Variants = append(cr.Variants, v)
	}
	assets, err := cat.Assets(ex)
	if err != nil {
		add(Variant{Name: "assets", Problem: err.Error()})
		return cr
	}
	if !ex.SupportsOS(tc.GOOS) {
		add(Variant{Name: "platform", OK: true, Problem: "skipped: not supported on " + tc.GOOS})
		return cr
	}
	eval := func(files map[string][]byte) *runner.ExerciseResult {
		return runner.Run(ctx, runner.Request{Exercise: ex, Files: files, Assets: assets, Toolchain: tc})
	}

	// Starter.
	starter, err := readFiles(assets, "starter", ex.StarterFiles)
	if err != nil {
		add(Variant{Name: "starter", Problem: err.Error()})
	} else {
		res := eval(starter)
		v := Variant{Name: "starter", Result: res, OK: true}
		first := res.FirstProblem()
		switch {
		case res.Status == runner.StatusPass:
			v.OK, v.Problem = false, "starter passes every check; it must fail until the learner fixes it"
		case res.Status == runner.StatusRunnerError || res.Status == runner.StatusTimeout:
			v.OK, v.Problem = false, fmt.Sprintf("starter produced %s instead of an ordinary failure", res.Status)
		case ex.StarterExpect == catalog.StarterExpectCompileError && (first == nil || first.Failure != runner.FailureCompile):
			v.OK, v.Problem = false, "starter_expect is compile_error but the starter compiles"
		case ex.StarterExpect == catalog.StarterExpectCheckFailure && first != nil && first.Failure == runner.FailureCompile:
			v.OK, v.Problem = false, "starter_expect is check_failure but the starter does not compile:\n"+first.Stderr
		}
		add(v)
	}

	// Solution.
	solution, err := readFiles(assets, "solution", ex.StarterFiles)
	if err != nil {
		add(Variant{Name: "solution", Problem: err.Error()})
		return cr
	}
	res := eval(solution)
	v := Variant{Name: "solution", Result: res, OK: res.Passed()}
	if !v.OK {
		v.Problem = "reference solution does not pass: " + describe(res)
	}
	add(v)

	// Mutants: solution overlaid with the mutant's files.
	entries, _ := fs.ReadDir(assets, "mutants")
	names := make([]string, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() {
			names = append(names, e.Name())
		}
	}
	sort.Strings(names)
	for _, name := range names {
		files := maps.Clone(solution)
		dir := path.Join("mutants", name)
		overlay, err := fs.ReadDir(assets, dir)
		for _, f := range overlay {
			if err != nil {
				break
			}
			var data []byte
			if data, err = fs.ReadFile(assets, path.Join(dir, f.Name())); err == nil {
				files[f.Name()] = data
			}
		}
		if err != nil {
			add(Variant{Name: "mutant:" + name, Problem: err.Error()})
			continue
		}
		res := eval(files)
		v := Variant{Name: "mutant:" + name, Result: res, OK: res.Status == runner.StatusFail}
		if !v.OK {
			v.Problem = fmt.Sprintf("mutant should fail the checks but got %s", res.Status)
		}
		add(v)
	}
	return cr
}

func readFiles(assets fs.FS, dir string, names []string) (map[string][]byte, error) {
	files := map[string][]byte{}
	for _, n := range names {
		data, err := fs.ReadFile(assets, path.Join(dir, n))
		if err != nil {
			return nil, err
		}
		files[n] = data
	}
	return files, nil
}

func describe(res *runner.ExerciseResult) string {
	c := res.FirstProblem()
	if c == nil {
		return string(res.Status)
	}
	msg := fmt.Sprintf("%s check: %s %s", c.Name, c.Status, c.Diagnostic)
	if c.Stderr != "" {
		msg += "\n" + c.Stderr
	}
	for _, t := range c.Tests {
		if t.Status == runner.StatusFail {
			msg += fmt.Sprintf("\n  FAIL %s\n%s", t.Name, t.Output)
		}
	}
	return msg
}
