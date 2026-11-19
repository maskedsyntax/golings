// Mutant: each -tag replaces the previous one.
package config

import (
	"errors"
	"flag"
	"io"
	"strings"
	"time"
)

type TagList []string

func (t *TagList) String() string { return strings.Join(*t, ",") }

func (t *TagList) Set(v string) error {
	if v == "" {
		return errors.New("empty tag")
	}
	*t = TagList{v}
	return nil
}

type Config struct {
	Tags    TagList
	Timeout time.Duration
	Verbose bool
	Files   []string
}

func Parse(args []string) (Config, error) {
	var cfg Config
	fs := flag.NewFlagSet("tool", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.Var(&cfg.Tags, "tag", "")
	fs.DurationVar(&cfg.Timeout, "timeout", 5*time.Second, "")
	fs.BoolVar(&cfg.Verbose, "v", false, "")
	if err := fs.Parse(args); err != nil {
		return Config{}, err
	}
	cfg.Files = fs.Args()
	return cfg, nil
}
