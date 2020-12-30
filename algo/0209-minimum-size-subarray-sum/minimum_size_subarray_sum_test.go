package _209_minimum_size_subarray_sum

import (
	"fmt"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	s := 7
	nums := []int{2,3,1,2,4,3}

	a := 0

	a = MinSubArrayLen(s, nums)
	fmt.Println(a)

	a = MinSubArrayLen1(s, nums)
	fmt.Println(a)
}