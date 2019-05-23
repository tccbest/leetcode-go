package _206_reverse_linked_list

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	node1 := new(ListNode)
	node1.Val = 1
	node2 := new(ListNode)
	node2.Val = 2
	node3 := new(ListNode)
	node3.Val = 3
	node4 := new(ListNode)
	node4.Val = 4
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	temp := node1
	for node1 != nil {
		fmt.Printf("%d->", node1.Val)
		node1 = node1.Next
	}
	fmt.Println()

	newNode := ReverseList(temp)

	for newNode != nil {
		fmt.Printf("%d->", newNode.Val)
		newNode = newNode.Next
	}
	fmt.Println()
}