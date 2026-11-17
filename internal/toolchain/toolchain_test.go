package toolchain

import "testing"

func TestParseVersion(t *testing.T) {
	tests := []struct {
		in           string
		major, minor int
		devel        bool
	}{
		{"go1.24.3", 1, 24, false},
		{"go1.24", 1, 24, false},
		{"go1.25rc1", 1, 25, false},
		{"go1.27.3 X:nocoverageredesign", 1, 27, false},
		{"devel go1.26-abcdef", 0, 0, true},
		{"", 0, 0, true},
	}
	for _, tt := range tests {
		major, minor, devel := ParseVersion(tt.in)
		if major != tt.major || minor != tt.minor || devel != tt.devel {
			t.Errorf("ParseVersion(%q) = %d, %d, %v; want %d, %d, %v", tt.in, major, minor, devel, tt.major, tt.minor, tt.devel)
		}
	}
}

func TestAtLeast(t *testing.T) {
	if !AtLeast(1, 24, 1, 24) || !AtLeast(1, 30, 1, 24) || !AtLeast(2, 0, 1, 24) {
		t.Error("AtLeast rejected a newer or equal version")
	}
	if AtLeast(1, 23, 1, 24) {
		t.Error("AtLeast accepted go1.23")
	}
}

func TestRaceSupported(t *testing.T) {
	if !RaceSupported("darwin", "arm64", false) {
		t.Error("macOS should support -race without cgo")
	}
	if RaceSupported("linux", "amd64", false) {
		t.Error("linux -race needs cgo")
	}
	if !RaceSupported("linux", "amd64", true) {
		t.Error("linux/amd64 with cgo supports -race")
	}
	if RaceSupported("linux", "386", true) {
		t.Error("linux/386 does not support -race")
	}
}
