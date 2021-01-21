package _153_find_minimum_in_rotated_sorted_array

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(nums, FindMin(nums))

	nums = []int{3, 1, 2}
	fmt.Println(nums, FindMin(nums))

	nums = []int{1, 2}
	fmt.Println(nums, FindMin(nums))

	nums = []int{2}
	fmt.Println(nums, FindMin(nums))

}
