package sort

//冒泡排序
func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array) - 1 - i; j++ {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

//选择排序
func SelectSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i+1; j < len(array); j++ {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}