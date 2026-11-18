// switch_01: No fallthrough (reference solution)
package weekend

// IsWeekend reports whether day is "Saturday" or "Sunday".
func IsWeekend(day string) bool {
	switch day {
	case "Saturday", "Sunday":
		return true
	}
	return false
}
