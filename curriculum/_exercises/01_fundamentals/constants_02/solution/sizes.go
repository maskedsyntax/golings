// constants_02: Human-readable sizes with iota (reference solution)
package sizes

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

// HumanSize formats a byte count using the largest fitting unit.
func HumanSize(n int64) string {
	switch {
	case n >= GB:
		return fmt.Sprintf("%.1f GB", float64(n)/GB)
	case n >= MB:
		return fmt.Sprintf("%.1f MB", float64(n)/MB)
	case n >= KB:
		return fmt.Sprintf("%.1f KB", float64(n)/KB)
	default:
		return fmt.Sprintf("%d B", n)
	}
}
