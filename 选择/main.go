package main

import "fmt"

func main() {
	arr1 := []int{1, 4, 5, 2, 6, 7, 44, 55, 64, 354, 756, 345, 12, 235, 65, 8, 9, 12}
	fmt.Println(xuAnZhe(arr1))
}

func xuAnZhe(arr []int) []int {
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		min := i
		for j := i + 1; j < arrLen; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
