package _033_search_in_rotated_sorted_array

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	a := []int{4,5,6,7,0,1,2}

	fmt.Println(Search(a, 0))
	fmt.Println(Search(a, 3))

	a = []int{1}
	fmt.Println(Search(a, 1))
	fmt.Println(Search(a, 2))

	a = []int{1, 3}
	fmt.Println(Search(a, 2))
	fmt.Println(Search(a, 3))
}