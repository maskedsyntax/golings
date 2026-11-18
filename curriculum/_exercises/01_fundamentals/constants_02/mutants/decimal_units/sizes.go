// Mutant: decimal (1000-based) units instead of 1024-based ones.
package sizes

import "fmt"

const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
)

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
