package _338_reduce_array_size_to_the_half

import (
	"sort"
)

func MinSetSize(arr []int) int {
	times := make(map[int]int)
	for _, v := range arr {
		times[v]++
	}

	var val []int
	for _, v := range times {
		val = append(val, v)
	}

	sort.Ints(val)

	temp := 0
	for i := len(val) - 1; i >= 0; i-- {
		temp += val[i]
		if temp >= len(arr)/2 {
			return len(val) - i
		}
	}

	return 0
}
