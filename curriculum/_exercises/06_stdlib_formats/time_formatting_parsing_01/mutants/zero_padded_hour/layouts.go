// Mutant: 12-hour clock with a leading zero.
package layouts

import "time"

func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func ParseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

func Clock(t time.Time) string {
	return t.Format("03:04 PM")
}
