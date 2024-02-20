package main

import "fmt"

func main() {
	arr := []int{45, 14, 25, 36, 4}
	fmt.Println(insertionSort(arr))
}

func insertionSort(arr []int) []int {
	var atual, index, pos int
	for index = 1; index < len(arr); index++ {
		atual = arr[index]
		for pos = index - 1; pos >= 0 && arr[pos] > atual; pos-- {
			arr[pos+1], arr[pos] = arr[pos], arr[pos+1]
		}
		fmt.Printf("esse é o arr, %v\n", arr)
	}
	return arr
}
