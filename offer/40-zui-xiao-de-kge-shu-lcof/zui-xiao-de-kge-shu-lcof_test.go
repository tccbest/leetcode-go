package _0_zui_xiao_de_kge_shu_lcof

import (
	"fmt"
	"testing"
)

func TestGetLeastNumbers(t *testing.T) {
	arr := []int{3,2,1}
	k := 2
	fmt.Println(GetLeastNumbers(arr, k))

	arr = []int{0,1,2,1}
	k = 1
	fmt.Println(GetLeastNumbers(arr, k))

	arr = []int{0,0,2,3,2,1,1,2,0,4}
	k = 10
	fmt.Println(GetLeastNumbers(arr, k))
}
