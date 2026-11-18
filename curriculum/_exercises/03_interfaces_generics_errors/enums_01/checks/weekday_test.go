package weekday

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	names := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	for i, want := range names {
		if got := Weekday(i).String(); got != want {
			t.Errorf("Weekday(%d).String() = %q, want %q", i, got, want)
		}
	}
	if got := fmt.Sprint(Wednesday); got != "Wednesday" {
		t.Errorf("fmt.Sprint(Wednesday) = %q", got)
	}
	for _, d := range []Weekday{7, 9, -1} {
		if got, want := d.String(), fmt.Sprintf("Weekday(%d)", int(d)); got != want {
			t.Errorf("Weekday(%d).String() = %q, want %q", int(d), got, want)
		}
	}
}

func TestIsWeekend(t *testing.T) {
	for d := Sunday; d <= Saturday; d++ {
		want := d == Saturday || d == Sunday
		if got := d.IsWeekend(); got != want {
			t.Errorf("%v.IsWeekend() = %v, want %v", d, got, want)
		}
	}
}

func TestNext(t *testing.T) {
	if Monday.Next() != Tuesday || Friday.Next() != Saturday {
		t.Error("Next() doesn't return the following day")
	}
	if got := Saturday.Next(); got != Sunday {
		t.Errorf("Saturday.Next() = %v, want Sunday", got)
	}
}
