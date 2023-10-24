package slice

func SumSlice(list []int) int {
	sum := 0
	for _, item := range list {
		sum += item
	}
	return sum
}
