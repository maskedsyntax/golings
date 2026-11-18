package safe

import "testing"

func noPanic(t *testing.T, what string, f func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("%s panicked: %v", what, r)
		}
	}()
	f()
}

func TestLast(t *testing.T) {
	var v int
	var ok bool
	noPanic(t, "Last([1 2 3])", func() { v, ok = Last([]int{1, 2, 3}) })
	if v != 3 || !ok {
		t.Fatalf("Last([1 2 3]) = %d, %v; want 3, true", v, ok)
	}
	noPanic(t, "Last(nil)", func() { v, ok = Last(nil) })
	if v != 0 || ok {
		t.Fatalf("Last(nil) = %d, %v; want 0, false", v, ok)
	}
}

func TestZeroRegistry(t *testing.T) {
	var r Registry
	noPanic(t, "Count on a zero Registry", func() {
		if r.Count("x") != 0 {
			t.Fatal("a zero Registry should count 0")
		}
	})
	noPanic(t, "Add on a zero Registry", func() {
		r.Add("go")
		r.Add("go")
		r.Add("rust")
	})
	if r.Count("go") != 2 || r.Count("rust") != 1 {
		t.Fatalf("counts = %d, %d; want 2, 1", r.Count("go"), r.Count("rust"))
	}
}
