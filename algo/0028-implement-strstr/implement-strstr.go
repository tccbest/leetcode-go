package _028_implement_strstr

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	left, right := 0, len(needle)
	for right <= len(haystack) {
		if needle == haystack[left:right] {
			return left
		}

		left++
		right++
	}

	return -1
}
