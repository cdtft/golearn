package simple

import (
	"fmt"
	"testing"
)

func TestTowSum(t *testing.T) {

	nums := []int{3, 2, 4}
	result := towSum(nums, 6)
	fmt.Println(result)

}

func towSum(nums []int, target int) []int {

	for i := 0; i < len(nums) - 1; i++ {
		for j := i + 1; j < len(nums); j ++ {
			x := nums[j]
			y := nums[i]
			if target == x + y {
				return []int{j, i}
			}
		}
	}
	return []int{}
}