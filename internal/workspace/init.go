package workspace

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/maskedsyntax/golings/internal/catalog"
)

// DefaultDir is the directory `golings init` creates when no path is given.
const DefaultDir = "golings-workspace"

// workspaceGoMod lets editors and gopls understand learner files. Canonical
// tests are never copied here, so this module only ever holds starters.
const workspaceGoMod = "module golings.workspace\n\ngo 1.24\n"

// InitOptions controls workspace creation.
type InitOptions struct {
	// AllowNonEmpty permits initializing into a directory that already has
	// files, as long as none of the paths golings writes already exist.
	AllowNonEmpty bool
	CLIVersion    string
}

// ErrNotEmpty means the destination already contains files.
var ErrNotEmpty = errors.New("destination directory is not empty")

// ErrAlreadyWorkspace means the destination already is a workspace.
var ErrAlreadyWorkspace = errors.New("destination is already a golings workspace")

// Init materializes a new workspace at dest from the catalog's embedded
// starters. It never overwrites an existing file.
func Init(dest string, cat *catalog.Catalog, opts InitOptions) (_ *Workspace, err error) {
	abs, err := filepath.Abs(dest)
	if err != nil {
		return nil, err
	}
	switch fi, statErr := os.Stat(abs); {
	case errors.Is(statErr, fs.ErrNotExist):
		// We create the directory, so remove it again if anything fails
		// and avoid leaving a half-written workspace behind.
		defer func() {
			if err != nil {
				os.RemoveAll(abs)
			}
		}()
	case statErr != nil:
		return nil, statErr
	case !fi.IsDir():
		return nil, fmt.Errorf("%s exists and is not a directory", abs)
	default:
		if _, err := os.Stat(filepath.Join(abs, MetaDir, metaFile)); err == nil {
			return nil, fmt.Errorf("%s: %w", abs, ErrAlreadyWorkspace)
		}
		entries, err := os.ReadDir(abs)
		if err != nil {
			return nil, err
		}
		if len(entries) > 0 {
			if !opts.AllowNonEmpty {
				return nil, fmt.Errorf("%s: %w", abs, ErrNotEmpty)
			}
			for _, p := range []string{MetaDir, exercisesDir, "go.mod", "README.md"} {
				if _, err := os.Lstat(filepath.Join(abs, p)); err == nil {
					return nil, fmt.Errorf("%s already contains %s; refusing to mix it with a workspace", abs, p)
				}
			}
		}
	}

	if err := os.MkdirAll(filepath.Join(abs, MetaDir), 0o755); err != nil {
		return nil, err
	}
	meta := Meta{
		WorkspaceFormatVersion: FormatVersion,
		CurriculumVersion:      cat.Version(),
		CreatedAt:              time.Now().UTC().Truncate(time.Second),
		CreatedBy:              "golings " + opts.CLIVersion,
	}
	metaJSON, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return nil, err
	}
	if err := createFile(filepath.Join(abs, "go.mod"), []byte(workspaceGoMod)); err != nil {
		return nil, err
	}
	if err := createFile(filepath.Join(abs, "README.md"), []byte(workspaceReadme(cat))); err != nil {
		return nil, err
	}

	w := &Workspace{Root: abs, Meta: meta}
	for _, ex := range cat.Exercises() {
		if err := w.materialize(cat, ex); err != nil {
			return nil, fmt.Errorf("writing %s: %w", ex.ID, err)
		}
	}
	// Written last: a workspace only "exists" once fully materialized.
	if err := createFile(filepath.Join(abs, MetaDir, metaFile), append(metaJSON, '\n')); err != nil {
		return nil, err
	}
	return w, nil
}

// materialize copies an exercise's starter files and README into the
// workspace.
func (w *Workspace) materialize(cat *catalog.Catalog, ex *catalog.Exercise) error {
	assets, err := cat.Assets(ex)
	if err != nil {
		return err
	}
	dir := w.ExerciseDir(ex)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	for _, name := range ex.StarterFiles {
		data, err := fs.ReadFile(assets, path.Join("starter", name))
		if err != nil {
			return err
		}
		if err := createFile(filepath.Join(dir, name), data); err != nil {
			return err
		}
	}
	if data, err := fs.ReadFile(assets, "README.md"); err == nil {
		if err := createFile(filepath.Join(dir, "README.md"), data); err != nil {
			return err
		}
	}
	return nil
}

// createFile writes a new file and fails if it already exists.
func createFile(name string, data []byte) error {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

func workspaceReadme(cat *catalog.Catalog) string {
	var b strings.Builder
	b.WriteString("# Golings workspace\n\n")
	b.WriteString("Learn Go by fixing small Go programs. Each folder under `exercises/` holds one\n")
	b.WriteString("exercise: open its `.go` file in any editor, read the instructions at the top or in\n")
	b.WriteString("its `README.md`, and make it pass.\n\n")
	b.WriteString("```sh\n")
	b.WriteString("golings next          # show the next unfinished exercise\n")
	b.WriteString("golings verify        # check the current exercise\n")
	b.WriteString("golings list          # see every exercise and its status\n")
	b.WriteString("golings progress      # completion summary\n")
	b.WriteString("```\n\n")
	fmt.Fprintf(&b, "Curriculum version %s, %d exercises. Progress is stored in `.golings/`.\n\n", cat.Version(), len(cat.Exercises()))
	b.WriteString("Exercises follow the topics of [Go by Example](https://gobyexample.com/). The learning model\n")
	b.WriteString("comes from [Rustlings](https://github.com/rust-lang/rustlings) and [Ziglings](https://ziglings.org/).\n")
	return b.String()
}
