package bankex

import "testing"

func WithdrawTest(t *testing.T) {
	Deposit(10)
	Withdraw(5)

	balance := Balance()
	if balance != 5 {
		t.Errorf("Balance is incorrect. Got: %d, want: %d.", balance, 5)
	}
}
