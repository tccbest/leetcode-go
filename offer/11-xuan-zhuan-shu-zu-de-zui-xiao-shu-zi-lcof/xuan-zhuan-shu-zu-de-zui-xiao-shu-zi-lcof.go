package _1_xuan_zhuan_shu_zu_de_zui_xiao_shu_zi_lcof

import "fmt"

func MinArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1

	//本身有序，直接返回
	if numbers[left] < numbers[right] {
		return numbers[left]
	}

	for left < right {	//循环退出条件，left=right
		mid := left + (right - left) >> 1
		if numbers[mid] < numbers[right] {	//最左元素小于mid元素，那么左侧单调有序
			right = mid	//之所以是mid不是mid-1，有可能mid是最小元素
		} else if numbers[mid] > numbers[right] {	//最右元素大于mid元素，那么右侧单调有序
			left = mid + 1
		} else {
			right--
		}
	}

	return numbers[left]
}