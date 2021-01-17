package _154_find_minimum_in_rotated_sorted_array_ii

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {
	numbers := []int{3, 4, 5, 1, 2}
	fmt.Println("FindMin", numbers, FindMin(numbers))

	numbers = []int{2, 2, 2, 0, 1}
	fmt.Println("FindMin", numbers, FindMin(numbers))

	numbers = []int{2, 2, 2, 0, 1, 2}
	fmt.Println("FindMin", numbers, FindMin(numbers))

	numbers = []int{3, 4, 5, 1, 2}
	fmt.Println("FindMin2", numbers, FindMin2(numbers))

	numbers = []int{2, 2, 2, 0, 1}
	fmt.Println("FindMin2", numbers, FindMin2(numbers))

	numbers = []int{2, 2, 2, 0, 1, 2}
	fmt.Println("FindMin2", numbers, FindMin2(numbers))
}