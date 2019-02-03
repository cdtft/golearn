package main

import (
	"fmt"
)

//冒泡排序
func bubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}


func main() {
	intSlice := [...]int{4, 4, 5, 1, 7, 8}
	bubbleSort(intSlice[:])
	fmt.Println(intSlice)
}
