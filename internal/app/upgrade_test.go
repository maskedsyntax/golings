package app_test

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/maskedsyntax/golings/internal/app"
	"github.com/maskedsyntax/golings/internal/catalog"
	"github.com/maskedsyntax/golings/internal/progress"
	"github.com/maskedsyntax/golings/internal/workspace"
)

type fakeEx struct{ module, id, topic, readme string }

// fakeCurriculum builds a small curriculum FS. topics maps topic -> module.
func fakeCurriculum(t *testing.T, version string, topics [][2]string, exercises []fakeEx, migrations string) *catalog.Catalog {
	t.Helper()
	var tj []string
	for i, tp := range topics {
		url := "https://gobyexample.com/" + strings.ReplaceAll(tp[0], "_", "-")
		tj = append(tj, `{"id":"`+tp[0]+`","title":"`+tp[0]+`","url":"`+url+`","module":"`+tp[1]+`","source_ordinal":`+string(rune('1'+i))+`}`)
	}
	f := fstest.MapFS{
		"topics.json": {Data: []byte(`{"schema_version":1,"source":"x","snapshot_date":"2026-11-17","topics":[` + strings.Join(tj, ",") + `]}`)},
	}
	if migrations != "" {
		f["migrations.json"] = &fstest.MapFile{Data: []byte(`{"schema_version":1,"migrations":[` + migrations + `]}`)}
	}
	var modules []string
	byModule := map[string][]string{}
	for _, ex := range exercises {
		if _, ok := byModule[ex.module]; !ok {
			modules = append(modules, ex.module)
		}
		byModule[ex.module] = append(byModule[ex.module], `"`+ex.id+`"`)
		dir := "_exercises/" + ex.module + "/" + ex.id + "/"
		f[dir+"exercise.json"] = &fstest.MapFile{Data: []byte(`{"schema_version":1,"id":"` + ex.id + `","topic":"` + ex.topic +
			`","title":"T","type":"implement","difficulty":"beginner","estimated_minutes":5,"prerequisites":[],` +
			`"starter_files":["a.go"],"fixture_files":[],"checks":["unit"],"timeout_seconds":10,"platforms":[],"hints":["h"],"starter_expect":""}`)}
		f[dir+"README.md"] = &fstest.MapFile{Data: []byte(ex.readme)}
		f[dir+"starter/a.go"] = &fstest.MapFile{Data: []byte("package a // starter " + ex.id + "\n")}
		f[dir+"checks/a_test.go"] = &fstest.MapFile{Data: []byte("package a\n")}
	}
	var mj []string
	for _, m := range modules {
		mj = append(mj, `{"id":"`+m+`","title":"`+m+`","exercises":[`+strings.Join(byModule[m], ",")+`]}`)
	}
	f["manifest.json"] = &fstest.MapFile{Data: []byte(`{"schema_version":1,"curriculum_version":"` + version +
		`","release_complete":false,"modules":[` + strings.Join(mj, ",") + `]}`)}
	cat, err := catalog.Load(f)
	if err != nil {
		t.Fatal(err)
	}
	return cat
}

