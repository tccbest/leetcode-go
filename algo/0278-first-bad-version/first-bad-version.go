package _278_first_bad_version

func firstBadVersion(n int) int {
	left := 0
	right := n

	for left <= right {
		mid := left + (right - left) >> 1
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if isBadVersion(left) {
		return left
	}

	return right
}

func isBadVersion(n int) bool {
	return true
}