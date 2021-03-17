package _125_valid_palindrome

import (
	"strings"
)

func IsPalindrome(s string) bool {
	if s == "" {
		return true
	}

	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		if !isOkStr(string(s[left])) {
			left++
			continue
		}

		if !isOkStr(string(s[right])) {
			right--
			continue
		}

		if string(s[left]) != string(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func isOkStr(s string) bool {
	if s >= "a" && s <= "z" || s >= "A" && s <= "Z" || s >= "0" && s <= "9" {
		return true
	}

	return false
}
