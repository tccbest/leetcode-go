package _2_lian_biao_zhong_dao_shu_di_kge_jie_dian_lcof

import "leetcode-go/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func GetKthFromEnd(head *common.ListNode, k int) *common.ListNode {
	l, r := head, head
	for i := 0; i < k; i++ {
		l = l.Next
	}

	for l != nil {
		l = l.Next
		r = r.Next
	}

	return r
}
