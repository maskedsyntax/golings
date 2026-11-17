// Package version reports the golings CLI version.
package version

import "runtime/debug"

// Version is set at release time with
// -ldflags "-X github.com/maskedsyntax/golings/internal/version.Version=v1.2.3".
var Version = ""

// String returns the CLI version, falling back to module build info (which
// is populated by `go install module@version`) and finally to "dev".
func String() string {
	if Version != "" {
		return Version
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		if v := info.Main.Version; v != "" && v != "(devel)" {
			return v
		}
	}
	return "dev"
}
