package catalog_test

import (
	"errors"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/maskedsyntax/golings/curriculum"
	"github.com/maskedsyntax/golings/internal/catalog"
)

func TestEmbeddedCurriculumLoads(t *testing.T) {
	cat, err := catalog.Load(curriculum.FS)
	if err != nil {
		t.Fatal(err)
	}
	if got := len(cat.Topics.Topics); got != 85 {
		t.Errorf("topics = %d, want 85", got)
	}
	for i, ex := range cat.Exercises() {
		if ex.Ordinal != i+1 {
			t.Errorf("%s ordinal = %d, want %d", ex.ID, ex.Ordinal, i+1)
		}
		if !strings.HasPrefix(ex.URL, catalog.TopicURLPrefix) {
			t.Errorf("%s has url %q", ex.ID, ex.URL)
		}
	}
}

const topicsJSON = `{"schema_version":1,"source":"x","snapshot_date":"2026-11-17","topics":[
 {"id":"slices","title":"Slices","url":"https://gobyexample.com/slices","module":"01_basics","source_ordinal":1},
 {"id":"maps","title":"Maps","url":"https://gobyexample.com/maps","module":"01_basics","source_ordinal":2}]}`

const manifestJSON = `{"schema_version":1,"curriculum_version":"0.0.1","release_complete":false,"modules":[
 {"id":"01_basics","title":"Basics","exercises":["slices_01","slices_02","maps_01"]}]}`

func exerciseJSON(id, topic string, extra string) string {
	return `{"schema_version":1,"id":"` + id + `","topic":"` + topic + `","title":"T","type":"implement",
 "difficulty":"beginner","estimated_minutes":5,"prerequisites":[],"starter_files":["a.go"],"fixture_files":[],
 "checks":["unit"],"timeout_seconds":10,"platforms":[],"hints":["h"],"starter_expect":""` + extra + `}`
}

func validFS() fstest.MapFS {
	return fstest.MapFS{
		"manifest.json": {Data: []byte(manifestJSON)},
		"topics.json":   {Data: []byte(topicsJSON)},
		"_exercises/01_basics/slices_01/exercise.json": {Data: []byte(exerciseJSON("slices_01", "slices", ""))},
		"_exercises/01_basics/slices_02/exercise.json": {Data: []byte(exerciseJSON("slices_02", "slices", ""))},
		"_exercises/01_basics/maps_01/exercise.json":   {Data: []byte(exerciseJSON("maps_01", "maps", ""))},
	}
}

func TestLoadValid(t *testing.T) {
	cat, err := catalog.Load(validFS())
	if err != nil {
		t.Fatal(err)
	}
	ex, ok := cat.Get("maps_01")
	if !ok || ex.Ordinal != 3 || ex.Module != "01_basics" || ex.URL != "https://gobyexample.com/maps" {
		t.Fatalf("maps_01 = %+v", ex)
	}
}

func TestLoadRejectsInvalidMetadata(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(fstest.MapFS)
		want   string
	}{
		{"unknown field", func(f fstest.MapFS) {
			f["_exercises/01_basics/maps_01/exercise.json"].Data = []byte(exerciseJSON("maps_01", "maps", `,"bogus":1`))
		}, "unknown field"},
		{"id mismatch topic", func(f fstest.MapFS) {
			f["_exercises/01_basics/maps_01/exercise.json"].Data = []byte(exerciseJSON("maps_01", "slices", ""))
		}, "must start with its topic"},
		{"bad check kind", func(f fstest.MapFS) {
			f["_exercises/01_basics/maps_01/exercise.json"].Data = []byte(strings.Replace(exerciseJSON("maps_01", "maps", ""), `["unit"]`, `["sh -c rm"]`, 1))
		}, "unknown check kind"},
		{"path in starter", func(f fstest.MapFS) {
			f["_exercises/01_basics/maps_01/exercise.json"].Data = []byte(strings.Replace(exerciseJSON("maps_01", "maps", ""), `["a.go"]`, `["../a.go"]`, 1))
		}, "invalid starter file"},
		{"later prerequisite", func(f fstest.MapFS) {
			f["_exercises/01_basics/slices_01/exercise.json"].Data = []byte(strings.Replace(exerciseJSON("slices_01", "slices", ""), `"prerequisites":[]`, `"prerequisites":["maps_01"]`, 1))
		}, "must come earlier"},
		{"wrong url", func(f fstest.MapFS) {
			f["topics.json"].Data = []byte(strings.Replace(topicsJSON, "gobyexample.com/maps", "gobyexample.com/map", 1))
		}, "want \"https://gobyexample.com/maps\""},
		{"duplicate id", func(f fstest.MapFS) {
			f["manifest.json"].Data = []byte(strings.Replace(manifestJSON, `"maps_01"]`, `"maps_01","maps_01"]`, 1))
		}, "duplicate exercise id"},
		{"missing exercise.json", func(f fstest.MapFS) {
			delete(f, "_exercises/01_basics/maps_01/exercise.json")
		}, "maps_01/exercise.json"},
		{"too many hints", func(f fstest.MapFS) {
			f["_exercises/01_basics/maps_01/exercise.json"].Data = []byte(strings.Replace(exerciseJSON("maps_01", "maps", ""), `["h"]`, `["a","b","c","d"]`, 1))
		}, "1 to 3 hints"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsys := validFS()
			tt.mutate(fsys)
			_, err := catalog.Load(fsys)
			var ve *catalog.ValidationError
			if !errors.As(err, &ve) {
				t.Fatalf("Load error = %v, want *ValidationError", err)
			}
			if !strings.Contains(err.Error(), tt.want) {
				t.Errorf("error %q does not mention %q", err, tt.want)
			}
		})
	}
}

