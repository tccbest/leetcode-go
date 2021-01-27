package _053_maximum_subarray

func MaxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		//如果前一个元素大于0，则nums[i]变成前一个元素和当前元素的和
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}
