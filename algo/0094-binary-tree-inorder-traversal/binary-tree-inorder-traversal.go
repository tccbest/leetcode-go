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
