// switch_03: Type switches (reference solution)
package describe

import "fmt"

// Describe names the dynamic type of v and summarizes its value.
func Describe(v any) string {
	switch x := v.(type) {
	case nil:
		return "nil"
	case int:
		return fmt.Sprintf("int %d", x)
	case string:
		return fmt.Sprintf("string %q", x)
	case bool:
		return fmt.Sprintf("bool %t", x)
	case []int:
		return fmt.Sprintf("[]int of length %d", len(x))
	default:
		return "unknown"
	}
}
