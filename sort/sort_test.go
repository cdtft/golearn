package sort

import (
	"fmt"
	"testing"
)

//【冒泡排序】测试
func TestBubbleSort(t *testing.T) {
	array := []int{1,23,44,3,442,32,44}
	BubbleSort(array)
	fmt.Println(array)
}

//【选择排序】测试
func TestSelectSort(t *testing.T) {
	array := []int{1,23,44,3,442,32,44}
	SelectSort(array)
	fmt.Println(array)
}