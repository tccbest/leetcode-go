package _107_binary_tree_level_order_traversal_ii

import "leetcode-go/common"

func levelOrderBottom(root *common.TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}

	queue := []*common.TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		temp := []int{}
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		ret = append(ret, temp)
	}

	l := len(ret)
	for i := 0; i < l/2; i++ {
		ret[i], ret[l-1-i] = ret[l-1-i], ret[i]
	}

	return ret
}
