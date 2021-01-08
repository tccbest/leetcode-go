package _102_binary_tree_level_order_traversal

/**
102. 二叉树的层序遍历
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。



示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}

	temp := []*TreeNode{root} //根节点
	for i := 0; len(temp) > 0; i++ {
		temp2 := []*TreeNode{} //队列
		ret = append(ret, []int{})
		for j := 0; j < len(temp); j++ {
			node := temp[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				temp2 = append(temp2, node.Left)
			}
			if node.Right != nil {
				temp2 = append(temp2, node.Right)
			}
		}

		temp = temp2
	}

	return ret
}