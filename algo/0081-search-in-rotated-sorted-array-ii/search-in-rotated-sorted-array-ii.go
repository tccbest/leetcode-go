package _081_search_in_rotated_sorted_array_ii

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return true
		}

		if nums[left] == nums[mid] {
			left++
			continue
		}

		//前半部分有序
		if nums[left] < nums[mid] {
			//target在前半
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			//后半部分有序
			//target在后半
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}
