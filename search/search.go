package search

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
