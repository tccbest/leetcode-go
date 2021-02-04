package _237_delete_node_in_a_linked_list

import "leetcode-go/common"

func deleteNode(node *common.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}