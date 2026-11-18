//go:build !windows

package render

import "os"

// EnableVT reports whether ANSI escapes can be used; terminals outside
// Windows support them natively.
func EnableVT(*os.File) bool { return true }
