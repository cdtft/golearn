package search

import (
	"demo/sort"
	"testing"
)

var array = []int{1, 23, 44, 3, 442, 32, 44}

func TestBinarySearch(t *testing.T) {
	sort.CountingSort(array, len(array))
	t.Log(array)
	index, count := BinarySearch(array, len(array), 32)
	t.Log(index, count)
}

func TestBinarySearchRecursion(t *testing.T) {
	sort.CountingSort(array, len(array))
	t.Log(array)
	index := BinarySearchRecursion(array, 0, len(array)-1, 32)
	t.Log(index)
}
