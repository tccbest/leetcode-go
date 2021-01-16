package _6_cong_wei_dao_tou_da_yin_lian_biao_lcof

import "leetcode-go/common"

func ReversePrint(head *common.ListNode) []int {
	n := 0 //链表长度
	var ret []int
	for head != nil {
		ret = append(ret, head.Val)
		n++
		head = head.Next
	}

	//ret := make([]int, 0)
	//for i := n - 1; n >= 0; n-- {
	//	ret = append(ret, stack[i])
	//}

	//翻转数组
	for i := 0; i < n/2; i++ {
		ret[i], ret[n-1-i ] = ret[n-1-i ], ret[i]
	}

	return ret
}
