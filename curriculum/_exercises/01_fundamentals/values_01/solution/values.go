// values_01: Mixing numbers and strings (reference solution)
package values

import "strconv"

// Average returns sum/count as an exact (non-truncated) value.
// A count of zero yields 0.
func Average(sum int, count int) float64 {
	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}

// Label describes a gopher's age, for example "Gopher is 13".
func Label(name string, age int) string {
	return name + " is " + strconv.Itoa(age)
}
