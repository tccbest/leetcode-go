package _083_remove_duplicates_from_sorted_list

import "leetcode-go/common"

func deleteDuplicates(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}

	current := head
	for current != nil {
		//链表的下一个节点值和当前节点值相等，则删除下一个节点
		if current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}
