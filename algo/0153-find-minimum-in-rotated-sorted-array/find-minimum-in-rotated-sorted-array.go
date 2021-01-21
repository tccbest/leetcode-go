package _153_find_minimum_in_rotated_sorted_array

func FindMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	if right == 0 {
		return nums[0]
	}

	for left < right {	//退出条件left=right
		mid := left + (right - left) >> 1
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return nums[left]
}