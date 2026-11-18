// Mutant: counts weekends too.
package bizdays

import "time"

func AddBusinessDays(t time.Time, n int) time.Time {
	return t.AddDate(0, 0, n)
}
