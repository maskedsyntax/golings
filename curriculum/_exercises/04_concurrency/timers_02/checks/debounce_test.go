package debounce

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestDebounceBurst(t *testing.T) {
	var calls atomic.Int32
	var firedAt atomic.Int64
	db := NewDebouncer(60*time.Millisecond, func() {
		calls.Add(1)
		firedAt.Store(time.Now().UnixNano())
	})
	var last time.Time
	for range 8 {
		db.Trigger()
		last = time.Now()
		time.Sleep(10 * time.Millisecond)
	}
	time.Sleep(250 * time.Millisecond)
	if n := calls.Load(); n != 1 {
		t.Fatalf("f ran %d times for one burst of 8 triggers, want 1", n)
	}
	if gap := time.Unix(0, firedAt.Load()).Sub(last); gap < 50*time.Millisecond {
		t.Fatalf("f ran %v after the last trigger, want about 60ms", gap)
	}
}

func TestDebounceTwoBursts(t *testing.T) {
	var calls atomic.Int32
	db := NewDebouncer(30*time.Millisecond, func() { calls.Add(1) })
	db.Trigger()
	db.Trigger()
	time.Sleep(150 * time.Millisecond)
	db.Trigger()
	time.Sleep(150 * time.Millisecond)
	if n := calls.Load(); n != 2 {
		t.Fatalf("f ran %d times for two separate bursts, want 2", n)
	}
}

func TestNoTriggerNoCall(t *testing.T) {
	var calls atomic.Int32
	NewDebouncer(10*time.Millisecond, func() { calls.Add(1) })
	time.Sleep(60 * time.Millisecond)
	if calls.Load() != 0 {
		t.Fatal("f ran without any Trigger")
	}
}
