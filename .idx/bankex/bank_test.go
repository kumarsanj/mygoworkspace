package bankex

import "testing"

func TestWithdrawHappy(t *testing.T) {
	Deposit(10)
	Withdraw(5)

	balance := Balance()
	if balance != 5 {
		t.Errorf("Balance is incorrect. Got: %d, want: %d.", balance, 5)
	}
}

func TestWithdrawNegative(t *testing.T) {
	Withdraw(5)
	Withdraw(10)
	balance := Balance()
	if balance != 0 {
		t.Errorf("Balance is incorrect. Got: %d, want: %d.", balance, 0)
	}
}
