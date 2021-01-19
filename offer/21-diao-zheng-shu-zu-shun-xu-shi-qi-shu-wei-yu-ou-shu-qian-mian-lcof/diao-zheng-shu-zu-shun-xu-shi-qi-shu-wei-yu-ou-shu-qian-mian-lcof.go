package _1_diao_zheng_shu_zu_shun_xu_shi_qi_shu_wei_yu_ou_shu_qian_mian_lcof

func Exchange(nums []int) []int {
	i := 0
	j := len(nums) - 1

	for i < j {
		if nums[i] % 2 == 0 && nums[j] % 2 != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}

		if nums[i] % 2 != 0 {
			i++
		}

		if nums[j] % 2 == 0 {
			j--
		}
	}

	return nums
}