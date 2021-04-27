package main

import "fmt"

func main() {
	arr1 := []int{4, 1, 5, 2, 6, 7, 44, 55, 64, 354, 756, 345, 12, 235, 65, 8, 9, 12}
	fmt.Println(quickSort(arr1))
}

func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index += 1
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
