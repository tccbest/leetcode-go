package _015_3sum

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	a := []int{-1, 0, 1, 2, -1, -4}

	resp := ThreeSum(a)

	fmt.Println(resp)
}
