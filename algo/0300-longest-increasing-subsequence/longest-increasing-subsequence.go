package _300_longest_increasing_subsequence

import "fmt"

func LengthOfLIS(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	ret := 1
	for i := 0; i < l; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		ret = max(ret, dp[i])
	}

	fmt.Println(dp)

	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
