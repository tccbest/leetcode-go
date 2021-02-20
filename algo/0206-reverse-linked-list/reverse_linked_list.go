package _206_reverse_linked_list

import "leetcode-go/common"

//链表反转 - 迭代
func ReverseList(head *common.ListNode) *common.ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	var pre *common.ListNode

	cur := head
	for cur != nil {
		// temp := cur.Next
		// cur.Next = pre
		// pre = cur
		// cur = temp
		cur.Next, pre, cur = pre, cur, cur.Next
	}

	return pre
}

//递归
func ReverseList2(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := ReverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
