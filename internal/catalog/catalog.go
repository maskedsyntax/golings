// Package catalog loads and validates the curriculum metadata: the ordered
// manifest, the Go by Example topic inventory, and per-exercise metadata.
//
// manifest.json is the single source of truth for module membership and
// ordering; an exercise's ordinal is its position in the manifest.
// exercise.json files carry everything else and never repeat ordering.
package catalog

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"path"
	"regexp"
	"slices"
	"strings"
)

// ExercisesDir is the directory (inside the curriculum FS) that holds
// exercise assets, laid out as <module>/<id>/.
const ExercisesDir = "_exercises"

// ExerciseType describes what the learner is asked to do.
type ExerciseType string

const (
	TypeCompileFix  ExerciseType = "compile_fix"
	TypeImplement   ExerciseType = "implement"
	TypeDebug       ExerciseType = "debug"
	TypeOutput      ExerciseType = "output"
	TypeIntegration ExerciseType = "integration"
	TypeConcurrency ExerciseType = "concurrency"
)

// CheckKind is one entry of the fixed check allowlist. Metadata strings are
// never turned into commands; each kind maps to a hard-coded invocation.
type CheckKind string

const (
	CheckBuild  CheckKind = "build"
	CheckUnit   CheckKind = "unit"
	CheckOutput CheckKind = "output"
	CheckVet    CheckKind = "vet"
	CheckRace   CheckKind = "race"
)

// Starter expectations used by the curriculum contract checks.
const (
	StarterExpectCompileError = "compile_error"
	StarterExpectCheckFailure = "check_failure"
)

var (
	validTypes = []ExerciseType{TypeCompileFix, TypeImplement, TypeDebug, TypeOutput, TypeIntegration, TypeConcurrency}
	validKinds = []CheckKind{CheckBuild, CheckUnit, CheckOutput, CheckVet, CheckRace}
	validDiffs = []string{"beginner", "intermediate", "advanced"}
	validOS    = []string{"darwin", "linux", "windows"}

	idRE       = regexp.MustCompile(`^[a-z][a-z0-9_]*_[0-9]{2}$`)
	topicRE    = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)
	moduleRE   = regexp.MustCompile(`^[0-9]{2}_[a-z][a-z0-9_]*$`)
	fileNameRE = regexp.MustCompile(`^[A-Za-z0-9_][A-Za-z0-9_.-]*$`)
)

// TopicURLPrefix is the base URL of every Go by Example topic.
const TopicURLPrefix = "https://gobyexample.com/"

// Topic is one entry in the Go by Example topic inventory.
type Topic struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	URL           string `json:"url"`
	Module        string `json:"module"`
	SourceOrdinal int    `json:"source_ordinal"`
}

// Topics is the topics.json document.
type Topics struct {
	SchemaVersion int     `json:"schema_version"`
	Source        string  `json:"source"`
	SnapshotDate  string  `json:"snapshot_date"`
	Topics        []Topic `json:"topics"`
}

// Module groups exercises; its order defines presentation order.
type Module struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Exercises []string `json:"exercises"`
}

// Manifest is the manifest.json document.
type Manifest struct {
	SchemaVersion     int      `json:"schema_version"`
	CurriculumVersion string   `json:"curriculum_version"`
	ReleaseComplete   bool     `json:"release_complete"`
	Modules           []Module `json:"modules"`
}

// Exercise is the metadata of one exercise (exercise.json) plus fields
// derived from the manifest.
type Exercise struct {
	SchemaVersion    int          `json:"schema_version"`
	ID               string       `json:"id"`
	Topic            string       `json:"topic"`
	Title            string       `json:"title"`
	Type             ExerciseType `json:"type"`
	Difficulty       string       `json:"difficulty"`
	EstimatedMinutes int          `json:"estimated_minutes"`
	Prerequisites    []string     `json:"prerequisites"`
	StarterFiles     []string     `json:"starter_files"`
	FixtureFiles     []string     `json:"fixture_files"`
	Checks           []CheckKind  `json:"checks"`
	TimeoutSeconds   int          `json:"timeout_seconds"`
	Platforms        []string     `json:"platforms"`
	Hints            []string     `json:"hints"`
	StarterExpect    string       `json:"starter_expect"`

	// Derived from the manifest and topic inventory.
	Ordinal int    `json:"-"`
	Module  string `json:"-"`
	URL     string `json:"-"`
}

// SupportsOS reports whether the exercise can be evaluated on goos. An empty
// platform list means every supported platform.
func (e *Exercise) SupportsOS(goos string) bool {
	return len(e.Platforms) == 0 || slices.Contains(e.Platforms, goos)
}

// AssetDir is the exercise's directory inside the curriculum FS.
func (e *Exercise) AssetDir() string {
	return path.Join(ExercisesDir, e.Module, e.ID)
}

