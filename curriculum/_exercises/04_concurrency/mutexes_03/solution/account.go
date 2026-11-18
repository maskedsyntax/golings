// mutexes_03: The forgotten unlock (reference solution)
package account

import (
	"errors"
	"sync"
)

// ErrInsufficient is returned when a withdrawal exceeds the balance.
var ErrInsufficient = errors.New("insufficient funds")

// Account is a bank account that is safe for concurrent use.
type Account struct {
	mu      sync.Mutex
	balance int
}

// Deposit adds n to the balance.
func (a *Account) Deposit(n int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += n
}

// Withdraw removes n from the balance, or returns ErrInsufficient.
func (a *Account) Withdraw(n int) error {
	a.mu.Lock()
	defer a.mu.Unlock() // released on every return path
	if n > a.balance {
		return ErrInsufficient
	}
	a.balance -= n
	return nil
}

// Balance returns the current balance.
func (a *Account) Balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}
