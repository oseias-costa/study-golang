package arrays

import "testing"

func TestSum(t *testing.T) {

	t.Run("colection with five numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		result := Sum(numbers)
		expect := 15

		if result != expect {
			t.Errorf("expect %d, result %d, data %v", expect, result, numbers)
		}
	})
}
