//go:build windows

package render

import (
	"os"

	"golang.org/x/sys/windows"
)

// EnableVT turns on ANSI escape processing for a Windows console. It reports
// whether escapes can be used.
func EnableVT(f *os.File) bool {
	h := windows.Handle(f.Fd())
	var mode uint32
	if err := windows.GetConsoleMode(h, &mode); err != nil {
		return false
	}
	if mode&windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING != 0 {
		return true
	}
	return windows.SetConsoleMode(h, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING) == nil
}
