package main

import "fmt"

//选择排序
func selectSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

func main() {
	testArray := [...]int{3,4,1,8,234,23,135,55}
	selectSort(testArray[:])
	fmt.Print(testArray)
}
