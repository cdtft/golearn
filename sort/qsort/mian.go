package main

import "fmt"

//快速排序
func quickSort(array []int, left int, right int) {
	if left >= right {
		return
	}

	value := array[left]
	key := left
	for i := left + 1; i <= right; i++ {
		if array[i] > value {
			array[key] = array[i]
			array[i] = array[key+1]
			key++
		}
	}
	array[key] = value
	quickSort(array, left, key-1)
	quickSort(array, key+1, right)
}

func main() {
	intSlice := [...]int{0, 32, 2, 2432, 3, 123, 3, 23, 42, 123, 23}
	quickSort(intSlice[:], 0, len(intSlice)-1)
	fmt.Println(intSlice)
}