func TestFind(t *testing.T) {
	cat, err := catalog.Load(validFS())
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range []struct{ q, want string }{
		{"maps_01", "maps_01"},
		{"MAPS_01", "maps_01"},
		{"3", "maps_01"},
		{"map", "maps_01"},
		{"slices_02", "slices_02"},
	} {
		ex, err := cat.Find(tt.q)
		if err != nil || ex.ID != tt.want {
			t.Errorf("Find(%q) = %v, %v; want %s", tt.q, ex, err, tt.want)
		}
	}

	var amb *catalog.AmbiguousError
	if _, err := cat.Find("slices"); !errors.As(err, &amb) || len(amb.Matches) != 2 {
		t.Errorf("Find(slices) error = %v, want ambiguous with 2 matches", err)
	}
	var nf *catalog.NotFoundError
	if _, err := cat.Find("slics_01"); !errors.As(err, &nf) || len(nf.Suggestions) == 0 || nf.Suggestions[0] != "slices_01" {
		t.Errorf("Find(slics_01) error = %v, want not found suggesting slices_01", err)
	}
	if _, err := cat.Find("99"); !errors.As(err, &nf) {
		t.Errorf("Find(99) error = %v, want not found", err)
	}
}

func TestRenamesFollowChains(t *testing.T) {
	f := validFS()
	f["manifest.json"].Data = []byte(strings.Replace(manifestJSON, `"slices_02"`, `"slices_03"`, 1))
	delete(f, "_exercises/01_basics/slices_02/exercise.json")
	f["_exercises/01_basics/slices_03/exercise.json"] = &fstest.MapFile{Data: []byte(exerciseJSON("slices_03", "slices", ""))}
	f["migrations.json"] = &fstest.MapFile{Data: []byte(`{"schema_version":1,"migrations":[
		{"to_version":"0.0.2","renamed":{"maps_09":"slices_02"},"removed":["maps_08"]},
		{"to_version":"0.0.3","renamed":{"slices_02":"slices_03"}}]}`)}
	cat, err := catalog.Load(f)
	if err != nil {
		t.Fatal(err)
	}
	got := cat.Renames()
	want := map[string]string{"maps_09": "slices_03", "slices_02": "slices_03", "maps_08": ""}
	if len(got) != len(want) {
		t.Fatalf("Renames = %v, want %v", got, want)
	}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("Renames[%s] = %q, want %q", k, got[k], v)
		}
	}
}

func TestMigrationsAreValidated(t *testing.T) {
	for name, migrations := range map[string]string{
		"rename to unknown":     `{"to_version":"0.0.2","renamed":{"old_01":"nope_01"}}`,
		"renamed id still used": `{"to_version":"0.0.2","renamed":{"maps_01":"slices_01"}}`,
		"removed id still used": `{"to_version":"0.0.2","removed":["maps_01"]}`,
		"missing version":       `{"renamed":{"old_01":"maps_01"}}`,
	} {
		f := validFS()
		f["migrations.json"] = &fstest.MapFile{Data: []byte(`{"schema_version":1,"migrations":[` + migrations + `]}`)}
		if _, err := catalog.Load(f); err == nil {
			t.Errorf("%s: Load accepted %s", name, migrations)
		}
	}
}

func TestChecksHashTracksChecksOnly(t *testing.T) {
	hash := func(f fstest.MapFS) string {
		t.Helper()
		cat, err := catalog.Load(f)
		if err != nil {
			t.Fatal(err)
		}
		ex, _ := cat.Get("maps_01")
		return ex.ChecksHash
	}
	f := validFS()
	f["_exercises/01_basics/maps_01/checks/a_test.go"] = &fstest.MapFile{Data: []byte("package a\n")}
	base := hash(f)
	if base == "" {
		t.Fatal("empty checks hash")
	}
	f["_exercises/01_basics/maps_01/README.md"] = &fstest.MapFile{Data: []byte("reworded")}
	f["_exercises/01_basics/maps_01/exercise.json"].Data = []byte(strings.Replace(exerciseJSON("maps_01", "maps", ""), `"hints":["h"]`, `"hints":["a better hint"]`, 1))
	if got := hash(f); got != base {
		t.Error("changing a README or hint changed the checks hash")
	}
	f["_exercises/01_basics/maps_01/checks/a_test.go"].Data = []byte("package a\n// stricter\n")
	if got := hash(f); got == base {
		t.Error("changing a canonical check did not change the checks hash")
	}
}
