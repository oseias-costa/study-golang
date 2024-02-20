package main

import "fmt"

func main() {
	arr := []int{42, 23, 4, 89, 7}
	fmt.Println(selectionSort(arr))
}

func selectionSort(arr []int) []int {
	var pos int
	var minor int
	for key, item := range arr {
		pos = key
		minor = item
		for i, verify := range arr {
			if verify < item && i > key {
				minor = verify
				pos = i
			}
		}
		if minor < item {
			arr[key], arr[pos] = arr[pos], arr[key]
		}
		fmt.Println(arr)
	}
	return arr
}
