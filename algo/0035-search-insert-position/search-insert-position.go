package _035_search_insert_position

func SearchInsert(nums []int, target int) int {
	l := len(nums)
	if l == 0 || target < nums[0] {
		return 0
	}

	if target > nums[l - 1] {
		return l
	}

	left := 0
	right := l - 1
	for left <= right {
		mid := left + (right-left)>>1
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}

	if target < nums[left] {
		return left
	}

	if target > nums[right] {
		return right + 1
	}

	return left
}
