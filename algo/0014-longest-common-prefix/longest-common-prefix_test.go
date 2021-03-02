package _014_longest_common_prefix

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower","flow","flight"}
	fmt.Println(LongestCommonPrefix(strs))
}