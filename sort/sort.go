package sort

//冒泡排序
func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

//选择排序
func SelectSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

//快速排序
//TODO 1.快速排序的思想还未完全理解
//     		1.1递归终止条件
//			1.2交换逻辑
func QuickSort(array []int, left int, right int) {
	if left >= right {
		return
	}
	index := left
	value := array[index]
	for i := left + 1; i <= right; i++ {
		if array[i] < value {
			array[index] = array[i]
			array[i] = array[index+1]
			index++
		}
	}
	array[index] = value
	QuickSort(array, left, index-1)
	QuickSort(array, index+1, right)
}
