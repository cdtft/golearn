package search

import (
	"demo/sort"
	"testing"
)

var array = []int{1, 23, 44, 3, 442, 32, 44}

//【二分查找】测试
func TestBinarySearch(t *testing.T) {
	sort.CountingSort(array, len(array))
	t.Log(array)
	index, count := BinarySearch(array, len(array), 32)
	t.Log(index, count)
}

//【二分查找】递归排序
func TestBinarySearchRecursion(t *testing.T) {
	sort.CountingSort(array, len(array))
	t.Log(array)
	index := BinarySearchRecursion(array, 0, len(array)-1, 32)
	t.Log(index)
}

//【二分查找】变形一
func TestBinarySearch_1(t *testing.T) {
	sort.QuickSort(array, 0, len(array)-1)
	t.Log(array)
	index, count := BinarySearch_1(array, len(array), 44)
	t.Log(index, count)
}
