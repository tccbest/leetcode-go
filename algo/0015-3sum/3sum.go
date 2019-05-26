package _015_3sum

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	resp := [][]int{}

	//对数组排序
	sort.Ints(nums)
	numsLen := len(nums)

	for i := range nums {
		//避免添加重复结果
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//定义两个索引
		j, k := i+1, numsLen-1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			switch {
			case sum < 0:
				j++
			case sum > 0:
				k--
			default:
				resp = append(resp, []int{nums[i], nums[j], nums[k]})

				//避免重复添加，j、k需要都移动到不同元素上
				for j < k {
					switch {
					case nums[j] == nums[j+1]:
						j++
					case nums[k] == nums[k-1]:
						k--
					default:
						j++
						k--
						goto end
					}
				}
				end:
			}
		}
	}

	return resp
}
