// switch_01: No fallthrough
//
// IsWeekend("Saturday") wrongly returns false. Why?
//
// Reference: https://gobyexample.com/switch
package weekend

// IsWeekend reports whether day is "Saturday" or "Sunday".
func IsWeekend(day string) bool {
	switch day {
	case "Saturday":
	case "Sunday":
		return true
	}
	return false
}
