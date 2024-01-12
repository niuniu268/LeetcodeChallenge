package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 递归
func maxdepth1(root *Methods.TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxdepth1(root.Left), maxdepth1(root.Right)) + 1
}

// 遍历
func maxdepth2(root *Methods.TreeNode) int {
	levl := 0
	queue := make([]*Methods.TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		for ; l > 0; l-- {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		levl++
		l = len(queue)
	}
	return levl
}
