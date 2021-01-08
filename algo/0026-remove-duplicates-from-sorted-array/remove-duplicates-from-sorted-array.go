package _026_remove_duplicates_from_sorted_array

func RemoveDuplicates(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}

	//双指针
	i := 0	//左指针，慢指针
	for j := 1; j < length; j++ {	//右指针，快指针
		 if nums[i] != nums[j] {
		 	nums[i + 1] = nums[j]
		 	i++
		 }
	}

	return i + 1
}
