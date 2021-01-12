package _069_sqrtx

func MySqrt(x int) int {
	left, right := 0, x
	ret := -1
	for left < right {
		mid := left + (right - left) / 2
		if mid * mid <= x {
			left = mid + 1
			ret = mid
		} else {
			right = mid - 1
		}
	}

	return ret
}