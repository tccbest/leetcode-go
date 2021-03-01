package _081_search_in_rotated_sorted_array_ii

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		for left < right && nums[left] == nums[left+1] {
			left++
		}

		for left < right && nums[right] == nums[right-1] {
			right--
		}

		mid := left + (right - left) >> 1
		if nums[mid] == target {
			return true
		}

		if nums[left] < nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid
			} else {
				left = mid
			}
		} else if nums[mid] < nums[right] {
			if nums[mid] <= target && target <= nums[right] {
				left = mid
			} else {
				right = mid
			}
		}
	}

	if nums[left] == target || nums[right] == target {
		return true
	}

	return false
}
