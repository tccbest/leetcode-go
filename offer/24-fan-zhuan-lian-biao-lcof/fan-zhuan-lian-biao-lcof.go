package _4_fan_zhuan_lian_biao_lcof

import "leetcode-go/common"

/**
1.定义两个指针： pre 和 cur ；pre 在前 cur 在后。
2.每次让 pre 的 next 指向 cur ，实现一次局部反转
3.局部反转完成之后， pre 和 cur 同时往前移动一个位置
4.循环上述过程，直至 pre 到达链表尾部
*/
func ReverseList(head *common.ListNode) *common.ListNode {
	var pre *common.ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next	//保存next链表
		cur.Next = pre	//单个元素反转
		pre = cur	//pre指针向前移动，指向cur的位置
		cur = tmp	//cur向前移动
	}

	return pre
}
