package _021_merge_two_sorted_lists

import "leetcode-go/common"

func MergeTwoLists(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	//设置哨兵节点，保存新链表
	dummy := &common.ListNode{}
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1	//值小的赋给当前节点的下一个节点
			cur = cur.Next  //指针往前移动
			l1 = l1.Next	//l1指针往前移动
		} else {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	return dummy.Next
}