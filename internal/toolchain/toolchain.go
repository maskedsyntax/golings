// Package toolchain locates the local Go toolchain and reports its
// capabilities.
package toolchain

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// MinMajor and MinMinor are the oldest Go release golings supports.
const (
	MinMajor = 1
	MinMinor = 24
)

// Info describes a detected Go toolchain.
type Info struct {
	GoBin         string `json:"go_bin"`
	Version       string `json:"version"` // e.g. "go1.24.3"
	Major         int    `json:"major"`
	Minor         int    `json:"minor"`
	Devel         bool   `json:"devel"`
	GOOS          string `json:"goos"`
	GOARCH        string `json:"goarch"`
	GOROOT        string `json:"goroot"`
	GOPATH        string `json:"gopath"`
	GOCACHE       string `json:"gocache"`
	GOMODCACHE    string `json:"gomodcache"`
	CGOEnabled    bool   `json:"cgo_enabled"`
	RaceSupported bool   `json:"race_supported"`
}

// ErrNotFound means no `go` executable was found on PATH.
var ErrNotFound = errors.New("the Go toolchain (`go`) was not found on PATH; install Go from https://go.dev/dl/")

// TooOldError reports a toolchain older than the supported minimum.
type TooOldError struct{ Version string }

func (e *TooOldError) Error() string {
	return fmt.Sprintf("Go %s is too old; golings needs Go %d.%d or newer (https://go.dev/dl/)", strings.TrimPrefix(e.Version, "go"), MinMajor, MinMinor)
}

// Detect finds `go` on PATH and queries it. It returns ErrNotFound,
// *TooOldError, or a wrapped error if the toolchain cannot be queried.
func Detect(ctx context.Context) (*Info, error) {
	bin, err := exec.LookPath("go")
	if err != nil {
		return nil, ErrNotFound
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, bin, "env", "-json",
		"GOVERSION", "GOOS", "GOARCH", "GOROOT", "GOPATH", "GOCACHE", "GOMODCACHE", "CGO_ENABLED")
	// Run outside any module so a go.mod in the current directory cannot
	// trigger a toolchain switch, and never download a toolchain.
	cmd.Dir = os.TempDir()
	cmd.Env = append(os.Environ(), "GOTOOLCHAIN=local", "GOFLAGS=")
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("running `go env`: %v: %s", err, strings.TrimSpace(stderr.String()))
	}
	var env map[string]string
	if err := json.Unmarshal(stdout.Bytes(), &env); err != nil {
		return nil, fmt.Errorf("parsing `go env -json` output: %w", err)
	}

	info := &Info{
		GoBin:      bin,
		Version:    env["GOVERSION"],
		GOOS:       env["GOOS"],
		GOARCH:     env["GOARCH"],
		GOROOT:     env["GOROOT"],
		GOPATH:     env["GOPATH"],
		GOCACHE:    env["GOCACHE"],
		GOMODCACHE: env["GOMODCACHE"],
		CGOEnabled: env["CGO_ENABLED"] == "1",
	}
	info.Major, info.Minor, info.Devel = ParseVersion(info.Version)
	info.RaceSupported = RaceSupported(info.GOOS, info.GOARCH, info.CGOEnabled)
	if !info.Devel && !AtLeast(info.Major, info.Minor, MinMajor, MinMinor) {
		return info, &TooOldError{Version: info.Version}
	}
	return info, nil
}

// ParseVersion parses GOVERSION strings such as "go1.24.3", "go1.25rc1" or
// "devel go1.26-abcdef". Unparseable or development versions report devel.
func ParseVersion(v string) (major, minor int, devel bool) {
	if strings.HasPrefix(v, "devel") {
		return 0, 0, true
	}
	v = strings.TrimPrefix(v, "go")
	// Cut a trailing "rc1"/"beta1" or patch component, and any build suffix.
	if i := strings.IndexAny(v, " -+"); i >= 0 {
		v = v[:i]
	}
	parts := strings.SplitN(v, ".", 3)
	if len(parts) < 2 {
		return 0, 0, true
	}
	mi := parts[1]
	if i := strings.IndexFunc(mi, func(r rune) bool { return r < '0' || r > '9' }); i >= 0 {
		mi = mi[:i]
	}
	var err1, err2 error
	major, err1 = strconv.Atoi(parts[0])
	minor, err2 = strconv.Atoi(mi)
	if err1 != nil || err2 != nil {
		return 0, 0, true
	}
	return major, minor, false
}

// AtLeast reports whether major.minor >= wantMajor.wantMinor.
func AtLeast(major, minor, wantMajor, wantMinor int) bool {
	return major > wantMajor || (major == wantMajor && minor >= wantMinor)
}

// raceTargets lists GOOS/GOARCH pairs where `go test -race` is available.
var raceTargets = map[string]bool{
	"linux/amd64": true, "linux/arm64": true, "linux/ppc64le": true, "linux/s390x": true,
	"darwin/amd64": true, "darwin/arm64": true,
	"freebsd/amd64": true, "netbsd/amd64": true,
	"windows/amd64": true,
}

// RaceSupported reports whether the race detector is usable. It requires
// cgo everywhere except macOS.
func RaceSupported(goos, goarch string, cgo bool) bool {
	if !raceTargets[goos+"/"+goarch] {
		return false
	}
	return cgo || goos == "darwin"
}
