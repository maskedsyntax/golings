package account

import (
	"errors"
	"sync"
	"testing"
	"time"
)

// within fails the test if f does not return within 2 seconds.
func within(t *testing.T, what string, f func()) {
	t.Helper()
	done := make(chan struct{})
	go func() { defer close(done); f() }()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatalf("%s did not return within 2s", what)
	}
}

func TestFailedWithdrawReleasesLock(t *testing.T) {
	var a Account
	a.Deposit(50)
	var err error
	within(t, "Withdraw", func() { err = a.Withdraw(80) })
	if !errors.Is(err, ErrInsufficient) {
		t.Fatalf("Withdraw(80) with balance 50 = %v, want ErrInsufficient", err)
	}
	within(t, "Deposit after a failed Withdraw (is the lock still held?)", func() { a.Deposit(10) })
	if got := a.Balance(); got != 60 {
		t.Fatalf("Balance = %d, want 60", got)
	}
}

func TestConcurrentWithdrawals(t *testing.T) {
	var a Account
	a.Deposit(100)
	var wg sync.WaitGroup
	var mu sync.Mutex
	ok := 0
	within(t, "200 concurrent withdrawals", func() {
		for range 200 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if a.Withdraw(1) == nil {
					mu.Lock()
					ok++
					mu.Unlock()
				}
			}()
		}
		wg.Wait()
	})
	if ok != 100 || a.Balance() != 0 {
		t.Fatalf("%d withdrawals of 1 succeeded from a balance of 100, leaving %d; want 100 and 0", ok, a.Balance())
	}
}
