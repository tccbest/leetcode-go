package _144_binary_tree_preorder_traversal

import "leetcode-go/common"

func preorderTraversal(root *common.TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}

	preorder(root, &ret)

	return ret
}

func preorder(root *common.TreeNode, ret *[]int) {
	if root != nil {
		*ret = append(*ret, root.Val)
		preorder(root.Left, ret)
		preorder(root.Right, ret)
	}
}

func preorderTraversal2(root *common.TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}

	stack := []*common.TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != nil {
			ret = append(ret, node.Val)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return ret
}
