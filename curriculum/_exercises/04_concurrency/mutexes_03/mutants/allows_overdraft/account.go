// Mutant: never refuses a withdrawal.
package account

import (
	"errors"
	"sync"
)

var ErrInsufficient = errors.New("insufficient funds")

type Account struct {
	mu      sync.Mutex
	balance int
}

func (a *Account) Deposit(n int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += n
}

func (a *Account) Withdraw(n int) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance -= n
	return nil
}

func (a *Account) Balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}
