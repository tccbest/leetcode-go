package _053_maximum_subarray

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}

	fmt.Println(MaxSubArray(nums))
}