// Catalog is the validated, immutable curriculum metadata.
type Catalog struct {
	Manifest  Manifest
	Topics    Topics
	fsys      fs.FS
	exercises []*Exercise
	byID      map[string]*Exercise
	topicByID map[string]*Topic
}

// ValidationError collects every problem found while loading, so authors see
// all of them at once.
type ValidationError struct{ Problems []string }

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid curriculum (%d problems):\n  - %s", len(e.Problems), strings.Join(e.Problems, "\n  - "))
}

// Load reads manifest.json, topics.json and every exercise.json from fsys
// and validates them. On failure the error is a *ValidationError listing all
// problems.
func Load(fsys fs.FS) (*Catalog, error) {
	c := &Catalog{
		fsys:      fsys,
		byID:      map[string]*Exercise{},
		topicByID: map[string]*Topic{},
	}
	var problems []string
	addf := func(format string, args ...any) { problems = append(problems, fmt.Sprintf(format, args...)) }

	if err := decodeStrict(fsys, "manifest.json", &c.Manifest); err != nil {
		return nil, &ValidationError{Problems: []string{err.Error()}}
	}
	if err := decodeStrict(fsys, "topics.json", &c.Topics); err != nil {
		return nil, &ValidationError{Problems: []string{err.Error()}}
	}

	m := &c.Manifest
	if m.SchemaVersion != 1 {
		addf("manifest.json: unsupported schema_version %d", m.SchemaVersion)
	}
	if m.CurriculumVersion == "" {
		addf("manifest.json: curriculum_version is required")
	}
	if len(m.Modules) == 0 {
		addf("manifest.json: no modules")
	}
	modules := map[string]bool{}
	for _, mod := range m.Modules {
		if !moduleRE.MatchString(mod.ID) {
			addf("manifest.json: invalid module id %q (want NN_name)", mod.ID)
		}
		if modules[mod.ID] {
			addf("manifest.json: duplicate module %q", mod.ID)
		}
		modules[mod.ID] = true
		if mod.Title == "" {
			addf("manifest.json: module %q has no title", mod.ID)
		}
	}

	if c.Topics.SchemaVersion != 1 {
		addf("topics.json: unsupported schema_version %d", c.Topics.SchemaVersion)
	}
	urls := map[string]bool{}
	for i := range c.Topics.Topics {
		t := &c.Topics.Topics[i]
		if !topicRE.MatchString(t.ID) {
			addf("topics.json: invalid topic id %q", t.ID)
		}
		if c.topicByID[t.ID] != nil {
			addf("topics.json: duplicate topic id %q", t.ID)
		}
		c.topicByID[t.ID] = t
		if want := TopicURLPrefix + strings.ReplaceAll(t.ID, "_", "-"); t.URL != want {
			addf("topics.json: topic %q has url %q, want %q", t.ID, t.URL, want)
		}
		if urls[t.URL] {
			addf("topics.json: duplicate url %q", t.URL)
		}
		urls[t.URL] = true
		if t.Title == "" {
			addf("topics.json: topic %q has no title", t.ID)
		}
		if !modules[t.Module] {
			addf("topics.json: topic %q references unknown module %q", t.ID, t.Module)
		}
		if t.SourceOrdinal != i+1 {
			addf("topics.json: topic %q has source_ordinal %d, want %d (keep source order)", t.ID, t.SourceOrdinal, i+1)
		}
	}

	ordinal := 0
	for _, mod := range m.Modules {
		for _, id := range mod.Exercises {
			ordinal++
			if !idRE.MatchString(id) {
				addf("manifest.json: invalid exercise id %q (want topic_NN)", id)
				continue
			}
			if c.byID[id] != nil {
				addf("manifest.json: duplicate exercise id %q", id)
				continue
			}
			ex := &Exercise{}
			file := path.Join(ExercisesDir, mod.ID, id, "exercise.json")
			if err := decodeStrict(fsys, file, ex); err != nil {
				addf("%v", err)
				continue
			}
			ex.Ordinal, ex.Module = ordinal, mod.ID
			for _, p := range validateExercise(ex, c.topicByID) {
				addf("%s: %s", file, p)
			}
			if t := c.topicByID[ex.Topic]; t != nil {
				ex.URL = t.URL
			}
			c.exercises = append(c.exercises, ex)
			c.byID[id] = ex
		}
	}

	// Prerequisites must name an earlier exercise.
	for _, ex := range c.exercises {
		for _, p := range ex.Prerequisites {
			pre := c.byID[p]
			switch {
			case pre == nil:
				addf("exercise %s: unknown prerequisite %q", ex.ID, p)
			case pre.Ordinal >= ex.Ordinal:
				addf("exercise %s: prerequisite %q must come earlier in the manifest", ex.ID, p)
			}
		}
	}

	if len(problems) > 0 {
		return nil, &ValidationError{Problems: problems}
	}
	return c, nil
}

