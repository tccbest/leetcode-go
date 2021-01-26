package _141_linked_list_cycle

import "leetcode-go/common"

/**
快慢指针
 */
func hasCycle(head *common.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

/**
使用map保存
 */
func hasCycle2(head *common.ListNode) bool {
	m := make(map[*common.ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}

		m[head] = struct{}{}
	}

	return false
}
