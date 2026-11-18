// Mutant: Scale has a value receiver, so it scales a copy.
package rect

type Rect struct {
	Width, Height float64
}

func (r Rect) Area() float64 { return r.Width * r.Height }

func (r Rect) Scale(f float64) {
	r.Width *= f
	r.Height *= f
	_ = r
}