func validateExercise(ex *Exercise, topics map[string]*Topic) []string {
	var p []string
	addf := func(format string, args ...any) { p = append(p, fmt.Sprintf(format, args...)) }

	if ex.SchemaVersion != 1 {
		addf("unsupported schema_version %d", ex.SchemaVersion)
	}
	if t := topics[ex.Topic]; t == nil {
		addf("unknown topic %q", ex.Topic)
	} else {
		if t.Module != ex.Module {
			addf("topic %q belongs to module %q but the exercise is listed under %q", ex.Topic, t.Module, ex.Module)
		}
		if !strings.HasPrefix(ex.ID, ex.Topic+"_") {
			addf("id %q must start with its topic %q", ex.ID, ex.Topic)
		}
	}
	if strings.TrimSpace(ex.Title) == "" {
		addf("title is required")
	}
	if !slices.Contains(validTypes, ex.Type) {
		addf("invalid type %q", ex.Type)
	}
	if !slices.Contains(validDiffs, ex.Difficulty) {
		addf("invalid difficulty %q", ex.Difficulty)
	}
	if ex.EstimatedMinutes <= 0 {
		addf("estimated_minutes must be positive")
	}
	if ex.TimeoutSeconds <= 0 || ex.TimeoutSeconds > 600 {
		addf("timeout_seconds must be between 1 and 600")
	}
	if len(ex.StarterFiles) == 0 {
		addf("starter_files must not be empty")
	}
	seen := map[string]bool{}
	for _, f := range ex.StarterFiles {
		if !fileNameRE.MatchString(f) || !strings.HasSuffix(f, ".go") || strings.HasSuffix(f, "_test.go") {
			addf("invalid starter file %q (want a plain non-test .go file name)", f)
		}
		if seen[f] {
			addf("duplicate starter file %q", f)
		}
		seen[f] = true
	}
	for _, f := range ex.FixtureFiles {
		if !fs.ValidPath(f) || f == "." {
			addf("invalid fixture path %q", f)
		}
	}
	if len(ex.Checks) == 0 {
		addf("checks must not be empty")
	}
	kinds := map[CheckKind]bool{}
	for _, k := range ex.Checks {
		if !slices.Contains(validKinds, k) {
			addf("unknown check kind %q", k)
		}
		if kinds[k] {
			addf("duplicate check %q", k)
		}
		kinds[k] = true
	}
	if ex.Type == TypeOutput && !kinds[CheckOutput] {
		addf("output exercises must declare the %q check", CheckOutput)
	}
	for _, goos := range ex.Platforms {
		if !slices.Contains(validOS, goos) {
			addf("unknown platform %q", goos)
		}
	}
	if len(ex.Hints) == 0 || len(ex.Hints) > 3 {
		addf("want 1 to 3 hints, got %d", len(ex.Hints))
	}
	for i, h := range ex.Hints {
		if strings.TrimSpace(h) == "" {
			addf("hint %d is empty", i+1)
		}
	}
	switch ex.StarterExpect {
	case "", StarterExpectCompileError, StarterExpectCheckFailure:
	default:
		addf("invalid starter_expect %q", ex.StarterExpect)
	}
	for _, pre := range ex.Prerequisites {
		if pre == ex.ID {
			addf("exercise lists itself as a prerequisite")
		}
	}
	return p
}

func decodeStrict(fsys fs.FS, name string, v any) error {
	data, err := fs.ReadFile(fsys, name)
	if err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(v); err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	if _, err := dec.Token(); !errors.Is(err, io.EOF) {
		return fmt.Errorf("%s: trailing data after JSON document", name)
	}
	return nil
}

// FS returns the underlying curriculum file system.
func (c *Catalog) FS() fs.FS { return c.fsys }

// Version returns the curriculum version.
func (c *Catalog) Version() string { return c.Manifest.CurriculumVersion }

// Exercises returns all exercises in manifest order. Callers must not modify
// the returned slice or its elements.
func (c *Catalog) Exercises() []*Exercise { return c.exercises }

// Get returns the exercise with the exact id.
func (c *Catalog) Get(id string) (*Exercise, bool) {
	ex, ok := c.byID[id]
	return ex, ok
}

// Topic returns the topic with the given id, or nil.
func (c *Catalog) Topic(id string) *Topic { return c.topicByID[id] }

// ModuleTitle returns the display title of a module id.
func (c *Catalog) ModuleTitle(id string) string {
	for _, m := range c.Manifest.Modules {
		if m.ID == id {
			return m.Title
		}
	}
	return id
}

// Assets returns the exercise's asset directory as a file system containing
// README.md, starter/, checks/, solution/, fixtures/ and mutants/.
func (c *Catalog) Assets(ex *Exercise) (fs.FS, error) {
	return fs.Sub(c.fsys, ex.AssetDir())
}
