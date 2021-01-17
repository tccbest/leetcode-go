package _154_find_minimum_in_rotated_sorted_array_ii

//右指针
func FindMin(numbers []int) int {
	left := 0
	right := len(numbers) - 1

	//本身有序，直接返回
	if numbers[left] < numbers[right] {
		return numbers[left]
	}

	for left < right { //循环退出条件，left=right
		mid := left + (right-left)>>1
		if numbers[mid] < numbers[right] { //中间元素小于最后元素，那么右侧单调有序
			right = mid //之所以是mid不是mid-1，有可能mid是最小元素
		} else if numbers[mid] > numbers[right] { //最右元素小于中间元素，那么左侧单调有序
			left = mid + 1
		} else {
			right--
		}
	}

	return numbers[left]
}

//左指针
func FindMin2(nums []int) int {
	left := 0
	right := len(nums) - 1

	//本身有序，直接返回
	if nums[left] < nums[right] {
		return nums[left]
	}

	for left < right { //循环退出条件，left=right
		if left >0 && nums[left] < nums[left - 1] {
			break
		}

		mid := left + (right-left)>>1
		if nums[left] < nums[mid] { //最左元素小于mid元素，那么左侧单调有序
			left = mid + 1
		} else if nums[mid] < nums[left] { //最左元素大于中间元素，那么右侧单调有序
			right = mid 	//之所以是mid不是mid-1，有可能mid是最小元素
		} else {
			left++
		}
	}

	return nums[left]
}