package _082_remove_duplicates_from_sorted_list_ii

import "leetcode-go/common"

func deleteDuplicates(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &common.ListNode{}
	dummy.Next = head

	//双指针，b在a之前
	a := dummy
	b := head
	for b != nil && b.Next != nil {
		//a指向哑巴节点，比较a的下一个节点和b的下一个节点
		if a.Next.Val != b.Next.Val {
			a = a.Next
			b = b.Next
		} else {
			//循环b节点，如果值相等，则删除，直到值不相等
			for b != nil && b.Next != nil && a.Next.Val == b.Next.Val {
				b = b.Next
			}

			a.Next = b.Next
			b = b.Next
		}
	}

	return dummy.Next
}