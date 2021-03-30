package _152_maximum_product_subarray

import (
	"fmt"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	var nums []int

	nums = []int{2, 3, -2, 4}
	fmt.Println(MaxProduct(nums))

	nums = []int{-2, 0, -1}
	fmt.Println(MaxProduct(nums))

	nums = []int{-2}
	fmt.Println(MaxProduct(nums))
}
