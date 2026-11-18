package jobs

import (
	"strings"
	"sync/atomic"
	"testing"
)

func TestRunJobs(t *testing.T) {
	var done atomic.Int32
	ok := func() { done.Add(1) }
	jobs := []func(){
		ok,
		func() { panic("disk on fire") },
		ok,
		func() { var p *int; _ = *p },
		ok,
	}
	errs := RunJobs(jobs)
	if len(errs) != 5 {
		t.Fatalf("got %d errors, want 5 (one per job)", len(errs))
	}
	for _, i := range []int{0, 2, 4} {
		if errs[i] != nil {
			t.Errorf("errs[%d] = %v, want nil", i, errs[i])
		}
	}
	if errs[1] == nil || !strings.Contains(errs[1].Error(), "disk on fire") {
		t.Errorf("errs[1] = %v, want an error mentioning the panic", errs[1])
	}
	if errs[3] == nil || !strings.Contains(errs[3].Error(), "nil pointer") {
		t.Errorf("errs[3] = %v, want an error mentioning the nil pointer dereference", errs[3])
	}
	if done.Load() != 3 {
		t.Errorf("%d of 3 healthy jobs finished", done.Load())
	}
}
