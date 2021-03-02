package _014_longest_common_prefix

import "math"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minIndex := math.MaxInt64
	firstLen := len(strs[0])
	for i := 0; i < len(strs); i++ {
		index := 0
		for j := 0; j < len(strs[i]); j++ {
			if index < firstLen && index <= j {
				if strs[0][index] == strs[i][j] {
					index++
				} else {
					break
				}
			}
		}

		if index == 0 {
			return ""
		}

		if index < minIndex {
			minIndex = index
		}
	}

	return strs[0][:minIndex]
}
