package _8_shan_chu_lian_biao_de_jie_dian_lcof

import "leetcode-go/common"

func DeleteNode(head *common.ListNode, val int) *common.ListNode {
	temp := &common.ListNode{}
	temp.Next = head
	pre := temp
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			break
		}

		pre = cur
		cur = cur.Next
	}

	return temp.Next
}