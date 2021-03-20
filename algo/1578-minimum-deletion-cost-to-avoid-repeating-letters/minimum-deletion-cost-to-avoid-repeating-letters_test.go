package _578_minimum_deletion_cost_to_avoid_repeating_letters

import (
	"fmt"
	"testing"
)

func TestMinCost(t *testing.T) {
	s := "abaac"
	cost := []int{1, 2, 3, 4, 5}
	fmt.Println(MinCost(s, cost))

	s = "abc"
	cost = []int{1, 2, 3}
	fmt.Println(MinCost(s, cost))

	s = "aabaa"
	cost = []int{1, 2, 3, 4, 1}
	fmt.Println(MinCost(s, cost))

	s = "aaabbbabbbb"
	cost = []int{3,5,10,7,5,3,5,5,4,8,1}
	fmt.Println(MinCost(s, cost))
}
