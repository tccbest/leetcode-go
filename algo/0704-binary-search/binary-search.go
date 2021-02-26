package _704_binary_search

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	if len(nums) == 0 || nums[left] > target || nums[right] < target {
		return -1
	}

	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
