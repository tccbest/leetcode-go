package _086_partition_list

import "leetcode-go/common"

func partition(head *common.ListNode, x int) *common.ListNode {
	left := &common.ListNode{}
	right := &common.ListNode{}

	leftHead := left
	rightHead := right

	for head != nil {
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}

		head = head.Next
	}

	right.Next = nil
	left.Next = rightHead.Next

	return leftHead.Next
}
