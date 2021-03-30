package _152_maximum_product_subarray

import "math"

func MaxProduct(nums []int) int {
	tmax, tmin, rmax := 1, 1, math.MinInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			tmax, tmin = tmin, tmax
		}

		tmax = max(nums[i], tmax*nums[i])
		tmin = min(nums[i], tmin*nums[i])

		rmax = max(rmax, tmax)
	}

	return rmax
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
