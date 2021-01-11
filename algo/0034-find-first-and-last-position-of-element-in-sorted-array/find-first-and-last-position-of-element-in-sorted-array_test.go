package _034_find_first_and_last_position_of_element_in_sorted_array

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	/**
	示例 1：

	输入：nums = [5,7,7,8,8,10], target = 8
	输出：[3,4]
	示例 2：

	输入：nums = [5,7,7,8,8,10], target = 6
	输出：[-1,-1]
	示例 3：

	输入：nums = [], target = 0
	输出：[-1,-1]

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
	 */
	var nums []int
	var target int

	nums = []int{5,7,7,8,8,10}
	target = 8
	fmt.Println(SearchRange(nums, target))

	nums = []int{5,7,7,8,8,10}
	target = 6
	fmt.Println(SearchRange(nums, target))

	nums = []int{}
	target = 0
	fmt.Println(SearchRange(nums, target))
}