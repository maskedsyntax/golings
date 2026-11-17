package diagnostics

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	out := "# golings.local/exercise\n./values.go:18:9: cannot use x\nvalues_test.go:12: got 1\nnoise\n"
	got := Parse(out)
	if len(got) != 2 {
		t.Fatalf("Parse = %+v", got)
	}
	if got[0] != (Diag{File: "./values.go", Line: 18, Col: 9, Message: "cannot use x"}) {
		t.Errorf("first = %+v", got[0])
	}
	if got[1].File != "values_test.go" || got[1].Line != 12 || got[1].Col != 0 {
		t.Errorf("second = %+v", got[1])
	}
}

func TestRewrite(t *testing.T) {
	staging := filepath.Join(string(filepath.Separator)+"tmp", "golings-123")
	r := Rewriter{
		StagingDir: staging,
		DisplayDir: filepath.Join("exercises", "01_fundamentals", "values_01"),
		Learner:    map[string]bool{"values.go": true},
	}
	in := "# golings.local/exercise\n" +
		"./values.go:18:9: cannot use x\n" +
		staging + string(filepath.Separator) + "values.go:3:1: other\n" +
		"./values_test.go:4: check file\n" +
		"FAIL\tgolings.local/exercise [build failed]\n"
	got := r.Rewrite(in)
	want := filepath.Join(r.DisplayDir, "values.go") + ":18:9: cannot use x\n" +
		filepath.Join(r.DisplayDir, "values.go") + ":3:1: other\n" +
		"values_test.go:4: check file\n"
	if got != want {
		t.Errorf("Rewrite =\n%s\nwant\n%s", got, want)
	}
}

func TestTruncate(t *testing.T) {
	s := strings.Repeat("line\n", 100)
	got, cut := Truncate(s, 52)
	if !cut || len(got) > 52 || !strings.HasPrefix(got, "line") {
		t.Errorf("Truncate = %q, %v", got, cut)
	}
	if got, cut := Truncate("short", 10); cut || got != "short" {
		t.Errorf("Truncate(short) = %q, %v", got, cut)
	}
}
