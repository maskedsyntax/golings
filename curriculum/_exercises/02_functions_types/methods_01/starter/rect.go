// methods_01: Receivers
//
// The tests call these as methods: r.Area() and r.Scale(2). Make them methods.
//
//	r := Rect{Width: 2, Height: 3}
//	r.Area() == 6
//	r.Scale(2) // r == Rect{4, 6}
//
// Reference: https://gobyexample.com/methods
package rect

// Rect is an axis-aligned rectangle.
type Rect struct {
	Width, Height float64
}

// Area returns the rectangle's area.
func Area(r Rect) float64 {
	return r.Width * r.Height
}

// Scale multiplies both sides of the rectangle by f.
func Scale(r Rect, f float64) {
	r.Width *= f
	r.Height *= f
}
