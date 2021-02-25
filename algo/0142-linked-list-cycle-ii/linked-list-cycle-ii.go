package _142_linked_list_cycle_ii

import "leetcode-go/common"

//空间复杂度O(n)
func detectCycle(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}

	temp := make(map[*common.ListNode]struct{})
	for head != nil {
		if _, ok := temp[head]; ok {
			return head
		}

		temp[head] = struct{}{}
		head = head.Next
	}

	return nil
}

//空间复杂度O(1)
//快慢指针检测链表是否有环，如果有环，慢指针回到开头，快慢指针同时移动，相遇点即为环点
func detectCycle2(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}

	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}

			return slow
		}
	}

	return nil
}