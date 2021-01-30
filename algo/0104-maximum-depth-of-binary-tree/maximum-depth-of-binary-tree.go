package _104_maximum_depth_of_binary_tree

import "leetcode-go/common"

//广度优先搜索
func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	ret := 0
	queue := []*common.TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		for l > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			l--
		}
		ret++
	}

	return ret
}


//递归
func maxDepth2(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth2(root.Left), maxDepth2(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
