package runner

import (
	"os"
	"strings"

	"github.com/maskedsyntax/golings/internal/toolchain"
)

// passthroughEnv lists the only variables inherited from the user's
// environment. Everything else is set explicitly so evaluation is
// reproducible and never touches the network.
var passthroughEnv = []string{
	"PATH", "HOME", "USER", "LOGNAME",
	"TMPDIR", "TMP", "TEMP",
	"XDG_CACHE_HOME", "XDG_CONFIG_HOME",
	"CC", "CXX", // the race detector needs a C toolchain on Linux
	// Windows essentials.
	"USERPROFILE", "LOCALAPPDATA", "APPDATA", "SystemRoot", "SYSTEMROOT",
	"SystemDrive", "WINDIR", "ComSpec", "PATHEXT", "ProgramData",
	"ProgramFiles", "ProgramFiles(x86)", "NUMBER_OF_PROCESSORS", "PROCESSOR_ARCHITECTURE",
}

// buildEnv returns the environment for Go tool and exercise processes.
func buildEnv(tc *toolchain.Info) []string {
	var env []string
	for _, k := range passthroughEnv {
		if v, ok := os.LookupEnv(k); ok {
			env = append(env, k+"="+v)
		}
	}
	cgo := "0"
	if tc.CGOEnabled {
		cgo = "1"
	}
	env = append(env,
		"GOROOT="+tc.GOROOT,
		"GOPATH="+tc.GOPATH,
		"GOCACHE="+tc.GOCACHE, // reuse the user's build cache for speed
		"GOMODCACHE="+tc.GOMODCACHE,
		"CGO_ENABLED="+cgo,
		"GOENV=off",         // ignore the user's go env file
		"GOFLAGS=",          // ignore user-wide flags such as -mod=vendor
		"GOPROXY=off",       // stdlib-only: never download modules
		"GOWORK=off",        // never pick up a go.work
		"GOTOOLCHAIN=local", // never download a toolchain
		"GO111MODULE=on",
		"GOTELEMETRY=off",
		"GOLINGS=1", // lets canonical tests detect they run under golings
	)
	return dedupeEnv(env)
}

// dedupeEnv keeps the last value for each key (case-insensitively on
// Windows, where SystemRoot and SYSTEMROOT are the same variable).
func dedupeEnv(env []string) []string {
	seen := map[string]int{}
	var out []string
	for _, kv := range env {
		k, _, _ := strings.Cut(kv, "=")
		key := k
		if isWindows {
			key = strings.ToUpper(k)
		}
		if i, ok := seen[key]; ok {
			out[i] = kv
			continue
		}
		seen[key] = len(out)
		out = append(out, kv)
	}
	return out
}
