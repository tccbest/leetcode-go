package _392_is_subsequence

import (
	"fmt"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	var s string
	var ts string

	s = "abc"
	ts = "ahbgdc"
	fmt.Println(IsSubsequence(s, ts))
}