func read(t *testing.T, p string) string {
	t.Helper()
	data, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func TestUpgradeCarriesWorkAcrossVersions(t *testing.T) {
	v1 := fakeCurriculum(t, "0.0.1",
		[][2]string{{"slices", "01_basics"}, {"maps", "01_basics"}},
		[]fakeEx{{"01_basics", "slices_01", "slices", "old readme"}, {"01_basics", "maps_01", "maps", "maps readme"}}, "")
	// v2 renames maps_01 to maps_02, moves the maps topic to a new module,
	// adds structs_01 and rewords the slices README.
	v2 := fakeCurriculum(t, "0.0.2",
		[][2]string{{"slices", "01_basics"}, {"maps", "02_more"}, {"structs", "02_more"}},
		[]fakeEx{{"01_basics", "slices_01", "slices", "new readme"}, {"02_more", "maps_02", "maps", "maps readme"},
			{"02_more", "structs_01", "structs", "structs readme"}},
		`{"to_version":"0.0.2","renamed":{"maps_01":"maps_02"}}`)

	root := filepath.Join(t.TempDir(), "ws")
	ws, err := workspace.Init(root, v1, workspace.InitOptions{})
	if err != nil {
		t.Fatal(err)
	}
	mapsV1 := filepath.Join(root, "exercises", "01_basics", "maps_01", "a.go")
	slicesFile := filepath.Join(root, "exercises", "01_basics", "slices_01", "a.go")
	os.WriteFile(mapsV1, []byte("package a // my maps work\n"), 0o644)
	os.WriteFile(slicesFile, []byte("package a // my slices work\n"), 0o644)
	a1 := app.New(v1, ws)
	if _, err := a1.Store.Update(func(s *progress.State) error {
		s.Entry("maps_01").Status = progress.Completed
		s.CurrentExerciseID = "maps_01"
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	// A dry run changes nothing.
	a2 := app.New(v2, ws)
	res, err := a2.Upgrade(true)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Added) != 1 || len(res.Moved) != 1 || res.ProgressRenamed["maps_01"] != "maps_02" {
		t.Fatalf("dry-run plan = %+v", res)
	}
	if read(t, mapsV1) != "package a // my maps work\n" {
		t.Fatal("dry run touched the workspace")
	}

	res, err = a2.Upgrade(false)
	if err != nil {
		t.Fatalf("Upgrade: %v", err)
	}
	mapsV2 := filepath.Join(root, "exercises", "02_more", "maps_02", "a.go")
	if got := read(t, mapsV2); got != "package a // my maps work\n" {
		t.Errorf("renamed exercise has %q; the learner's work should have moved with it", got)
	}
	if _, err := os.Stat(filepath.Dir(mapsV1)); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("old directory %s still exists", filepath.Dir(mapsV1))
	}
	if got := read(t, slicesFile); got != "package a // my slices work\n" {
		t.Errorf("upgrade rewrote a learner file: %q", got)
	}
	if got := read(t, filepath.Join(root, "exercises", "01_basics", "slices_01", "README.md")); got != "new readme" {
		t.Errorf("README not refreshed: %q", got)
	}
	if _, err := os.Stat(filepath.Join(root, "exercises", "02_more", "structs_01", "a.go")); err != nil {
		t.Errorf("new exercise not created: %v", err)
	}
	if res.StateBackup == "" {
		t.Error("progress was changed without a backup")
	} else if _, err := os.Stat(res.StateBackup); err != nil {
		t.Errorf("backup missing: %v", err)
	}

	reopened, err := workspace.Open(root)
	if err != nil {
		t.Fatal(err)
	}
	if reopened.Meta.CurriculumVersion != "0.0.2" || reopened.Meta.UpgradedAt == nil {
		t.Errorf("workspace meta = %+v", reopened.Meta)
	}
	st, err := app.New(v2, reopened).Store.Load()
	if err != nil {
		t.Fatal(err)
	}
	if st.Get("maps_02").Status != progress.Completed || st.Exercises["maps_01"] != nil || st.CurrentExerciseID != "maps_02" {
		t.Errorf("progress not migrated: current %q, entries %v", st.CurrentExerciseID, st.Exercises)
	}

	// Upgrading again is a no-op, and going back to v1 is refused.
	res, err = app.New(v2, reopened).Upgrade(false)
	if err != nil || !res.Empty() || len(res.ProgressRenamed) != 0 {
		t.Errorf("second upgrade = %+v, %v; want nothing to do", res, err)
	}
	if _, err := app.New(v1, reopened).Upgrade(true); !errors.Is(err, workspace.ErrWorkspaceNewer) {
		t.Errorf("downgrade error = %v, want ErrWorkspaceNewer", err)
	}
}

func TestRecheckAfterChecksChange(t *testing.T) {
	cat := fakeCurriculum(t, "0.0.1", [][2]string{{"maps", "01_basics"}},
		[]fakeEx{{"01_basics", "maps_01", "maps", "r"}}, "")
	ws, err := workspace.Init(filepath.Join(t.TempDir(), "ws"), cat, workspace.InitOptions{})
	if err != nil {
		t.Fatal(err)
	}
	a := app.New(cat, ws)
	ex, _ := cat.Get("maps_01")
	files, _ := ws.ReadLearnerFiles(ex)
	st := &progress.State{Exercises: map[string]*progress.ExerciseState{
		"maps_01": {Status: progress.Completed, SourceFingerprint: workspace.Fingerprint(files), ChecksHash: ex.ChecksHash},
	}}
	if v := a.ViewOf(ex, st); v.Status != app.DisplayCompleted {
		t.Fatalf("status = %s, want completed", v.Status)
	}
	st.Exercises["maps_01"].ChecksHash = "from-an-older-curriculum"
	if v := a.ViewOf(ex, st); v.Status != app.DisplayRecheck || v.Status.Done() {
		t.Fatalf("status = %s, want recheck (not done)", v.Status)
	}
	if next := a.NextUnfinished(st); next == nil || next.ID != "maps_01" {
		t.Fatalf("NextUnfinished = %v, want maps_01 to be re-verified", next)
	}
}

func TestCompareVersions(t *testing.T) {
	for _, tt := range []struct {
		a, b string
		want int
	}{
		{"0.1.0", "0.1.0", 0}, {"0.9.0", "0.10.0", -1}, {"1.0", "0.99.9", 1}, {"v1.2.0", "1.2", 0}, {"1.0.0-rc1", "1.0.0-rc2", -1},
	} {
		if got := workspace.CompareVersions(tt.a, tt.b); got != tt.want {
			t.Errorf("CompareVersions(%q, %q) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
