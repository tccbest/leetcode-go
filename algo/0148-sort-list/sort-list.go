package _148_sort_list

import "leetcode-go/common"

func sortList(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}

	return mergeSort(head)
}

func findMiddle(head *common.ListNode) *common.ListNode {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func mergeSort(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	//寻找中间节点
	middle := findMiddle(head)

	//从中间节点断开，分成两半
	tail := middle.Next
	middle.Next = nil

	left := mergeSort(head)
	right := mergeSort(tail)

	return mergeTwoList(left, right)
}

func mergeTwoList(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	dummy := &common.ListNode{}
	head := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}

		head = head.Next
	}

	if l1 != nil {
		head.Next = l1.Next
	}

	if l2 != nil {
		head.Next = l2.Next
	}

	return dummy.Next
}