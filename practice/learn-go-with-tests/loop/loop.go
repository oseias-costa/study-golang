package loop

func Loop(lether string) string {
	var result string

	for i := 0; i < 4; i++ {
		result += lether
	}
	return result
}
