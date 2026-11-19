// command_line_flags_02: Custom flag types (reference solution)
package config

import (
	"errors"
	"flag"
	"io"
	"strings"
	"time"
)

// TagList collects every -tag value.
type TagList []string

// String implements flag.Value.
func (t *TagList) String() string { return strings.Join(*t, ",") }

// Set implements flag.Value; it is called once per -tag.
func (t *TagList) Set(v string) error {
	if v == "" {
		return errors.New("empty tag")
	}
	*t = append(*t, v)
	return nil
}

// Config is the parsed command line.
type Config struct {
	Tags    TagList
	Timeout time.Duration
	Verbose bool
	Files   []string
}

// Parse parses args (without the program name).
func Parse(args []string) (Config, error) {
	var cfg Config
	fs := flag.NewFlagSet("tool", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.Var(&cfg.Tags, "tag", "a tag (repeatable)")
	fs.DurationVar(&cfg.Timeout, "timeout", 5*time.Second, "how long to wait")
	fs.BoolVar(&cfg.Verbose, "v", false, "verbose output")
	if err := fs.Parse(args); err != nil {
		return Config{}, err
	}
	cfg.Files = fs.Args()
	return cfg, nil
}
