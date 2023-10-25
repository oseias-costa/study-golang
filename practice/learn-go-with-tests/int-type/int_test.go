package inttype

import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 2)
	expect := 4

	if result != expect {
		t.Errorf("expect '%d', result '%d' ", expect, result)
	}
}
