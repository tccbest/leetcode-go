package _026_remove_duplicates_from_sorted_array

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	a := []int{0, 0, 1, 1, 1, 2, 3, 4, 5, 5, 6, 7, 7}

	fmt.Println(RemoveDuplicates(a))
}
