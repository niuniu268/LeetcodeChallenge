package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

func Merge(node1, node2 *Methods.TreeNode) *Methods.TreeNode {

	if node1 != nil && node2 == nil {

		return node1
	}

	if node2 != nil && node1 == nil {
		return node2
	}

	node1.Val += node2.Val

	node1.Left = Merge(node1.Left, node2.Left)
	node1.Right = Merge(node1.Right, node2.Right)

	return node1
}
