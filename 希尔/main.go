package main

import "fmt"

func main() {
	arr := []int{4, 1, 5, 2, 6, 7, 44, 55, 64, 354, 756, 345, 12, 235, 65, 8, 9, 12}
	shellSort(arr, len(arr)/2)
	fmt.Println(arr)
}

/**
希尔排序：把切片分成n个batch，对每个batch进行插入排序；然后减小batch，再对每个batch进行插入排序；直到bathc等于1
*/
func shellSort(arr []int, batchSize int) {
	if batchSize < 1 {
		return
	}
	// k : 每个batch中的元素所在batch的index， 介于[0, batchSize)
	for k := 0; k < batchSize; k++ {
		// 用到了插入排序
		for j := 1; batchSize*j+k < len(arr); j++ { // j: 用来获取每个batch所在的第k个元素，拥有多少个batch
			for n := j; n > 0; n-- {
				pre := batchSize*(n-1) + k
				next := batchSize*n + k
				if arr[next] < arr[pre] {
					arr[next], arr[pre] = arr[pre], arr[next]
				}
			}

		}
	}
	shellSort(arr, batchSize/2)
}
