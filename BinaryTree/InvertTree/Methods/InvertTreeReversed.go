package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

func invertTree(root *Methods.TreeNode) *Methods.TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left //交换

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
