package main

import "testing"

func TestSum(t *testing.T) {
	ans := sum(1, 3)
	if ans != 4 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", ans, 4)
	}
}
