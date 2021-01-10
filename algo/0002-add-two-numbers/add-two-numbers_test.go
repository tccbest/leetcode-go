package _002_add_two_numbers

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	a := &ListNode{Val:2}
	a.Next = &ListNode{Val:4}
	a.Next.Next = &ListNode{Val:3}

	b := &ListNode{Val:5}
	b.Next = &ListNode{Val:6}
	b.Next.Next = &ListNode{Val:4}

	PrintListNode(AddTwoNumbers(a, b))
}

func PrintListNode(node *ListNode)  {
	for {
		fmt.Print(node.Val)
		if node.Next == nil {
			fmt.Println()
			break
		}
		node = node.Next
	}
}