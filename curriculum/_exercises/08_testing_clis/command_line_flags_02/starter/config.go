// command_line_flags_02: Custom flag types
//
// Make TagList a flag.Value and implement Parse. See README.md.
//
// Reference: https://gobyexample.com/command-line-flags
package config

import (
	"errors"
	"time"
)

// TagList collects every -tag value.
type TagList []string

// TODO: String and Set methods.

// Config is the parsed command line.
type Config struct {
	Tags    TagList
	Timeout time.Duration
	Verbose bool
	Files   []string
}

// Parse parses args (without the program name).
func Parse(args []string) (Config, error) {
	// TODO
	return Config{}, errors.New("not implemented")
}
