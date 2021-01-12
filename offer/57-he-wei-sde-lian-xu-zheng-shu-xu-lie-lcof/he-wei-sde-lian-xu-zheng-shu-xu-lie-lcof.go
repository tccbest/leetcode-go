package _7_he_wei_sde_lian_xu_zheng_shu_xu_lie_lcof

func FindContinuousSequence(target int) [][]int {
	l, r := 1, 2
	ret := [][]int{}
	for l < r {
		sum := (l + r) * (r - l + 1) / 2
		if sum == target {
			temp := []int{}
			for i := l; i <= r; i++ {
				temp = append(temp, i)
			}
			ret = append(ret, temp)
			l++
		} else if sum < target {
			r++
		} else {
			l++
		}
	}

	return ret
}