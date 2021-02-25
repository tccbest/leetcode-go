package _234_palindrome_linked_list

import "leetcode-go/common"

func isPalindrome(head *common.ListNode) bool {
	if head == nil {
		return true
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tail := slow
	slow = nil

	tail = reverseList(tail)

	for head != nil && tail != nil{
		if head.Val != tail.Val {
			return false
		}

		head = head.Next
		tail = tail.Next
	}

	return true
}

func reverseList(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}

	var pre *common.ListNode
	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}

	return pre
}
