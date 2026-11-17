// Mutant: converting after integer division loses the fraction.
package values

import "strconv"

func Average(sum int, count int) float64 {
	if count == 0 {
		return 0
	}
	return float64(sum / count)
}

func Label(name string, age int) string {
	return name + " is " + strconv.Itoa(age)
}
