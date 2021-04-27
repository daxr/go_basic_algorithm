package main

import "fmt"

func main() {
	arr1 := []int{4, 1, 5, 2, 6, 7, 44, 55, 64, 354, 756, 345, 12, 235, 65, 8, 9, 12}
	fmt.Println(chaRu(arr1))
}

func chaRu(arr []int) []int {
	for i := range arr {
		pre := i - 1
		current := arr[i]
		for pre >= 0 && arr[pre] > current {
			arr[pre+1] = arr[pre]
			pre -= 1
		}
		arr[pre+1] = current
	}
	return arr
}
