// interfaces_02: Type switches (reference solution)
package describe

import "fmt"

// Describe returns a short description of the dynamic type and value of x.
func Describe(x any) string {
	switch v := x.(type) {
	case nil:
		return "nil"
	case int:
		return fmt.Sprintf("int %d", v)
	case string:
		return fmt.Sprintf("string %q", v)
	case []int:
		return fmt.Sprintf("%d ints", len(v))
	case fmt.Stringer:
		return "stringer " + v.String()
	default:
		return "unknown"
	}
}
