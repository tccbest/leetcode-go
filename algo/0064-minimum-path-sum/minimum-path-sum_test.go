package _064_minimum_path_sum

import (
	"fmt"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	var grid [][]int

	grid = [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	fmt.Println(MinPathSum(grid))

	grid = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}
	fmt.Println(MinPathSum(grid))
}
