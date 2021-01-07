package _29_shun_shi_zhen_da_yin_ju_zhen_lcof

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	a := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	ret := SpiralOrder(a)

	fmt.Println(ret)
}
