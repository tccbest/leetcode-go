package _1_diao_zheng_shu_zu_shun_xu_shi_qi_shu_wei_yu_ou_shu_qian_mian_lcof

import (
	"fmt"
	"testing"
)

func TestExchange(t *testing.T) {
	nums := []int{1,2,3,4}
	fmt.Println(nums, Exchange(nums))

	nums = []int{1}
	fmt.Println(nums, Exchange(nums))
}