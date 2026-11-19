package intmin

// IntMin returns the smaller of a and b.
//
// It is declared as a variable so golings can swap in broken versions and
// check that your tests notice. Call it like any other function.
var IntMin = func(a, b int) int {
	if a < b {
		return a
	}
	return b
}
