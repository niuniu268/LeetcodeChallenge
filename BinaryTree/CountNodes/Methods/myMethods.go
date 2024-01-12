package Methods

import (
	"Leetcode/BinaryTree/OrderTraversal/InOrderTraversal/Methods"
)

func LeftLevel(node *Methods.TreeNode, leftDepth int) int {
	if node == nil {
		return leftDepth
	}

	return LeftLevel(node.Left, leftDepth) + 1
}

func RightLevel(node *Methods.TreeNode, rightDepth int) int {
	if node == nil {
		return rightDepth
	}

	return RightLevel(node.Right, rightDepth) + 1
}

func countNodes(root *Methods.TreeNode) int {

	if root == nil {
		return 0

	}

	leftDepth := 0
	rightDepth := 0

	leftDepth = LeftLevel(root, leftDepth)
	rightDepth = RightLevel(root, rightDepth)

	if leftDepth == rightDepth {
		return leftDepth*leftDepth - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}
