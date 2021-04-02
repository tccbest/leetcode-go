package _392_is_subsequence

func IsSubsequence(s string, t string) bool {
	ls := len(s)
	lt := len(t)
	if ls > lt || ls == lt && s != t {
		return false
	}

	i, j := 0, 0	//i为s的索引，j为t的索引
	for i < ls && j < lt {
		if s[i] == t[j] {
			i++
		}

		j++
	}

	if i == ls && (j < lt || j == lt && s[i] == t[j]) {
		return true
	}

	return false
}