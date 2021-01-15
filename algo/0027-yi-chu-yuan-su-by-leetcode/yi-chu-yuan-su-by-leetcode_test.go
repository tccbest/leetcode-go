package _027_yi_chu_yuan_su_by_leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3,2,2,3}
	val := 3
	fmt.Println(RemoveElement(nums, val))

	nums = []int{0,1,2,2,3,0,4,2}
	val = 2
	fmt.Println(RemoveElement(nums, val))

	nums = []int{0,0}
	val = 0
	fmt.Println(RemoveElement(nums, val))
}