// Mutant: the balance is fixed but overdrawn is still shadowed.
package variables

func Balance(transactions []int) (balance int, overdrawn bool) {
	for _, t := range transactions {
		balance += t
		if balance < 0 {
			overdrawn := true
			_ = overdrawn
		}
	}
	return balance, overdrawn
}
