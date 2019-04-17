package _019_remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//设置哨兵节点
	temp := &ListNode{0, nil}
	temp.Next = head

	//第一次遍历获取链表长度
	lLen := 0
	lenTemp := temp
	for nil != lenTemp.Next {
		lLen++
		lenTemp = lenTemp.Next
	}

	//第二次遍历获取待删除节点
	index := lLen - n
	first := temp
	for index > 0 {
		index--
		first = first.Next
	}

	//设置节点的指针
	first.Next = first.Next.Next

	return temp.Next
}
