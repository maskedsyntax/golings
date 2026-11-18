package rect

import "testing"

func TestArea(t *testing.T) {
	if got := (Rect{Width: 2, Height: 3}).Area(); got != 6 {
		t.Errorf("Rect{2, 3}.Area() = %v, want 6", got)
	}
	if got := (Rect{}).Area(); got != 0 {
		t.Errorf("Rect{}.Area() = %v, want 0", got)
	}
}

func TestScaleModifiesReceiver(t *testing.T) {
	r := Rect{Width: 2, Height: 3}
	r.Scale(2)
	if r != (Rect{Width: 4, Height: 6}) {
		t.Errorf("after r.Scale(2), r = %+v, want {Width:4 Height:6}", r)
	}
	p := &r
	p.Scale(0.5)
	if r != (Rect{Width: 2, Height: 3}) {
		t.Errorf("after p.Scale(0.5), r = %+v, want {Width:2 Height:3}", r)
	}
}
