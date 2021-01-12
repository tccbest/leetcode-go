package _088_merge_sorted_array

import (
	"fmt"
	"testing"
)


func TestMerge(t *testing.T) {
	//示例 1：
	//输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
	//输出：[1,2,2,3,5,6]
	nums1 := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	Merge(nums1, m, nums2, n)
	fmt.Println(nums1)


	//示例 2：
	//输入：nums1 = [1], m = 1, nums2 = [], n = 0
	//输出：[1]
	nums1 = []int{0}
	m = 0
	nums2 = []int{1}
	n = 1
	Merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}