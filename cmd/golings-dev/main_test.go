package main

import (
	"strings"
	"testing"

	"github.com/maskedsyntax/golings/curriculum"
	"github.com/maskedsyntax/golings/internal/catalog"
)

const samplePage = `<html><body><ul>
  <li><a href="hello-world">Hello World</a></li>
  <li><a href="values">Values</a></li>
  <li><a href="time-formatting-parsing">Time Formatting / Parsing</a></li>
  <li><a href="brand-new">Brand &amp; New</a></li>
</ul><p><a href="https://example.com">not a topic</a></p></body></html>`

func TestParseIndex(t *testing.T) {
	got := parseIndex(samplePage)
	want := []indexEntry{{"hello-world", "Hello World"}, {"values", "Values"},
		{"time-formatting-parsing", "Time Formatting / Parsing"}, {"brand-new", "Brand & New"}}
	if len(got) != len(want) {
		t.Fatalf("parseIndex = %v", got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("entry %d = %v, want %v", i, got[i], want[i])
		}
	}
}

func TestDiffTopics(t *testing.T) {
	ours := []catalog.Topic{
		{Title: "Hello World", URL: catalog.TopicURLPrefix + "hello-world"},
		{Title: "Values", URL: catalog.TopicURLPrefix + "values"},
		{Title: "Variables", URL: catalog.TopicURLPrefix + "variables"},
	}
	live := []indexEntry{{"values", "Values"}, {"hello-world", "Hello, World"}, {"constants", "Constants"}}
	got := strings.Join(diffTopics(ours, live), "\n")
	for _, want := range []string{
		`title changed: hello-world: "Hello World" -> "Hello, World"`,
		`removed upstream: variables`,
		`new upstream: constants`,
		`order changed: values now comes before hello-world`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("diff is missing %q:\n%s", want, got)
		}
	}
	if d := diffTopics(ours[:2], []indexEntry{{"hello-world", "Hello World"}, {"values", "Values"}}); len(d) != 0 {
		t.Errorf("identical lists gave %v", d)
	}
}

func TestNextExerciseID(t *testing.T) {
	cat, err := catalog.Load(curriculum.FS)
	if err != nil {
		t.Fatal(err)
	}
	id, last := nextExerciseID(cat, "maps")
	if id != "maps_05" || last != "maps_04" {
		t.Errorf("nextExerciseID(maps) = %s, %s; want maps_05, maps_04", id, last)
	}
}

func TestInsertIntoManifest(t *testing.T) {
	m := `{"modules": [
    {
      "id": "01_x",
      "title": "X",
      "exercises": [
        "a_01", "a_02",
        "b_01"
      ]
    },
    {"id": "02_y", "title": "Y", "exercises": []}
  ]}`
	got, err := insertIntoManifest(m, "01_x", "a_02", "a_03")
	if err != nil || !strings.Contains(got, `"a_01", "a_02", "a_03",`) {
		t.Errorf("after a_02:\n%s (%v)", got, err)
	}
	got, err = insertIntoManifest(m, "02_y", "", "c_01")
	if err != nil || !strings.Contains(got, `"exercises": ["c_01"]`) {
		t.Errorf("into empty module:\n%s (%v)", got, err)
	}
	got, err = insertIntoManifest(m, "01_x", "", "z_01")
	if err != nil || !strings.Contains(got, "\"exercises\": [\n        \"z_01\",\n        \"a_01\"") {
		t.Errorf("first in module:\n%s (%v)", got, err)
	}
	if _, err := insertIntoManifest(m, "09_none", "", "q_01"); err == nil {
		t.Error("unknown module accepted")
	}
}

func TestShards(t *testing.T) {
	cat, err := catalog.Load(curriculum.FS)
	if err != nil {
		t.Fatal(err)
	}
	all := cat.Exercises()
	seen := map[string]int{}
	for i := 1; i <= 3; i++ {
		part := shardOf(all, i, 3)
		if len(part) < len(all)/3 || len(part) > len(all)/3+1 {
			t.Errorf("shard %d/3 has %d of %d exercises", i, len(part), len(all))
		}
		for _, ex := range part {
			seen[ex.ID]++
		}
	}
	if len(seen) != len(all) {
		t.Errorf("shards cover %d of %d exercises", len(seen), len(all))
	}
	for id, n := range seen {
		if n != 1 {
			t.Errorf("%s is in %d shards", id, n)
		}
	}
	for _, bad := range []string{"0/3", "4/3", "1", "a/b", "1/0"} {
		if _, _, err := parseShard(bad); err == nil {
			t.Errorf("parseShard(%q) accepted", bad)
		}
	}
}
