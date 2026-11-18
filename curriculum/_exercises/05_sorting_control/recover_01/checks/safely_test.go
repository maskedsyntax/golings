package safely

import "testing"

func TestSafelyNoPanic(t *testing.T) {
	ran := false
	if err := Safely(func() { ran = true }); err != nil || !ran {
		t.Fatalf("Safely(ok) = %v (ran=%v), want nil", err, ran)
	}
}

func TestSafelyRecovers(t *testing.T) {
	var err error
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("the panic escaped from Safely: %v", r)
			}
		}()
		err = Safely(func() { panic("a problem") })
	}()
	if err == nil || err.Error() != "recovered: a problem" {
		t.Fatalf("Safely(panicking) = %v, want \"recovered: a problem\"", err)
	}
}

func TestSafelyRuntimePanic(t *testing.T) {
	err := Safely(func() {
		var m map[string]int
		m["x"] = 1
	})
	if err == nil {
		t.Fatal("Safely should report a runtime panic as an error")
	}
}
