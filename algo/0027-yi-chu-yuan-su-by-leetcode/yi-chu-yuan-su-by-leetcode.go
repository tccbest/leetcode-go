package _027_yi_chu_yuan_su_by_leetcode

func RemoveElement(nums []int, val int) int {
	l := len(nums)
	i := 0
	for i < l {
		if nums[i] != val {
			i++
		} else {
			nums[i] = nums[l-1]
			l--
		}
	}

	return l
g}
