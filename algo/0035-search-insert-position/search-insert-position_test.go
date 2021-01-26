package _035_search_insert_position

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	var nums []int
	var target int

	nums = []int{1,3,5,6}
	target = 5
	fmt.Println(SearchInsert(nums, target))	//2

	nums = []int{1,3,5,6}
	target = 2
	fmt.Println(SearchInsert(nums, target))	//1

	nums = []int{1,3,5,6}
	target = 7
	fmt.Println(SearchInsert(nums, target))	//4

	nums = []int{1,3,5,6}
	target = 0
	fmt.Println(SearchInsert(nums, target))	//0
}
