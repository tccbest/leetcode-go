package _092_reverse_linked_list_ii

import "leetcode-go/common"

func reverseBetween(head *common.ListNode, left int, right int) *common.ListNode {
	// 思路：先遍历到m处，翻转，再拼接后续，注意指针处理
	// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
	if head == nil {
		return head
	}

	dummy := &common.ListNode{}
	dummy.Next = head
	head = dummy

	pre := &common.ListNode{}

	i := 0
	for i < left {
		pre = head
		head = head.Next
		i++
	}

	j := i
	next := &common.ListNode{}
	mid := head

	for j <= right && head != nil {
		temp := head.Next
		head.Next = next
		next = head
		head = temp
		j++
	}

	pre.Next = next
	mid.Next = head

	return dummy.Next
}

func reverseBetween2(head *common.ListNode, left int, right int) *common.ListNode {
	// 思路：先遍历到m处，翻转，再拼接后续，注意指针处理
	// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
	if head == nil {
		return head
	}

	dummy := &common.ListNode{}
	dummy.Next = head
	head = dummy

	pre := &common.ListNode{}

	i := 0
	for i < left {
		pre = head
		head = head.Next
		i++
	}

	j := i
	next := &common.ListNode{}
	mid := head

	for j <= right && head != nil {
		temp := head.Next
		head.Next = next
		next = head
		head = temp
		j++
	}

	pre.Next = next
	mid.Next = head

	return dummy.Next
}
