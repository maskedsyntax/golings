// methods_01: Receivers (reference solution)
package rect

// Rect is an axis-aligned rectangle.
type Rect struct {
	Width, Height float64
}

// Area returns the rectangle's area.
func (r Rect) Area() float64 {
	return r.Width * r.Height
}

// Scale multiplies both sides of the rectangle by f.
func (r *Rect) Scale(f float64) {
	r.Width *= f
	r.Height *= f
}
