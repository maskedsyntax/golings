// string_formatting_02: Human-readable sizes (reference solution)
package size

import "fmt"

// FormatBytes formats n bytes using B, KiB, MiB, GiB or TiB.
func FormatBytes(n int64) string {
	if n < 1024 {
		return fmt.Sprintf("%d B", n)
	}
	units := []string{"KiB", "MiB", "GiB", "TiB"}
	v := float64(n)
	for i, u := range units {
		v /= 1024
		if v < 1024 || i == len(units)-1 {
			return fmt.Sprintf("%.1f %s", v, u)
		}
	}
	panic("unreachable")
}
