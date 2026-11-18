// if_else_01: Letter grades (reference solution)
package grade

// Grade converts a 0-100 score to a letter grade, or "invalid".
func Grade(score int) string {
	if score < 0 || score > 100 {
		return "invalid"
	} else if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	}
	return "F"
}
