package main

import "fmt"

//选择排序
func insertSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j] < array[j -1] {
				break
			}
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
}

func main()  {
	intSlice := [...]int{0}
	insertSort(intSlice[:])
	fmt.Println(intSlice)
}
