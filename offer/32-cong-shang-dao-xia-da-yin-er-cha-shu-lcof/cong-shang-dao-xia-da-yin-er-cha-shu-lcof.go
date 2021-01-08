package _2_cong_shang_dao_xia_da_yin_er_cha_shu_lcof

/**
剑指 Offer 32 - I. 从上到下打印二叉树
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。



例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：

[3,9,20,15,7]

提示：

节点总数 <= 1000
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}

	temp := []*TreeNode{root} //根节点
	for len(temp) > 0 {
		temp2 := []*TreeNode{} //队列
		for j := 0; j < len(temp); j++ {
			node := temp[j]
			ret = append(ret, node.Val)
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
