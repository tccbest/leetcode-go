package _215_kth_largest_element_in_an_array

import (
	"fmt"
	"testing"
)

func TestFindKthLargest1(t *testing.T) {
	fmt.Println(FindKthLargest1([]int{3,2,3,1,2,4,5,5,6}, 4))

	fmt.Println(FindKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4))
}