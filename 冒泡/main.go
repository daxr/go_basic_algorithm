package main

import "fmt"

func main() {
	arr1 := []int{1, 4, 5, 2, 6, 7, 44, 55, 64, 354, 756, 345, 12, 235, 65, 8, 9, 12}
	fmt.Println(maoPao(arr1))
}

func maoPao(arr []int) []int {
	ArrLen := len(arr)
	for i := 0; i < ArrLen; i++ {
		for j := 0; j < ArrLen-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
