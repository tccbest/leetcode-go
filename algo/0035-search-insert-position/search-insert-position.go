package _035_search_insert_position

func SearchInsert(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	left := 0
	right := l - 1

	for left < right {
		mid := left + (right-left)>>1

		if target <= nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if target <= nums[left] {
		return left
	}

	if target > nums[right] {
		return right + 1
	}

	return left
}
