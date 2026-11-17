// Command golings-dev is the maintainer tool for the curriculum:
//
//	golings-dev validate-curriculum   structural and coverage validation
//	golings-dev check [--all | id...]  contract checks: starter fails, solution passes, mutants fail
//	golings-dev coverage [--write]     render docs/curriculum-coverage.md
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"sync"

	"github.com/maskedsyntax/golings/curriculum"
	"github.com/maskedsyntax/golings/internal/catalog"
	"github.com/maskedsyntax/golings/internal/toolchain"
	"github.com/maskedsyntax/golings/internal/validate"
)

const coveragePath = "docs/curriculum-coverage.md"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	os.Exit(run(ctx, os.Args[1:], os.Stdout, os.Stderr))
}

func run(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		usage(stderr)
		return 2
	}
	cat, err := catalog.Load(curriculum.FS)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	switch args[0] {
	case "validate-curriculum":
		return validateCmd(cat, stdout, stderr)
	case "check":
		return checkCmd(ctx, cat, args[1:], stdout, stderr)
	case "coverage":
		return coverageCmd(cat, args[1:], stdout, stderr)
	case "help", "-h", "--help":
		usage(stdout)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown command %q\n", args[0])
		usage(stderr)
		return 2
	}
}

func usage(w io.Writer) {
	fmt.Fprintln(w, `usage:
  golings-dev validate-curriculum          validate metadata, assets and coverage
  golings-dev check [-j N] (--all | id...) starter fails, solution passes, mutants fail
  golings-dev coverage [--write]           print or write `+coveragePath)
}

func validateCmd(cat *catalog.Catalog, stdout, stderr io.Writer) int {
	rep := validate.Curriculum(cat)
	for _, w := range rep.Warnings {
		fmt.Fprintln(stdout, "warning:", w)
	}
	for _, e := range rep.Errors {
		fmt.Fprintln(stderr, "error:", e)
	}
	// The coverage document must be up to date.
	if data, err := os.ReadFile(coveragePath); err == nil && string(data) != validate.CoverageMarkdown(cat) {
		fmt.Fprintf(stderr, "error: %s is out of date; run `go run ./cmd/golings-dev coverage --write`\n", coveragePath)
		return 1
	}
	if !rep.OK() {
		return 1
	}
	fmt.Fprintf(stdout, "ok: %d exercises, %d topics, %d warnings\n", len(cat.Exercises()), len(cat.Topics.Topics), len(rep.Warnings))
	return 0
}

func checkCmd(ctx context.Context, cat *catalog.Catalog, args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("check", flag.ContinueOnError)
	fs.SetOutput(stderr)
	all := fs.Bool("all", false, "check every exercise")
	jobs := fs.Int("j", max(1, runtime.NumCPU()/2), "exercises to check in parallel")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	var exs []*catalog.Exercise
	switch {
	case *all:
		exs = cat.Exercises()
	case fs.NArg() == 0:
		fmt.Fprintln(stderr, "check: pass --all or exercise ids")
		return 2
	default:
		for _, q := range fs.Args() {
			ex, err := cat.Find(q)
			if err != nil {
				fmt.Fprintln(stderr, err)
				return 2
			}
			exs = append(exs, ex)
		}
	}
	tc, err := toolchain.Detect(ctx)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 4
	}

	results := make([]*validate.ContractResult, len(exs))
	sem := make(chan struct{}, max(1, *jobs))
	var wg sync.WaitGroup
	for i, ex := range exs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			results[i] = validate.Contract(ctx, cat, ex, tc)
		}()
	}
	wg.Wait()

	failed := 0
	for _, r := range results {
		if r.OK {
			names := make([]string, 0, len(r.Variants))
			for _, v := range r.Variants {
				names = append(names, v.Name)
			}
			fmt.Fprintf(stdout, "ok    %-18s %s\n", r.ExerciseID, strings.Join(names, ", "))
			continue
		}
		failed++
		fmt.Fprintf(stdout, "FAIL  %s\n", r.ExerciseID)
		for _, v := range r.Variants {
			if !v.OK {
				fmt.Fprintf(stdout, "      %s: %s\n", v.Name, strings.ReplaceAll(v.Problem, "\n", "\n        "))
			}
		}
	}
	fmt.Fprintf(stdout, "\n%d/%d exercises satisfy the contract\n", len(results)-failed, len(results))
	if failed > 0 || ctx.Err() != nil {
		return 1
	}
	return 0
}

func coverageCmd(cat *catalog.Catalog, args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("coverage", flag.ContinueOnError)
	fs.SetOutput(stderr)
	write := fs.Bool("write", false, "write "+coveragePath+" instead of printing")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	md := validate.CoverageMarkdown(cat)
	if !*write {
		fmt.Fprint(stdout, md)
		return 0
	}
	if err := os.WriteFile(coveragePath, []byte(md), 0o644); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	fmt.Fprintln(stdout, "wrote", coveragePath)
	return 0
}
