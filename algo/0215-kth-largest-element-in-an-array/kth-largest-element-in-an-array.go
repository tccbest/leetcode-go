package _215_kth_largest_element_in_an_array

import (
	"math/rand"
	"sort"
	"time"
)

// 解法一 排序
func FindKthLargest1(nums []int, k int) int {
	sort.Ints(nums)

	return nums[len(nums)-k]
}

func FindKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())

	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l, r, index int) int {
	i := partition(nums, l, r)
	if i == index {
		return nums[i]
	} else if i < index {
		return quickSelect(nums, i+1, r, index)
	} else {
		return quickSelect(nums, l, i-1, index)
	}
}

func partition(nums []int, l, r int) int {
	x := nums[r]

	i := l	//i用来存放已排序的
	for j := l; j < r; j++ {	//j用来存放未排序的，j指针向后移，比较和基准数大小
		if nums[j] <= x {	//当前数小于等于x，则需要交换，放左边
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	//交换i和r，i就是基准点位置
	nums[i], nums[r] = nums[r], nums[i]

	return i
}
