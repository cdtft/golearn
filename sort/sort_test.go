package sort

import (
	"fmt"
	"testing"
)

var array = []int{1, 23, 44, 3, 442, 32, 44}

//【冒泡排序】测试
func TestBubbleSort(t *testing.T) {
	BubbleSort(array)
	fmt.Println(array)
}

//【选择排序】测试
func TestSelectSort(t *testing.T) {
	SelectSort(array)
	fmt.Println(array)
}

//【快速排序】测试
func TestQuickSort(t *testing.T) {
	QuickSort(array, 0, len(array)-1)
	fmt.Println(array)
}

func TestMergeSort(t *testing.T) {
	MergeSort(array)
	fmt.Println(array)
}
