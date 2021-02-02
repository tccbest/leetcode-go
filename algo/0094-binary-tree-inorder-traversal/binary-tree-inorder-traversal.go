package _094_binary_tree_inorder_traversal

import "leetcode-go/common"

func inorderTraversal(root *common.TreeNode) []int {
	ret := make([]int, 0)

	if root == nil {
		return ret
	}

	inorder(root, &ret)

	return ret
}

//递归
func inorder(root *common.TreeNode, ret *[]int) {
	if root != nil {
		inorder(root.Left, ret)
		*ret = append(*ret, root.Val)
		inorder(root.Right, ret)
	}
}

//非递归
func inorderTraversal2(root *common.TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}

	stack := []*common.TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
		root = root.Right
	}

	return ret
}
