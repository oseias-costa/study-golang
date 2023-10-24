package slice

import "testing"

func TestSumSlice(t *testing.T) {
	slice := []int{1, 2, 3, 20}
	result := SumSlice(slice)
	expect := 26

	if result != expect {
		t.Errorf("expect %d, result %d, slice %v", expect, result, slice)
	}
}
