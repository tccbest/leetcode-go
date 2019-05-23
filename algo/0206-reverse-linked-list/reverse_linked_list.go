package _206_reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	var pre *ListNode

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
