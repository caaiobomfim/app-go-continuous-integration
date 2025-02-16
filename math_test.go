package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(15, 15)

	if result != 30 {
		t.Errorf("RESULT IS INVALID: EXPECTED RESULT %d.: %d ", result, 30)
	}
}
