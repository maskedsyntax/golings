package validate_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/maskedsyntax/golings/curriculum"
	"github.com/maskedsyntax/golings/internal/catalog"
	"github.com/maskedsyntax/golings/internal/toolchain"
	"github.com/maskedsyntax/golings/internal/validate"
)

func load(t *testing.T) *catalog.Catalog {
	t.Helper()
	cat, err := catalog.Load(curriculum.FS)
	if err != nil {
		t.Fatal(err)
	}
	return cat
}

func TestCurriculumIntegrity(t *testing.T) {
	rep := validate.Curriculum(load(t))
	for _, e := range rep.Errors {
		t.Error(e)
	}
}

func TestCoverageDocIsCurrent(t *testing.T) {
	want := validate.CoverageMarkdown(load(t))
	got, err := os.ReadFile(filepath.Join("..", "..", "docs", "curriculum-coverage.md"))
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want {
		t.Error("docs/curriculum-coverage.md is out of date; run `go run ./cmd/golings-dev coverage --write`")
	}
}

// TestCurriculumContract is the mandatory per-exercise contract (spec §10.3):
// every starter fails, every solution passes, every mutant fails.
func TestCurriculumContract(t *testing.T) {
	if testing.Short() {
		t.Skip("contract checks compile every exercise; skipped in -short mode")
	}
	tc, err := toolchain.Detect(context.Background())
	if err != nil {
		t.Skipf("Go toolchain unavailable: %v", err)
	}
	cat := load(t)
	for _, ex := range cat.Exercises() {
		t.Run(ex.ID, func(t *testing.T) {
			t.Parallel()
			res := validate.Contract(context.Background(), cat, ex, tc)
			for _, v := range res.Variants {
				if !v.OK {
					t.Errorf("%s: %s", v.Name, v.Problem)
				}
			}
		})
	}
}
