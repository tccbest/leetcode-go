package _034_find_first_and_last_position_of_element_in_sorted_array

func SearchRange(nums []int, target int) []int {
	ret := []int{-1, -1}
	n := len(nums)

	if n == 0 {
		return ret
	}

	ret[0] = searchLeft(nums, target)
	ret[1] = searchRight(nums, target)

	return ret
}

func searchLeft(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			right = mid
		} else {
			right = mid - 1
		}
	}

	if nums[left] == target {
		return left
	}

	return -1
}

func searchRight(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		mid := (left + right + 1) / 2	//+1避免left，right想临时，mid=target死循环
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			left = mid
		} else {
			right = mid - 1
		}
	}

	if nums[right] == target {
		return right
	}

	return -1
}
