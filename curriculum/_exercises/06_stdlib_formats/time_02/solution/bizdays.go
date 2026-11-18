// time_02: Business days (reference solution)
package bizdays

import "time"

// AddBusinessDays returns t moved forward by n weekdays (n >= 0).
func AddBusinessDays(t time.Time, n int) time.Time {
	for n > 0 {
		t = t.AddDate(0, 0, 1)
		if wd := t.Weekday(); wd != time.Saturday && wd != time.Sunday {
			n--
		}
	}
	return t
}
