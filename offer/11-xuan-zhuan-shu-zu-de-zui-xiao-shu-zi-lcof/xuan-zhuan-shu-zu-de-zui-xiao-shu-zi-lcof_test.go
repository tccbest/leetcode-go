package _1_xuan_zhuan_shu_zu_de_zui_xiao_shu_zi_lcof

import (
	"fmt"
	"testing"
)

func TestMinArray(t *testing.T) {
	numbers := []int{3, 4, 5, 1, 2}
	fmt.Println(numbers, MinArray(numbers))

	numbers = []int{2, 2, 2, 0, 1}
	fmt.Println(numbers, MinArray(numbers))

	numbers = []int{2, 2, 2, 0, 1, 2}
	fmt.Println(numbers, MinArray(numbers))
}
