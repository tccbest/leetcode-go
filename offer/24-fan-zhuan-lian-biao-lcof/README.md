### 剑指 Offer 24. 反转链表
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

示例:
```
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
``` 

限制：

- 0 <= 节点个数 <= 5000

注意：本题与主站 206 题相同：https://leetcode-cn.com/problems/reverse-linked-list/


#### 题解
双指针
- 定义两个指针： prepre 和 curcur ；prepre 在前 curcur 在后。
- 每次让 prepre 的 nextnext 指向 curcur ，实现一次局部反转
- 局部反转完成之后， prepre 和 curcur 同时往前移动一个位置
- 循环上述过程，直至 prepre 到达链表尾部

https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/solution/fan-zhuan-lian-biao-yi-dong-de-shuang-zhi-zhen-jia/