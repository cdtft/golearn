package search

//二分查找
func BinarySearch(array []int, length int, value int) (index int, count int) {
	_count := 0
	start := 0
	end := length - 1
	for ; start <= end; {
		_count++
		mid := (start + end) / 2
		if array[mid] == value {
			return mid, _count
		}
		if array[mid] < value {
			start = mid + 1
		}
		if array[mid] > value {
			end = mid - 1
		}
	}
	return -1, _count
}

//二分查找递归实现
func BinarySearchRecursion(array []int, low int, high int, value int) int {
	if low >= high {
		return -1
	}
	mid := (high - low) / 2
	if array[mid] == value {
		return mid
	}
	if array[mid] < value {
		return BinarySearchRecursion(array, low, mid-1, value)
	} else {
		return BinarySearchRecursion(array, mid+1, high, value)
	}
}

//二分查找变形问题一，查找第一个等于给定值位置
func BinarySearch_1(array []int, length int, value int) (index int, count int) {
	_count := 0
	start := 0
	end := length - 1
	for ; start <= end; {
		_count++
		mid := end + (start-end)>>1
		if array[mid] == value {
			if mid == 0 || array[mid-1] != value {
				return mid, _count
			}
			end = mid - 1
		}
		if array[mid] < value {
			start = mid + 1
		}
		if array[mid] > value {
			end = mid - 1
		}
	}
	return -1, _count
}
