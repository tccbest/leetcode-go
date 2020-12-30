package _209_minimum_size_subarray_sum

import (
	"fmt"
	"math"
)

// 双重循环，暴力破解法
// 时间复杂度O(n²)
func MinSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	ret := math.MaxInt32
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum >= s {
				if j-i+1 < ret {
					ret = j - i + 1
				}
				break
			}
		}
	}

	if ret == math.MaxInt32 {
		return 0
	}

	return ret
}

// 滑动窗口，双指针
// 时间复杂度O(n)
func MinSubArrayLen1(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	start, end := 0, 0
	ret := math.MaxInt32
	sum := 0
	for end < n {
		fmt.Println(start, end)
		sum += nums[end]
		for sum >= s {
			subArrLen := end - start + 1
			if ret > subArrLen {
				ret = subArrLen
			}

			sum -= nums[start]
			start++
		}

		end++
	}

	if ret == math.MaxInt32 {
		return 0
	}

	return ret
}
