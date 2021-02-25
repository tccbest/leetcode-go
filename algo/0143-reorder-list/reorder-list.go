package _143_reorder_list

import "leetcode-go/common"

func reorderList(head *common.ListNode) {
	if head == nil {
		return
	}

	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tail := slow.Next
	slow.Next = nil
	tail = reverseList(tail)

	head = mergeTwoList(head, tail)
}

func reverseList(head *common.ListNode) *common.ListNode {
	var pre *common.ListNode
	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}

	return pre
}

func mergeTwoList(l1, l2 *common.ListNode) *common.ListNode {
	dummy := &common.ListNode{}
	head := dummy
	for l1 != nil && l2 != nil {
		head.Next = l1
		head = head.Next
		l1 = l1.Next

		head.Next = l2
		head = head.Next
		l2 = l2.Next
	}

	for l1 != nil {
		head.Next = l1
		head = head.Next
		l1 = l1.Next
	}

	for l2 != nil {
		head.Next = l2
		head = head.Next
		l2 = l2.Next
	}

	return dummy.Next
}
