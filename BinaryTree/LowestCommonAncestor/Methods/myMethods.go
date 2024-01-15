package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

func lowerCommonAncestor(root, p, q *Methods.TreeNode) *Methods.TreeNode {

	if root == nil || p == root || q == root {
		return root
	}

	var left *Methods.TreeNode
	var right *Methods.TreeNode

	left = lowerCommonAncestor(root.Left, p, q)
	right = lowerCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil && right != nil {
		return left
	}

	if left != nil && right == nil {
		return right
	}

	return nil

}
