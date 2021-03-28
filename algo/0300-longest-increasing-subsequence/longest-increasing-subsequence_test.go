package _300_longest_increasing_subsequence

import (
	"fmt"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	nums := []int{1,3,6,7,9,4,10,5,6}
	fmt.Println(LengthOfLIS(nums))

	nums = []int{10,9,2,5,3,7,101,18}
	fmt.Println(LengthOfLIS(nums))
}
